package cmd

import (
	pb "github.com/tiennv147/restless/base/pb/base"
	"github.com/tiennv147/restless/base/repository"
	"github.com/tiennv147/restless/base/service"
	"github.com/tiennv147/restless/base/transport"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func RunGRPC(listener net.Listener) error {

	baseRepository, err := repository.NewBaseRepository(Config.Database)
	CheckError(err)
	schemaService, err := service.NewBaseService(baseRepository, *Config)
	CheckError(err)

	s := grpc.NewServer()
	pb.RegisterBaseServer(s, transport.NewGRPCServer(schemaService))
	// Register reflection service on gRPC server.
	reflection.Register(s)

	return s.Serve(listener)
}
