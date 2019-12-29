package endpoint

import (
	"context"
	"github.com/tiennv147/mazti-commons/endpoints"

	"github.com/go-kit/kit/endpoint"

	"github.com/tiennv147/restless/meta/dto"
	"github.com/tiennv147/restless/meta/service"
)

type RequestWithID struct {
	ID string
}

func MakeCreateMetaEndpoint(h service.MetaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.CreateMetaReq)
		return h.Create(ctx, req)
	}
}

func MakeGetMetaEndpoint(h service.MetaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RequestWithID)
		return h.Get(ctx, req.ID)
	}
}

func MakeListMetaEndpoint(h service.MetaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(endpoints.ListRequest)
		return h.List(ctx, req.Offset, req.Limit)
	}
}

func MakeUpdateMetaEndpoint(h service.MetaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.UpdateMetaReq)
		return h.Update(ctx, req)
	}
}

func MakeDeleteMetaEndpoint(h service.MetaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RequestWithID)
		return nil, h.Delete(ctx, req.ID)
	}
}
