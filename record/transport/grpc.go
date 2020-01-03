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
func (g *grpcServer) ListRecords(context.Context, *pb.ListRecordsRequest) (*pb.ListRecordsReply, error) {
	return nil, errors.New("please implement me")
}

func (g *grpcServer) CreateRecords(ctx context.Context, req *pb.CreateRecordsRequest) (*pb.ListRecordsReply, error) {
	return nil, errors.New("please implement me")
}
