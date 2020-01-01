package transport

import (
	"github.com/tiennv147/restless/meta/dto"
	pb "github.com/tiennv147/restless/meta/pb/meta"
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

func (g *grpcServer) Create(ctx netcontext.Context, req *pb.CreateMetaRequest) (*pb.CreateMetaReply, error) {
	resp, err := g.metaService.Create(ctx, dto.CreateMetaReq{Name: req.Name,})
	if err != nil {
		return nil, err
	}
	return &pb.CreateMetaReply{
		Id:   resp.ID,
		Name: resp.Name,
	}, nil
}

func (g *grpcServer) Get(ctx netcontext.Context, req *pb.GetMetaRequest) (*pb.GetMetaReply, error) {
	resp, err := g.metaService.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetMetaReply{
		Id:   resp.ID,
		Name: resp.Name,
	}, nil
}

func (g *grpcServer) List(ctx netcontext.Context, req *pb.ListMetaRequest) (*pb.ListMetaReply, error) {
	resp, err := g.metaService.List(ctx, int(req.Offset), int(req.Limit))
	if err != nil {
		return nil, err
	}
	var metas []*pb.MetaModel
	for _, u := range resp.Results {
		su := pb.MetaModel{Id: u.ID, Name: u.Name}
		metas = append(metas, &su)
	}

	return &pb.ListMetaReply{
		Metas: metas,
		Metadata: &pb.ListMetadata{
			Count:  int32(resp.Metadata.Count),
			Limit:  int32(resp.Metadata.Limit),
			Offset: int32(resp.Metadata.Offset),
			Total:  resp.Metadata.Total,
		},
	}, nil
}

func (g *grpcServer) Update(ctx netcontext.Context, req *pb.UpdateMetaRequest) (*pb.UpdateMetaReply, error) {
	resp, err := g.metaService.Update(ctx, dto.UpdateMetaReq{ID: req.Id, Name: req.Name})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateMetaReply{
		Id:   resp.ID,
		Name: resp.Name,
	}, nil
}

func (g *grpcServer) Delete(ctx netcontext.Context, req *pb.DeleteMetaRequest) (*pb.DeleteMetaReply, error) {
	err := g.metaService.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteMetaReply{}, nil
}
