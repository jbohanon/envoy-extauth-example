package main

import (
	"context"
	"regexp"

	extauth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	"google.golang.org/genproto/googleapis/rpc/status"
)

var contentTypeJSON = regexp.MustCompile("^(application|text)/json(;.*)?$")

var _ extauth.AuthorizationServer = &authorizationService{}

type authorizationService struct{}

func (s *authorizationService) Check(_ context.Context, _ *extauth.CheckRequest) (*extauth.CheckResponse, error) {
	// Do nothing but return ok
	return &extauth.CheckResponse{
		Status: &status.Status{},
		HttpResponse: &extauth.CheckResponse_OkResponse{
			OkResponse: &extauth.OkHttpResponse{},
		},
	}, nil
}
