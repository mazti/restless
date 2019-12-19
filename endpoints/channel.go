package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	cendpoints "github.com/tiennv147/restless/commons/endpoints"
	"github.com/tiennv147/restless/dto"
	"github.com/tiennv147/restless/service"
)

func MakeCreateChannelEndpoint(h service.ChannelService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.CreateChannelReq)
		return h.Create(ctx, req)
	}
}

func MakeGetChannelEndpoint(h service.ChannelService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(cendpoints.RequestWithID)
		return h.Get(ctx, req.ID)
	}
}

func MakeListChannelEndpoint(h service.ChannelService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(cendpoints.ListRequest)
		return h.List(ctx, req.Offset, req.Limit)
	}
}

func MakeUpdateChannelEndpoint(h service.ChannelService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.UpdateChannelReq)
		return h.Update(ctx, req)
	}
}

func MakeDeleteChannelEndpoint(h service.ChannelService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(cendpoints.RequestWithID)
		return nil, h.Delete(ctx, req.ID)
	}
}
