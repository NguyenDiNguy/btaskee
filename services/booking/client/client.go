package client

import (
	"btaskee/services/booking/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	proto.BookingClient
}

func NewBookingClient(url string) (proto.BookingClient, error) {
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient(url, options...)
	if err != nil {
		return nil, err
	}

	client := proto.NewBookingClient(conn)
	return &Client{BookingClient: client}, nil
}
