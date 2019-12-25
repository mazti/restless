package transport

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/tiennv147/mazti-commons/errors"
	ctransport "github.com/tiennv147/mazti-commons/transport"
	"github.com/tiennv147/restless/meta/dto"
	"github.com/tiennv147/restless/meta/endpoint"
)

func NewHTTPHandler(endpoints endpoint.Endpoints, logger log.Logger) http.Handler {
	r := mux.NewRouter()

	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(ctransport.EncodeError),
	}

	// Meta
	r.Methods("POST").
		Path("/meta").
		Handler(httptransport.NewServer(
			endpoints.CreateMeta,
			decodeCreateMetaRequest,
			ctransport.EncodeCommonResponse,
			options...,
		))
	r.Methods("GET").
		Path("/meta/{id}").
		Handler(httptransport.NewServer(
			endpoints.GetMeta,
			ctransport.DecodeRequestCommonWithID,
			ctransport.EncodeCommonResponse,
			options...,
		))
	r.Methods("GET").
		Path("/meta").
		Handler(httptransport.NewServer(
			endpoints.ListMeta,
			ctransport.DecodeListCommonRequest,
			ctransport.EncodeCommonResponse,
			options...,
		))
	r.Methods("PATCH").
		Path("/meta/{id}").
		Handler(httptransport.NewServer(
			endpoints.UpdateMeta,
			decodeUpdateMetaRequest,
			ctransport.EncodeCommonResponse,
			options...,
		))
	r.Methods("DELETE").
		Path("/meta/{id}").
		Handler(httptransport.NewServer(
			endpoints.DeleteMeta,
			ctransport.DecodeRequestCommonWithID,
			ctransport.EncodeCommonResponse,
			options...,
		))

	return r
}

func decodeCreateMetaRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req dto.CreateMetaReq
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func decodeUpdateMetaRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		return nil, errors.ErrBadRouting
	}

	var req dto.UpdateMetaReq
	err := json.NewDecoder(r.Body).Decode(&req)
	req.ID = idStr
	return req, err
}
