package grapql

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/graphql-go/graphql"
)

type Server struct {
	ctx    context.Context
	e      endpoint.Endpoint
	before []RequestFunc
	after  []ResponseFunc
	dec    DecodeRequestFunc
	logger log.Logger
}

func NewServer(
	ctx context.Context,
	e endpoint.Endpoint,
	dec DecodeRequestFunc,
	options ...ServerOption) *Server {
	s := &Server{
		ctx:    ctx,
		e:      e,
		dec:    dec,
		logger: log.NewNopLogger(),
	}

	for _, option := range options {
		option(s)
	}
	return s
}

type ServerOption func(*Server)

// ServerBefore functions are executed on the HTTP request object before the
// request is decoded.
func ServerBefore(before ...RequestFunc) ServerOption {
	return func(s *Server) { s.before = before }
}

// ServerAfter functions are executed on the HTTP response writer after the
// endpoint is invoked, but before anything is written to the client.
func ServerAfter(after ...ResponseFunc) ServerOption {
	return func(s *Server) { s.after = append(s.after, after...) }
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := s.ctx

	for _, f := range s.before {
		ctx = f(ctx, r)
	}

	request, err := s.dec(ctx, r)
	if err != nil {
		s.logger.Log("err", err)
		s.errorEncoder(ctx, err, w)
		return
	}

	response, err := s.e(ctx, request)
	if err != nil {
		s.logger.Log("err", err)
		s.errorEncoder(ctx, err, w)
		return
	}

	for _, f := range s.after {
		ctx = f(ctx, w)
	}

	s.encodeResponse(ctx, w, graphql.Do(response.(graphql.Params)))
}

func (s Server) errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
		"data":  nil,
	})
}

func (s Server) encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(error); ok {
		s.errorEncoder(ctx, e, w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
