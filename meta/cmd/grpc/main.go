package main

import (
	"github.com/go-kit/kit/log"
	"github.com/tiennv147/restless/meta/config"
	"github.com/tiennv147/restless/meta/endpoint"
	"github.com/tiennv147/restless/meta/model"
	"github.com/tiennv147/restless/meta/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"

	"github.com/tiennv147/mazti-commons/db"
	"github.com/tiennv147/restless/meta/repository"
	"github.com/tiennv147/restless/meta/service"
	"github.com/tiennv147/restless/meta/transport"
)

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	cfg, err := config.New()
	checkError(err, logger)

	lis, err := net.Listen("tcp", cfg.GRPC.ListenAddr)
	checkError(err, logger)

	restlessDB, err := db.NewDB(cfg.Database)
	checkError(err, logger)

	restlessDB.AutoMigrate(model.Meta{})

	schemaRepository := repository.NewMetaRepository(restlessDB)
	schemaService := service.NewMetaService(schemaRepository, logger)
	schemaEndpoints := endpoint.New(schemaService, logger)

	s := grpc.NewServer()
	pb.RegisterMetaServer(s, transport.NewGRPCServer(schemaEndpoints))
	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		checkError(err, logger)
	}
}

func checkError(err error, logger log.Logger) {
	if err != nil {
		logger.Log(err.Error())
		os.Exit(-1)
	}
}
