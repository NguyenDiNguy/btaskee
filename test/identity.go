package test

import (
	"btaskee/services/identity/proto"
	"context"
)

func SignUp(cli proto.IdentityClient, email, pass string) error {
	_, err := cli.SignUp(context.Background(), &proto.SignUpRequest{
		Email:    email,
		Password: pass,
	})
	return err
}

func SignIn(cli proto.IdentityClient, email, pass string) (*proto.SignInResponse, error) {
	rs, err := cli.SignIn(context.Background(), &proto.SignInRequest{
		Email:    email,
		Password: pass,
	})
	return rs, err
}
