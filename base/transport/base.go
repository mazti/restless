package transport

import (
	"github.com/mazti/restless/base/dto"
	base "github.com/mazti/restless/base/pb"
	"github.com/mazti/restless/base/service"
	shared "github.com/mazti/restless/shared/proto"
	netcontext "golang.org/x/net/context"
)

type baseGRPCServer struct {
	baseService service.BaseService
}

func NewBaseGRPCServer(service service.BaseService) base.BaseServer {
	return &baseGRPCServer{
		baseService: service,
	}
}

// Implementations

func (g *baseGRPCServer) Create(ctx netcontext.Context, req *base.CreateBaseRequest) (*base.CreateBaseReply, error) {
	resp, err := g.baseService.Create(ctx, dto.CreateBaseReq{Base: req.Base})
	if err != nil {
		return nil, err
	}
	return &base.CreateBaseReply{
		Id:        resp.ID,
		Base:      resp.Base,
		CreatedAt: resp.CreatedAt,
		UpdatedAt: resp.UpdatedAt,
	}, nil
}

func (g *baseGRPCServer) Get(ctx netcontext.Context, req *shared.GetRequest) (*base.GetBaseReply, error) {
	resp, err := g.baseService.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &base.GetBaseReply{
		Id:        resp.ID,
		Base:      resp.Base,
		CreatedAt: resp.CreatedAt,
		UpdatedAt: resp.UpdatedAt,
	}, nil
}

func (g *baseGRPCServer) List(ctx netcontext.Context, req *shared.ListRequest) (*base.ListBaseReply, error) {
	resp, err := g.baseService.List(ctx, int(req.Offset), int(req.Limit))
	if err != nil {
		return nil, err
	}
	var bases []*base.BaseModel
	for _, u := range resp.Results {
		su := base.BaseModel{
			Id:        u.ID,
			Base:      u.Base,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		}
		bases = append(bases, &su)
	}
	return &base.ListBaseReply{
		Bases: bases,
		Metadata: &shared.ListMetadata{
			Count:  int32(resp.Metadata.Count),
			Limit:  int32(resp.Metadata.Limit),
			Offset: int32(resp.Metadata.Offset),
			Total:  int32(resp.Metadata.Total),
		},
	}, nil
}

func (g *baseGRPCServer) Update(ctx netcontext.Context, req *base.UpdateBaseRequest) (*base.UpdateBaseReply, error) {
	resp, err := g.baseService.Update(ctx, dto.UpdateBaseReq{ID: req.Id, Base: req.Base})
	if err != nil {
		return nil, err
	}
	return &base.UpdateBaseReply{
		Id:        resp.ID,
		Base:      resp.Base,
		CreatedAt: resp.CreatedAt,
		UpdatedAt: resp.UpdatedAt,
	}, nil
}

func (g *baseGRPCServer) Delete(ctx netcontext.Context, req *shared.DeleteRequest) (*shared.EmptyMsg, error) {
	err := g.baseService.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &shared.EmptyMsg{}, nil
}
