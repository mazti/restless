package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/tiennv147/mazti-commons/endpoints"

	"github.com/tiennv147/restless/record/dto"
	"github.com/tiennv147/restless/record/service"
)

type RecordEndpoints struct {
	SelectRecords endpoint.Endpoint
}

func NewRecordEndpoints(schemaService service.RecordService, logger log.Logger) RecordEndpoints {

	var selectRecordsEndpoint endpoint.Endpoint
	{
		selectRecordsEndpoint = MakeSelectRecordsEndpoint(schemaService)
		selectRecordsEndpoint = endpoints.LoggingMiddleware(log.With(logger, "method", "Create"))(selectRecordsEndpoint)
	}

	return RecordEndpoints{
		SelectRecords: selectRecordsEndpoint,
	}
}

func MakeSelectRecordsEndpoint(h service.RecordService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.SelectRecordsReq)
		return nil, h.Select(ctx, req)
	}
}
