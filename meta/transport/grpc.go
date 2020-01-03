package transport

import (
	"github.com/tiennv147/restless/meta/dto"
	meta "github.com/tiennv147/restless/meta/pb"
	"github.com/tiennv147/restless/meta/service"
	shared "github.com/tiennv147/restless/shared/proto"

	netcontext "golang.org/x/net/context"
)

type grpcServer struct {
	metaService service.MetaService
}

func NewGRPCServer(service service.MetaService) meta.MetaServer {
	return &grpcServer{
		metaService: service,
	}
}

// Implementations

func (g *grpcServer) Create(ctx netcontext.Context, req *meta.CreateMetaRequest) (*meta.CreateMetaReply, error) {
	resp, err := g.metaService.Create(ctx, dto.CreateMetaReq{Name: req.Name,})
	if err != nil {
		return nil, err
	}
	return &meta.CreateMetaReply{
		Id:   resp.ID,
		Name: resp.Name,
	}, nil
}

func (g *grpcServer) Get(ctx netcontext.Context, req *shared.GetRequest) (*meta.GetMetaReply, error) {
	resp, err := g.metaService.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &meta.GetMetaReply{
		Id:   resp.ID,
		Name: resp.Name,
	}, nil
}

func (g *grpcServer) List(ctx netcontext.Context, req *shared.ListRequest) (*meta.ListMetaReply, error) {
	resp, err := g.metaService.List(ctx, int(req.Offset), int(req.Limit))
	if err != nil {
		return nil, err
	}
	var metas []*meta.MetaModel
	for _, u := range resp.Results {
		su := meta.MetaModel{Id: u.ID, Name: u.Name}
		metas = append(metas, &su)
	}

	return &meta.ListMetaReply{
		Metas: metas,
		Metadata: &shared.ListMetadata{
			Count:  int32(resp.Metadata.Count),
			Limit:  int32(resp.Metadata.Limit),
			Offset: int32(resp.Metadata.Offset),
			Total:  resp.Metadata.Total,
		},
	}, nil
}

func (g *grpcServer) Update(ctx netcontext.Context, req *meta.UpdateMetaRequest) (*meta.UpdateMetaReply, error) {
	resp, err := g.metaService.Update(ctx, dto.UpdateMetaReq{ID: req.Id, Name: req.Name})
	if err != nil {
		return nil, err
	}
	return &meta.UpdateMetaReply{
		Id:   resp.ID,
		Name: resp.Name,
	}, nil
}

func (g *grpcServer) Delete(ctx netcontext.Context, req *shared.DeleteRequest) (*shared.EmptyMsg, error) {
	err := g.metaService.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &shared.EmptyMsg{}, nil
}
