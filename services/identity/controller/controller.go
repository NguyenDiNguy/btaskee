package controller

import (
	"btaskee/services/identity/proto"
	"btaskee/services/identity/service"
	"context"

	"github.com/rs/zerolog/log"
)

type controller struct {
	proto.UnimplementedIdentityServer
	svc service.IService
}

func NewController(svc service.IService) proto.IdentityServer {
	return &controller{
		svc: svc,
	}
}

func (c *controller) SignUp(ctx context.Context, req *proto.SignUpRequest) (*proto.SignUpResponse, error) {
	log.Info().Msg("SignUp")

	return c.svc.SignUp(ctx, req)
}

func (c *controller) SignIn(ctx context.Context, req *proto.SignInRequest) (*proto.SignInResponse, error) {
	log.Info().Msg("SignIn")

	return c.svc.SignIn(ctx, req)
}
