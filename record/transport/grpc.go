package transport

import (
	"context"
	"errors"

	"github.com/tiennv147/restless/record/pb"
	"github.com/tiennv147/restless/record/service"
)

type grpcServer struct {
	RecordService service.RecordService
}

func NewGRPCServer(service service.RecordService) pb.RecordServer {
	return &grpcServer{
		RecordService: service,
	}
}

// Implementations
func (g *grpcServer) Create(context.Context, *pb.ListRecordsRequest) (*pb.Records, error) {
	return nil, errors.New("please implement me")
}
