package transport

import (
	"github.com/tiennv147/restless/base/dto"
	base "github.com/tiennv147/restless/base/pb"
	"github.com/tiennv147/restless/base/service"
	shared "github.com/tiennv147/restless/shared/proto"
	netcontext "golang.org/x/net/context"
)

type tableGRPCServer struct {
	tableService service.TableService
}

func NewTableGRPCServer(service service.TableService) base.TableServer {
	return &tableGRPCServer{
		tableService: service,
	}
}

// Implementations

func (g *tableGRPCServer) Create(ctx netcontext.Context, req *base.CreateTableRequest) (*shared.EmptyMsg, error) {
	var columns []dto.Column
	for _, c := range req.Columns {
		columns = append(columns, dto.Column{Name: c.Name, Attributes: c.Attributes})
	}
	tableReq := dto.CreateTableReq{Schema: req.Schema, Name: req.Name, Columns: columns}
	err := g.tableService.Create(ctx, tableReq)
	if err != nil {
		return nil, err
	}
	return &shared.EmptyMsg{}, nil
}
