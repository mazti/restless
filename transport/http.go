package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/tiennv147/restless/commons/errors"
	ctransport "github.com/tiennv147/restless/commons/transport"
	"github.com/tiennv147/restless/dto"
	"github.com/tiennv147/restless/endpoints"
)

func NewHTTPHandler(endpoints endpoints.Endpoints, logger log.Logger) http.Handler {
	r := mux.NewRouter()

	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(ctransport.EncodeError),
	}

	// Channel
	r.Methods("POST").
		Path("/channels").
		Handler(httptransport.NewServer(
			endpoints.CreateChannel,
			decodeCreateChannelRequest,
			ctransport.EncodeCommonResponse,
			options...,
		))
	r.Methods("GET").
		Path("/channels/{id}").
		Handler(httptransport.NewServer(
			endpoints.GetChannel,
			ctransport.DecodeRequestCommonWithID,
			ctransport.EncodeCommonResponse,
			options...,
		))
	r.Methods("GET").
		Path("/channels").
		Handler(httptransport.NewServer(
			endpoints.ListChannel,
			ctransport.DecodeListCommonRequest,
			ctransport.EncodeCommonResponse,
			options...,
		))
	r.Methods("PATCH").
		Path("/channels/{id}").
		Handler(httptransport.NewServer(
			endpoints.UpdateChannel,
			decodeUpdateChannelRequest,
			ctransport.EncodeCommonResponse,
			options...,
		))
	r.Methods("DELETE").
		Path("/channels/{id}").
		Handler(httptransport.NewServer(
			endpoints.DeleteChannel,
			ctransport.DecodeRequestCommonWithID,
			ctransport.EncodeCommonResponse,
			options...,
		))

	// TVProgram
	r.Methods("POST").
		Path("/tvprograms").
		Handler(httptransport.NewServer(
			endpoints.CreateTVProgram,
			decodeCreateTVProgramRequest,
			ctransport.EncodeCommonResponse,
			options...,
		))
	r.Methods("GET").
		Path("/tvprograms/{id}").
		Handler(httptransport.NewServer(
			endpoints.GetTVProgram,
			ctransport.DecodeRequestCommonWithID,
			ctransport.EncodeCommonResponse,
			options...,
		))
	r.Methods("GET").
		Path("/tvprograms").
		Handler(httptransport.NewServer(
			endpoints.ListTVProgram,
			ctransport.DecodeListCommonRequest,
			ctransport.EncodeCommonResponse,
			options...,
		))
	r.Methods("PATCH").
		Path("/tvprograms/{id}").
		Handler(httptransport.NewServer(
			endpoints.UpdateTVProgram,
			decodeUpdateTVProgramRequest,
			ctransport.EncodeCommonResponse,
			options...,
		))
	r.Methods("DELETE").
		Path("/tvprograms/{id}").
		Handler(httptransport.NewServer(
			endpoints.DeleteTVProgram,
			ctransport.DecodeRequestCommonWithID,
			ctransport.EncodeCommonResponse,
			options...,
		))

	return r
}

func decodeCreateChannelRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req dto.CreateChannelReq
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func decodeUpdateChannelRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		return nil, errors.ErrBadRouting
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return nil, errors.ErrInconsistentIDs
	}

	var req dto.UpdateChannelReq
	err = json.NewDecoder(r.Body).Decode(&req)
	req.ID = id
	return req, err
}

func decodeCreateTVProgramRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req dto.CreateTVProgramReq
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func decodeUpdateTVProgramRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		return nil, errors.ErrBadRouting
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return nil, errors.ErrInconsistentIDs
	}

	var req dto.UpdateTVProgramReq
	err = json.NewDecoder(r.Body).Decode(&req)
	req.ID = id
	return req, err
}
