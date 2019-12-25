package endpoint

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"

	cendpoints "github.com/tiennv147/mazti-commons/endpoints"
	"github.com/tiennv147/restless/meta/service"
)

type Endpoints struct {
	CreateMeta endpoint.Endpoint
	GetMeta    endpoint.Endpoint
	ListMeta   endpoint.Endpoint
	UpdateMeta endpoint.Endpoint
	DeleteMeta endpoint.Endpoint
}

func New(metaService service.MetaService, logger log.Logger) Endpoints {

	var createMetaEndpoint endpoint.Endpoint
	{
		createMetaEndpoint = MakeCreateMetaEndpoint(metaService)
		createMetaEndpoint = cendpoints.LoggingMiddleware(log.With(logger, "method", "Create"))(createMetaEndpoint)
	}

	var getMetaEndpoint endpoint.Endpoint
	{
		getMetaEndpoint = MakeGetMetaEndpoint(metaService)
		getMetaEndpoint = cendpoints.LoggingMiddleware(log.With(logger, "method", "Get"))(getMetaEndpoint)
	}

	var listMetaEndpoint endpoint.Endpoint
	{
		listMetaEndpoint = MakeListMetaEndpoint(metaService)
		listMetaEndpoint = cendpoints.LoggingMiddleware(log.With(logger, "method", "List"))(listMetaEndpoint)
	}

	var updateMetaEndpoint endpoint.Endpoint
	{
		updateMetaEndpoint = MakeUpdateMetaEndpoint(metaService)
		updateMetaEndpoint = cendpoints.LoggingMiddleware(log.With(logger, "method", "Update"))(updateMetaEndpoint)
	}

	var deleteMetaEndpoint endpoint.Endpoint
	{
		deleteMetaEndpoint = MakeDeleteMetaEndpoint(metaService)
		deleteMetaEndpoint = cendpoints.LoggingMiddleware(log.With(logger, "method", "Delete"))(deleteMetaEndpoint)
	}

	return Endpoints{
		CreateMeta: createMetaEndpoint,
		GetMeta:    getMetaEndpoint,
		ListMeta:   listMetaEndpoint,
		UpdateMeta: updateMetaEndpoint,
		DeleteMeta: deleteMetaEndpoint,
	}
}
