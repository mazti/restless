package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/tiennv147/mazti-commons/endpoints"

	"github.com/tiennv147/restless/dto"
	"github.com/tiennv147/restless/service"
)

type SchemaEndpoints struct {
	CreateSchema endpoint.Endpoint
}

func NewSchemaEndpoints(schemaService service.SchemaService, logger log.Logger) SchemaEndpoints {

	var createSchemaEndpoint endpoint.Endpoint
	{
		createSchemaEndpoint = MakeCreateSchemaEndpoint(schemaService)
		createSchemaEndpoint = endpoints.LoggingMiddleware(log.With(logger, "method", "Create"))(createSchemaEndpoint)
	}

	return SchemaEndpoints{
		CreateSchema: createSchemaEndpoint,
	}
}

func MakeCreateSchemaEndpoint(h service.SchemaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.CreateSchemaReq)
		return nil, h.Create(ctx, req)
	}
}
