package cmd

import (
	"github.com/tiennv147/restless/base/pb/base"
	"log"
	"net"
	"net/http"
	"path"
	"strings"

	"google.golang.org/grpc"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
)

func RunHTTP(listener net.Listener, opts ...runtime.ServeMuxOption) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := http.NewServeMux()

	mux.HandleFunc("/swagger/", serveSwagger)

	gateway, err := newGateway(ctx, opts...)
	if err != nil {
		return err
	}
	mux.Handle("/", gateway)

	return http.Serve(listener, allowCORS(mux))
}

func newGateway(ctx context.Context, opts ...runtime.ServeMuxOption) (http.Handler, error) {
	mux := runtime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}
	err := base.RegisterBaseHandlerFromEndpoint(ctx, mux, "0.0.0.0"+Config.GRPC.ListenAddr, dialOpts)
	if err != nil {
		return nil, err
	}

	err = base.RegisterTableHandlerFromEndpoint(ctx, mux, "0.0.0.0"+Config.GRPC.ListenAddr, dialOpts)
	if err != nil {
		return nil, err
	}

	return mux, nil
}

func serveSwagger(w http.ResponseWriter, r *http.Request) {
	if !strings.HasSuffix(r.URL.Path, ".swagger.json") {
		log.Println("ERROR", "Swagger JSON not Found: %s", r.URL.Path)
		http.NotFound(w, r)
		return
	}

	log.Println("INFO", "Serving %s", r.URL.Path)
	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	p = path.Join(*Config.SwaggerDir, p)
	http.ServeFile(w, r, p)
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	log.Println("INFO", "preflight request for %s", r.URL.Path)
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
