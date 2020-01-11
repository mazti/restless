package transport

import (
	"github.com/mazti/restless/base/dto"
	column "github.com/mazti/restless/base/pb"
	"github.com/mazti/restless/base/service"
	shared "github.com/mazti/restless/shared/proto"
	netcontext "golang.org/x/net/context"
)

type columnGRPCServer struct {
	columnService service.ColumnService
}

func NewColumnGRPCServer(service service.ColumnService) column.ColumnServer {
	return &columnGRPCServer{
		columnService: service,
	}
}

// Implementations

func (g *columnGRPCServer) Create(ctx netcontext.Context, req *column.CreateColumnRequest) (*column.CreateColumnReply, error) {
	columnReq := dto.CreateColumnReq{
		Table:      req.Table,
		Name:       req.Name,
		Type:       req.Type,
		Default:    req.Default,
		TypeOption: dto.TypeOption{},
	}
	if req.TypeOption != nil {
		columnReq.TypeOption.Format = req.TypeOption.Format
		columnReq.TypeOption.Precision = req.TypeOption.Precision
	}
	resp, err := g.columnService.Create(ctx, columnReq)
	if err != nil {
		return nil, err
	}
	option := &column.TypeOption{}
	if resp.TypeOption != nil {
		option.Format = resp.TypeOption.Format
		option.Precision = resp.TypeOption.Precision
	}
	return &column.CreateColumnReply{
		Id:         resp.ID,
		Name:       resp.Name,
		Type:       resp.Type,
		Default:    resp.Default,
		TypeOption: option,
		CreatedAt:  resp.CreatedAt,
		UpdatedAt:  resp.UpdatedAt,
	}, nil
}

func (g *columnGRPCServer) Update(ctx netcontext.Context, req *column.UpdateColumnRequest) (*column.UpdateColumnReply, error) {
	columnReq := dto.UpdateColumnReq{
		ID:      req.Id,
		Table:   req.Table,
		Name:    req.Name,
		Type:    req.Type,
		Default: req.Default,
		TypeOption: dto.TypeOption{
			Format:    req.TypeOption.Format,
			Precision: req.TypeOption.Precision,
		},
	}
	resp, err := g.columnService.Update(ctx, columnReq)
	if err != nil {
		return nil, err
	}
	return &column.UpdateColumnReply{
		Id:      resp.ID,
		Name:    resp.Name,
		Type:    resp.Type,
		Default: resp.Default,
		TypeOption: &column.TypeOption{
			Format:    resp.TypeOption.Format,
			Precision: resp.TypeOption.Precision,
		},
		CreatedAt: resp.CreatedAt,
		UpdatedAt: resp.UpdatedAt,
	}, nil
}

func (g *columnGRPCServer) Delete(ctx netcontext.Context, req *shared.DeleteRequest) (*shared.EmptyMsg, error) {
	err := g.columnService.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &shared.EmptyMsg{}, nil
}
