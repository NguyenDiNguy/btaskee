package main

import (
	"btaskee/services/identity/client"
	"btaskee/services/identity/proto"
	"context"
	"fmt"
	"testing"
)

func TestSignUp(t *testing.T) {
	cli, err := client.NewIdentityClient("127.0.0.1:9080")
	if err != nil {
		panic(err)
	}
	rs, err := cli.SignUp(context.Background(), &proto.SignUpRequest{
		Email:    "kennynguyen2080@gmail.com",
		Password: "abcd1234",
	})
	if err != nil {
		t.Error(err)
	}

	fmt.Println(rs)
}

func TestSignIn(t *testing.T) {
	cli, err := client.NewIdentityClient("127.0.0.1:9080")
	if err != nil {
		panic(err)
	}
	rs, err := cli.SignIn(context.Background(), &proto.SignInRequest{
		Email:    "kennynguyen2080@gmail.com",
		Password: "abcd1234",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(rs)
}
