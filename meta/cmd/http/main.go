package main

import (
	"github.com/go-kit/kit/log"
	"github.com/tiennv147/restless/meta/config"
	"net/http"
	"os"
	"path"
	"strings"

	"google.golang.org/grpc"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	gw "github.com/tiennv147/restless/meta/pb"
	"golang.org/x/net/context"
)

var (
	conf   *config.Config
	logger log.Logger
)

func newGateway(ctx context.Context, opts ...runtime.ServeMuxOption) (http.Handler, error) {
	mux := runtime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterMetaHandlerFromEndpoint(ctx, mux, "0.0.0.0"+conf.GRPC.ListenAddr, dialOpts)
	if err != nil {
		return nil, err
	}

	err = gw.RegisterMetaHandlerFromEndpoint(ctx, mux, "0.0.0.0"+conf.GRPC.ListenAddr, dialOpts)
	if err != nil {
		return nil, err
	}

	return mux, nil
}

func serveSwagger(w http.ResponseWriter, r *http.Request) {
	if !strings.HasSuffix(r.URL.Path, ".swagger.json") {
		logger.Log("ERROR", "Swagger JSON not Found: %s", r.URL.Path)
		http.NotFound(w, r)
		return
	}

	logger.Log("INFO", "Serving %s", r.URL.Path)
	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	p = path.Join(*conf.SwaggerDir, p)
	http.ServeFile(w, r, p)
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	logger.Log("INFO", "preflight request for %s", r.URL.Path)
	return
}

// allowCORS allows Cross Origin Resoruce Sharing from any origin.
// Don't do this without consideration in production systems.
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func Run(opts ...runtime.ServeMuxOption) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	//mux := runtime.NewServeMux()
	mux := http.NewServeMux()

	mux.HandleFunc("/swagger/", serveSwagger)

	gateway, err := newGateway(ctx, opts...)
	if err != nil {
		return err
	}
	mux.Handle("/", gateway)

	return http.ListenAndServe(conf.HTTP.ListenAddr, allowCORS(mux))

}

func main() {
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	var err error
	conf, err = config.New()
	checkError(err)

	err = Run()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		logger.Log("ERROR", err.Error())
		os.Exit(-1)
	}
}
