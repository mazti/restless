package cmd

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/mazti/restless/record/pb"
	"github.com/mazti/restless/record/service"
	"github.com/mazti/restless/record/transport"
)

func RunGRPC(listener net.Listener) error {

	recordService := service.NewRecordService()

	s := grpc.NewServer()
	pb.RegisterRecordServer(s, transport.NewGRPCServer(recordService))
	// Register reflection service on gRPC server.
	reflection.Register(s)

	return s.Serve(listener)
}
