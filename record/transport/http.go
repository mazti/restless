package transport

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"

	ctransport "github.com/tiennv147/mazti-commons/transport"
	"github.com/tiennv147/restless/record/dto"
	"github.com/tiennv147/restless/record/endpoints"
)

func NewHTTPHandler(endpoints endpoints.SchemaEndpoints,
	recordEndpoints endpoints.RecordEndpoints, logger log.Logger) http.Handler {
	r := mux.NewRouter()

	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(ctransport.EncodeError),
	}

	// Schema-side
	r.Methods("POST").
		Path("/schema").
		Handler(httptransport.NewServer(
			endpoints.CreateSchema,
			decodeCreateSchemaRequest,
			ctransport.EncodeCommonResponse,
			options...,
		))

	// Data-side
	r.Methods("GET").
		Path("/data/{database-id}/{table-name}").
		Handler(httptransport.NewServer(
			recordEndpoints.SelectRecords,
			ctransport.DecodeListCommonRequest,
			ctransport.EncodeCommonResponse,
			options...,
		))

	return r
}

func decodeCreateSchemaRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req dto.CreateSchemaReq
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
