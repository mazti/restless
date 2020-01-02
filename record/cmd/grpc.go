package cmd

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/tiennv147/restless/record/pb"
	"github.com/tiennv147/restless/record/service"
	"github.com/tiennv147/restless/record/transport"
)

func RunGRPC(listener net.Listener) error {

	recordService := service.NewRecordService()

	s := grpc.NewServer()
	pb.RegisterRecordServer(s, transport.NewGRPCServer(recordService))
	// Register reflection service on gRPC server.
	reflection.Register(s)

	return s.Serve(listener)
}
