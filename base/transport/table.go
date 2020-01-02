package transport

import (
	"github.com/tiennv147/restless/base/dto"
	pb "github.com/tiennv147/restless/base/pb/base"
	"github.com/tiennv147/restless/base/service"
	netcontext "golang.org/x/net/context"
)

type tableGRPCServer struct {
	tableService service.TableService
}

func NewTableGRPCServer(service service.TableService) pb.TableServer {
	return &tableGRPCServer{
		tableService: service,
	}
}

// Implementations

func (g *tableGRPCServer) Create(ctx netcontext.Context, req *pb.CreateTableRequest) (*pb.EmptyMsg, error) {
	var columns []dto.Column
	for _, c := range req.Columns {
		columns = append(columns, dto.Column{Name: c.Name, Attributes: c.Attributes})
	}
	tableReq := dto.CreateTableReq{Schema: req.Schema, Name: req.Name, Columns: columns}
	err := g.tableService.Create(ctx, tableReq)
	if err != nil {
		return nil, err
	}
	return &pb.EmptyMsg{}, nil
}
