package transport

import (
	"context"
	"github.com/tiennv147/restless/meta/dto"
	"github.com/tiennv147/restless/meta/endpoint"
	"github.com/tiennv147/restless/meta/transport/pb"

	"github.com/go-kit/kit/transport/grpc"
	netcontext "golang.org/x/net/context"
)

type grpcServer struct {
	create grpc.Handler
	get    grpc.Handler
	list   grpc.Handler
}

func NewGRPCServer(endpoints endpoint.Endpoints) pb.MetaServer {
	return &grpcServer{
		create: makeCreateHandler(endpoints),
		get:    makeGetHandler(endpoints),
		list:   makeListHandler(endpoints),
	}
}

func makeCreateHandler(endpoints endpoint.Endpoints) grpc.Handler {
	return grpc.NewServer(endpoints.CreateMeta, decodeCreateRequest, encodeRegisterResponse)
}

func makeGetHandler(endpoints endpoint.Endpoints) grpc.Handler {
	return grpc.NewServer(endpoints.GetMeta, decodeGetRequest, encodeGetResponse)
}

func makeListHandler(endpoints endpoint.Endpoints) grpc.Handler {
	return grpc.NewServer(endpoints.ListMeta, decodeListRequest, encodeListResponse)
}

// Implementations

func (g *grpcServer) Create(ctx netcontext.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	_, rep, err := g.create.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateResponse), nil
}

func (g *grpcServer) Get(ctx netcontext.Context, req *pb.GetRequest) (*pb.GetReply, error) {
	_, rep, err := g.get.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetReply), nil
}

func (g *grpcServer) List(ctx netcontext.Context, req *pb.ListRequest) (*pb.ListReply, error) {
	_, rep, err := g.list.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ListReply), nil
}

// Decoder and encoder

func decodeCreateRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.CreateRequest)
	return dto.CreateMetaReq{
		Name: req.Name,
	}, nil
}

func encodeRegisterResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(dto.MetaResp)
	return &pb.CreateResponse{
		Id: resp.ID,
	}, nil
}

func decodeGetRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.GetRequest)
	return endpoint.RequestWithID{
		ID: req.Id,
	}, nil
}

func encodeGetResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(dto.MetaResp)
	return &pb.GetReply{
		Id:   resp.ID,
		Name: resp.Name,
	}, nil
}

func decodeListRequest(_ context.Context, _ interface{}) (interface{}, error) {
	return dto.ListMetaResp{}, nil
}

func encodeListResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(dto.ListMetaResp)
	var metas []*pb.MetaModel
	for _, u := range resp.Results {
		su := pb.MetaModel{Id: u.ID, Name: u.Name}
		metas = append(metas, &su)
	}
	return &pb.ListReply{
		Metas: metas,
	}, nil
}
