package client

import (
	"btaskee/services/identity/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	proto.IdentityClient
}

func NewIdentityClient(url string) (proto.IdentityClient, error) {
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient(url, options...)
	if err != nil {
		return nil, err
	}

	client := proto.NewIdentityClient(conn)
	return &Client{IdentityClient: client}, nil
}
