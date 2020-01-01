package cmd

import (
	"github.com/tiennv147/mazti-commons/db"
	"github.com/tiennv147/restless/meta/model"
	pb "github.com/tiennv147/restless/meta/pb/meta"
	"github.com/tiennv147/restless/meta/repository"
	"github.com/tiennv147/restless/meta/service"
	"github.com/tiennv147/restless/meta/transport"
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
