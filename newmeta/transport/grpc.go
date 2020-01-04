package transport

import (
	"github.com/mazti/restless/newmeta/dto"
	meta "github.com/mazti/restless/newmeta/pb"
	"github.com/mazti/restless/newmeta/service"
	shared "github.com/mazti/restless/shared/proto"

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
	resp, err := g.metaService.Create(ctx, dto.CreateMetaReq{Base: req.Base})
	if err != nil {
		return nil, err
	}
	return &meta.CreateMetaReply{
		Id:     resp.ID,
		Base:   resp.Base,
		Schema: resp.Schema,
	}, nil
}

func (g *grpcServer) Get(ctx netcontext.Context, req *shared.GetRequest) (*meta.GetMetaReply, error) {
	resp, err := g.metaService.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &meta.GetMetaReply{
		Id:     resp.ID,
		Base:   resp.Base,
		Schema: resp.Schema,
	}, nil
}

func (g *grpcServer) List(ctx netcontext.Context, req *shared.ListRequest) (*meta.ListMetaReply, error) {
	resp, err := g.metaService.List(ctx, int(req.Offset), int(req.Limit))
	if err != nil {
		return nil, err
	}
	var metas []*meta.MetaModel
	for _, u := range resp.Results {
		su := meta.MetaModel{Id: u.ID, Base: u.Base, Schema: u.Schema}
		metas = append(metas, &su)
	}

	return &meta.ListMetaReply{
		Metas: metas,
		Metadata: &shared.ListMetadata{
			Count:  int32(resp.Metadata.Count),
			Limit:  int32(resp.Metadata.Limit),
			Offset: int32(resp.Metadata.Offset),
			Total:  int32(resp.Metadata.Total),
		},
	}, nil
}

func (g *grpcServer) Update(ctx netcontext.Context, req *meta.UpdateMetaRequest) (*meta.UpdateMetaReply, error) {
	resp, err := g.metaService.Update(ctx, dto.UpdateMetaReq{ID: req.Id, Base: req.Base})
	if err != nil {
		return nil, err
	}
	return &meta.UpdateMetaReply{
		Id:     resp.ID,
		Base:   resp.Base,
		Schema: resp.Schema,
	}, nil
}

func (g *grpcServer) Delete(ctx netcontext.Context, req *shared.DeleteRequest) (*shared.EmptyMsg, error) {
	err := g.metaService.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &shared.EmptyMsg{}, nil
}
