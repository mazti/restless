package transport

import (
	"github.com/tiennv147/restless/base/dto"
	pb "github.com/tiennv147/restless/base/pb/base"
	"github.com/tiennv147/restless/base/service"
	netcontext "golang.org/x/net/context"
)

type grpcServer struct {
	baseService service.BaseService
}

func NewGRPCServer(service service.BaseService) pb.BaseServer {
	return &grpcServer{
		baseService: service,
	}
}

// Implementations

func (g *grpcServer) Create(ctx netcontext.Context, req *pb.CreateBaseRequest) (*pb.CreateBaseReply, error) {
	base, err := g.baseService.Create(ctx, dto.CreateBaseReq{Name: req.Name,})
	if err != nil {
		return nil, err
	}
	return &pb.CreateBaseReply{
		Id:   base.ID,
		Name: base.Name,
	}, nil
}

func (g *grpcServer) Get(ctx netcontext.Context, req *pb.GetBaseRequest) (*pb.GetBaseReply, error) {
	resp, err := g.baseService.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetBaseReply{
		Id:   resp.ID,
		Name: resp.Name,
	}, nil
}

func (g *grpcServer) List(ctx netcontext.Context, req *pb.ListBaseRequest) (*pb.ListBaseReply, error) {
	resp, err := g.baseService.List(ctx, int(req.Offset), int(req.Limit))
	if err != nil {
		return nil, err
	}
	var bases []*pb.BaseModel
	for _, u := range resp.Results {
		su := pb.BaseModel{Id: u.ID, Name: u.Name}
		bases = append(bases, &su)
	}
	return &pb.ListBaseReply{
		Bases: bases,
		Metadata: &pb.ListMetadata{
			Count:  int32(resp.Metadata.Count),
			Limit:  int32(resp.Metadata.Limit),
			Offset: int32(resp.Metadata.Offset),
			Total:  resp.Metadata.Total,
		},
	}, nil
}

func (g *grpcServer) Update(ctx netcontext.Context, req *pb.UpdateBaseRequest) (*pb.UpdateBaseReply, error) {
	resp, err := g.baseService.Update(ctx, dto.UpdateBaseReq{ID: req.Id, Name: req.Name})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateBaseReply{
		Id:   resp.ID,
		Name: resp.Name,
	}, nil
}

func (g *grpcServer) Delete(ctx netcontext.Context, req *pb.DeleteBaseRequest) (*pb.DeleteBaseReply, error) {
	err := g.baseService.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteBaseReply{}, nil
}
