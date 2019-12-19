package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	cendpoints "github.com/tiennv147/restless/commons/endpoints"
	"github.com/tiennv147/restless/dto"
	"github.com/tiennv147/restless/service"
)

func MakeCreateTVProgramEndpoint(h service.TVProgramService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.CreateTVProgramReq)
		return h.Create(ctx, req)
	}
}

func MakeGetTVProgramEndpoint(h service.TVProgramService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(cendpoints.RequestWithID)
		return h.Get(ctx, req.ID)
	}
}

func MakeListTVProgramEndpoint(h service.TVProgramService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(cendpoints.ListRequest)
		return h.List(ctx, req.Offset, req.Limit)
	}
}

func MakeUpdateTVProgramEndpoint(h service.TVProgramService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.UpdateTVProgramReq)
		return h.Update(ctx, req)
	}
}

func MakeDeleteTVProgramEndpoint(h service.TVProgramService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(cendpoints.RequestWithID)
		return nil, h.Delete(ctx, req.ID)
	}
}
