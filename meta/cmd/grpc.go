package cmd

import (
	"github.com/mazti/restless/meta/model"
	"github.com/mazti/restless/meta/pb"
	"github.com/mazti/restless/meta/repository"
	"github.com/mazti/restless/meta/service"
	"github.com/mazti/restless/meta/transport"
	"github.com/tiennv147/mazti-commons/db"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func RunGRPC(listener net.Listener) error {
	restlessDB, err := db.NewDB(Config.Database)
	CheckError(err)

	restlessDB.AutoMigrate(model.Meta{})

	schemaRepository := repository.NewMetaRepository(restlessDB)
	schemaService := service.NewMetaService(schemaRepository)

	s := grpc.NewServer()
	pb.RegisterMetaServer(s, transport.NewGRPCServer(schemaService))
	// Register reflection service on gRPC server.
	reflection.Register(s)

	return s.Serve(listener)
}
