package endpoints

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"

	cendpoints "github.com/tiennv147/restless/commons/endpoints"
	"github.com/tiennv147/restless/service"
)

type Endpoints struct {
	CreateChannel endpoint.Endpoint
	GetChannel    endpoint.Endpoint
	ListChannel   endpoint.Endpoint
	UpdateChannel endpoint.Endpoint
	DeleteChannel endpoint.Endpoint

	CreateTVProgram endpoint.Endpoint
	GetTVProgram    endpoint.Endpoint
	ListTVProgram   endpoint.Endpoint
	UpdateTVProgram endpoint.Endpoint
	DeleteTVProgram endpoint.Endpoint
}

func New(channelService service.ChannelService,
	tvProgramService service.TVProgramService,
	logger log.Logger) Endpoints {

	var createChannelEndpoint endpoint.Endpoint
	{
		createChannelEndpoint = MakeCreateChannelEndpoint(channelService)
		createChannelEndpoint = cendpoints.LoggingMiddleware(log.With(logger, "method", "Create"))(createChannelEndpoint)
	}

	var getChannelEndpoint endpoint.Endpoint
	{
		getChannelEndpoint = MakeGetChannelEndpoint(channelService)
		getChannelEndpoint = cendpoints.LoggingMiddleware(log.With(logger, "method", "Get"))(getChannelEndpoint)
	}

	var listChannelEndpoint endpoint.Endpoint
	{
		listChannelEndpoint = MakeListChannelEndpoint(channelService)
		listChannelEndpoint = cendpoints.LoggingMiddleware(log.With(logger, "method", "List"))(listChannelEndpoint)
	}

	var updateChannelEndpoint endpoint.Endpoint
	{
		updateChannelEndpoint = MakeUpdateChannelEndpoint(channelService)
		updateChannelEndpoint = cendpoints.LoggingMiddleware(log.With(logger, "method", "Update"))(updateChannelEndpoint)
	}

	var deleteChannelEndpoint endpoint.Endpoint
	{
		deleteChannelEndpoint = MakeDeleteChannelEndpoint(channelService)
		deleteChannelEndpoint = cendpoints.LoggingMiddleware(log.With(logger, "method", "Delete"))(deleteChannelEndpoint)
	}

	var createTVProgramEndpoint endpoint.Endpoint
	{
		createTVProgramEndpoint = MakeCreateTVProgramEndpoint(tvProgramService)
		createTVProgramEndpoint = cendpoints.LoggingMiddleware(log.With(logger, "method", "Create"))(createTVProgramEndpoint)
	}

	var getTVProgramEndpoint endpoint.Endpoint
	{
		getTVProgramEndpoint = MakeGetTVProgramEndpoint(tvProgramService)
		getTVProgramEndpoint = cendpoints.LoggingMiddleware(log.With(logger, "method", "Get"))(getTVProgramEndpoint)
	}

	var listTVProgramEndpoint endpoint.Endpoint
	{
		listTVProgramEndpoint = MakeListTVProgramEndpoint(tvProgramService)
		listTVProgramEndpoint = cendpoints.LoggingMiddleware(log.With(logger, "method", "List"))(listTVProgramEndpoint)
	}

	var updateTVProgramEndpoint endpoint.Endpoint
	{
		updateTVProgramEndpoint = MakeUpdateTVProgramEndpoint(tvProgramService)
		updateTVProgramEndpoint = cendpoints.LoggingMiddleware(log.With(logger, "method", "Update"))(updateTVProgramEndpoint)
	}

	var deleteTVProgramEndpoint endpoint.Endpoint
	{
		deleteTVProgramEndpoint = MakeDeleteTVProgramEndpoint(tvProgramService)
		deleteTVProgramEndpoint = cendpoints.LoggingMiddleware(log.With(logger, "method", "Delete"))(deleteTVProgramEndpoint)
	}

	return Endpoints{
		CreateChannel: createChannelEndpoint,
		GetChannel:    getChannelEndpoint,
		ListChannel:   listChannelEndpoint,
		UpdateChannel: updateChannelEndpoint,
		DeleteChannel: deleteChannelEndpoint,

		CreateTVProgram: createTVProgramEndpoint,
		GetTVProgram:    getTVProgramEndpoint,
		ListTVProgram:   listTVProgramEndpoint,
		UpdateTVProgram: updateTVProgramEndpoint,
		DeleteTVProgram: deleteTVProgramEndpoint,
	}
}
