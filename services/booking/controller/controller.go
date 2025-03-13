package controller

import (
	"btaskee/libs/jwt"
	"btaskee/services/booking/proto"
	"btaskee/services/booking/service"
	"context"

	"github.com/rs/zerolog/log"
)

type controller struct {
	proto.UnimplementedBookingServer
	svc service.IService
}

func NewController(svc service.IService) proto.BookingServer {
	return &controller{
		svc: svc,
	}
}

func (c *controller) CreateTask(ctx context.Context, req *proto.CreateTaskRequest) (*proto.CreateTaskResponse, error) {
	log.Info().Msg("CreateTask")
	token, err := jwt.Claim(ctx, "authorization")
	if err != nil {
		return nil, err
	}

	userId, err := jwt.ValidateJWT(token)
	if err != nil {
		return nil, err
	}

	return c.svc.CreateTask(ctx, userId, req)
}

func (c *controller) GetTask(ctx context.Context, req *proto.GetTaskRequest) (*proto.GetTaskResponse, error) {
	return c.svc.GetTask(ctx, req)
}

func (c *controller) AcceptTask(ctx context.Context, req *proto.AcceptTaskRequest) (*proto.AcceptTaskResponse, error) {
	log.Info().Msg("AcceptTask")
	token, err := jwt.Claim(ctx, "authorization")
	if err != nil {
		return nil, err
	}

	userId, err := jwt.ValidateJWT(token)
	if err != nil {
		return nil, err
	}

	return c.svc.AcceptTask(ctx, userId, req)
}

func (c *controller) ConfirmTasker(ctx context.Context, req *proto.ConfirmTaskerRequest) (*proto.ConfirmTaskerResponse, error) {
	log.Info().Msg("ConfirmTasker")
	token, err := jwt.Claim(ctx, "authorization")
	if err != nil {
		return nil, err
	}

	userId, err := jwt.ValidateJWT(token)
	if err != nil {
		return nil, err
	}

	return c.svc.ConfirmTasker(ctx, userId, req)
}
