package test

import (
	"btaskee/services/booking/proto"
	"context"

	"google.golang.org/grpc/metadata"
)

func CreateTask(cli proto.BookingClient, detail, jwt string) (*proto.CreateTaskResponse, error) {
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("authorization", jwt))
	return cli.CreateTask(ctx, &proto.CreateTaskRequest{
		Detail: detail,
	})
}

func AcceptTask(cli proto.BookingClient, taskId, jwt string) (*proto.AcceptTaskResponse, error) {
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("authorization", jwt))
	return cli.AcceptTask(ctx, &proto.AcceptTaskRequest{
		TaskId: taskId,
	})
}

func ConfirmTasker(cli proto.BookingClient, taskId, taskerId, jwt string) (*proto.ConfirmTaskerResponse, error) {
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("authorization", jwt))
	return cli.ConfirmTasker(ctx, &proto.ConfirmTaskerRequest{
		TaskId:   taskId,
		TaskerId: taskerId,
	})
}
