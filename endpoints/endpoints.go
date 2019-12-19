package endpoints

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"

	cendpoints "github.com/tiennv147/mazti-commons/endpoints"
	"github.com/tiennv147/restless/service"
)

type Endpoints struct {
	CreateSample endpoint.Endpoint
	GetSample    endpoint.Endpoint
	ListSample   endpoint.Endpoint
	UpdateSample endpoint.Endpoint
	DeleteSample endpoint.Endpoint
}

func New(channelService service.SampleService, logger log.Logger) Endpoints {

	var createSampleEndpoint endpoint.Endpoint
	{
		createSampleEndpoint = MakeCreateSampleEndpoint(channelService)
		createSampleEndpoint = cendpoints.LoggingMiddleware(log.With(logger, "method", "Create"))(createSampleEndpoint)
	}

	var getSampleEndpoint endpoint.Endpoint
	{
		getSampleEndpoint = MakeGetSampleEndpoint(channelService)
		getSampleEndpoint = cendpoints.LoggingMiddleware(log.With(logger, "method", "Get"))(getSampleEndpoint)
	}

	var listSampleEndpoint endpoint.Endpoint
	{
		listSampleEndpoint = MakeListSampleEndpoint(channelService)
		listSampleEndpoint = cendpoints.LoggingMiddleware(log.With(logger, "method", "List"))(listSampleEndpoint)
	}

	var updateSampleEndpoint endpoint.Endpoint
	{
		updateSampleEndpoint = MakeUpdateSampleEndpoint(channelService)
		updateSampleEndpoint = cendpoints.LoggingMiddleware(log.With(logger, "method", "Update"))(updateSampleEndpoint)
	}

	var deleteSampleEndpoint endpoint.Endpoint
	{
		deleteSampleEndpoint = MakeDeleteSampleEndpoint(channelService)
		deleteSampleEndpoint = cendpoints.LoggingMiddleware(log.With(logger, "method", "Delete"))(deleteSampleEndpoint)
	}

	return Endpoints{
		CreateSample: createSampleEndpoint,
		GetSample:    getSampleEndpoint,
		ListSample:   listSampleEndpoint,
		UpdateSample: updateSampleEndpoint,
		DeleteSample: deleteSampleEndpoint,
	}
}
