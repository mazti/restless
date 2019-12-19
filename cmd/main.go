package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/oklog/oklog/pkg/group"
	"github.com/tiennv147/mazti-commons/db"

	"github.com/tiennv147/restless/config"
	"github.com/tiennv147/restless/endpoints"

	"github.com/tiennv147/restless/repository"
	"github.com/tiennv147/restless/service"
	"github.com/tiennv147/restless/transport"
)

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	cfg, err := config.New()
	if err != nil {
		logger.Log("Failed to reading config file", err)
		os.Exit(-1)
	}

	db, err := db.NewDB(cfg.Database)
	if err != nil {
		logger.Log("Failed to initialize model for operating all service, %s\n", err)
	}

	cRepo := repository.NewSampleRepository(db)
	cService := service.NewSampleService(cRepo, logger)

	cEndpoints := endpoints.New(cService, logger)
	cHandler := transport.NewHTTPHandler(cEndpoints, logger)

	var g group.Group
	{
		// The debug listener mounts the http.DefaultServeMux, and serves up
		// stuff like the Prometheus metrics route, the Go debug and profiling
		// routes, and so on.
		debugListener, err := net.Listen("tcp", cfg.HTTP.DebugAddr)
		if err != nil {
			_ = logger.Log("transport", "debug/HTTP", "during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(func() error {
			logger.Log("transport", "debug/HTTP", "addr", cfg.HTTP.DebugAddr)
			return http.Serve(debugListener, http.DefaultServeMux)
		}, func(error) {
			_ = debugListener.Close()
		})
	}

	{
		// The HTTP listener mounts the Go kit HTTP handler we created.
		httpListener, err := net.Listen("tcp", cfg.HTTP.ListenAddr)
		if err != nil {
			logger.Log("transport", "HTTP", "during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(func() error {
			logger.Log("transport", "HTTP", "addr", cfg.HTTP.ListenAddr)
			return http.Serve(httpListener, cHandler)
		}, func(error) {
			_ = httpListener.Close()
		})
	}

	{
		// This function just sits and waits for ctrl-C.
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}
	logger.Log("exit", g.Run())
}
