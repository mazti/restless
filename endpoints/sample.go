package endpoints

import (
	"context"
	"github.com/tiennv147/mazti-commons/endpoints"

	"github.com/go-kit/kit/endpoint"

	"github.com/tiennv147/restless/dto"
	"github.com/tiennv147/restless/service"
)

func MakeCreateSampleEndpoint(h service.SampleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.CreateSampleReq)
		return h.Create(ctx, req)
	}
}

func MakeGetSampleEndpoint(h service.SampleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(endpoints.RequestWithID)
		return h.Get(ctx, req.ID)
	}
}

func MakeListSampleEndpoint(h service.SampleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(endpoints.ListRequest)
		return h.List(ctx, req.Offset, req.Limit)
	}
}

func MakeUpdateSampleEndpoint(h service.SampleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.UpdateSampleReq)
		return h.Update(ctx, req)
	}
}

func MakeDeleteSampleEndpoint(h service.SampleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(endpoints.RequestWithID)
		return nil, h.Delete(ctx, req.ID)
	}
}
