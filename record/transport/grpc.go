package transport

import (
	"context"
	"errors"
	"github.com/mazti/restless/record/dto"
	"github.com/mazti/restless/record/pb"
	"github.com/mazti/restless/record/service"
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
func (g *grpcServer) ListRecords(ctx context.Context, req *pb.ListRecordsRequest) (*pb.ListRecordsReply, error) {
	var (
		records    []interface{}
		recordsMsg []*pb.RecordMsg
		err        error
	)
	if records, err = g.RecordService.List(ctx, req.Base, req.Table, req.MaxRecords, req.PageSize); err != nil {
		return nil, err
	}
	if recordsMsg, err = dto.ConvertRecords(records); err != nil {
		return nil, err
	}
	return &pb.ListRecordsReply{
		Records: recordsMsg,
	}, nil
}

func (g *grpcServer) RetrieveRecord(ctx context.Context, req *pb.RetrievesRequest) (*pb.RecordMsg, error) {
	var (
		record interface{}
		err    error
	)
	if record, err = g.RecordService.Get(ctx, req.Base, req.Table, req.RecordID); err != nil {
		return nil, err
	}
	return dto.ConvertRecord(record)
}

func (g *grpcServer) UpdateRecords(ctx context.Context, req *pb.CreateRecordsRequest) (*pb.ListRecordsReply, error) {
	return nil, errors.New("please implement me")
}

func (g *grpcServer) CreateRecords(ctx context.Context, req *pb.CreateRecordsRequest) (*pb.ListRecordsReply, error) {
	return nil, errors.New("please implement me")
}

func (g *grpcServer) DeleteRecords(ctx context.Context, req *pb.DeleteRecordsRequest) (*pb.DeleteRecordsReply, error) {
	return nil, errors.New("please implement me")
}
