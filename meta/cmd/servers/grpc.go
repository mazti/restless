package servers

import (
	"github.com/tiennv147/mazti-commons/db"
	"github.com/tiennv147/restless/meta/endpoint"
	"github.com/tiennv147/restless/meta/model"
	"github.com/tiennv147/restless/meta/pb"
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
	schemaService := service.NewMetaService(schemaRepository, Logger)
	schemaEndpoints := endpoint.New(schemaService, Logger)

	s := grpc.NewServer()
	pb.RegisterMetaServer(s, transport.NewGRPCServer(schemaEndpoints))
	// Register reflection service on gRPC server.
	reflection.Register(s)

	return s.Serve(listener)
}
