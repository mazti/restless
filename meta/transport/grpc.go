package transport

import (
	"github.com/tiennv147/restless/meta/dto"
	"github.com/tiennv147/restless/meta/pb"
	"github.com/tiennv147/restless/meta/service"

	netcontext "golang.org/x/net/context"
)

type grpcServer struct {
	metaService service.MetaService
}

func NewGRPCServer(service service.MetaService) pb.MetaServer {
	return &grpcServer{
		metaService: service,
	}
}

// Implementations

func (g *grpcServer) Create(ctx netcontext.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	resp, err := g.metaService.Create(ctx, dto.CreateMetaReq{Name: req.Name,})
	if err != nil {
		return nil, err
	}
	return &pb.CreateResponse{
		Id:   resp.ID,
		Name: resp.Name,
	}, nil
}

func (g *grpcServer) Get(ctx netcontext.Context, req *pb.GetRequest) (*pb.GetReply, error) {
	resp, err := g.metaService.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetReply{
		Id:   resp.ID,
		Name: resp.Name,
	}, nil
}

func (g *grpcServer) List(ctx netcontext.Context, req *pb.ListRequest) (*pb.ListReply, error) {
	resp, err := g.metaService.List(ctx, int(req.Offset), int(req.Limit))
	if err != nil {
		return nil, err
	}
	var metas []*pb.MetaModel
	for _, u := range resp.Results {
		su := pb.MetaModel{Id: u.ID, Name: u.Name}
		metas = append(metas, &su)
	}
	return &pb.ListReply{
		Metas: metas,
	}, nil
}

func (g *grpcServer) Update(ctx netcontext.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	resp, err := g.metaService.Update(ctx, dto.UpdateMetaReq{ID: req.Id, Name: req.Name})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateResponse{
		Id:   resp.ID,
		Name: resp.Name,
	}, nil
}

func (g *grpcServer) Delete(ctx netcontext.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	err := g.metaService.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteResponse{}, nil
}