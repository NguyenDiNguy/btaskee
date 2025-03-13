package service

import (
	"btaskee/libs/jwt"
	"btaskee/services/identity/proto"
	"btaskee/services/identity/repository"
	"context"

	"golang.org/x/crypto/bcrypt"
)

type IService interface {
	SignUp(ctx context.Context, req *proto.SignUpRequest) (*proto.SignUpResponse, error)
	SignIn(ctx context.Context, req *proto.SignInRequest) (*proto.SignInResponse, error)
}

type service struct {
	repo repository.IRepository
}

func NewService(repo repository.IRepository) IService {
	return &service{
		repo: repo,
	}
}

// CreateTask ...
func (s *service) SignUp(ctx context.Context, req *proto.SignUpRequest) (*proto.SignUpResponse, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	err = s.repo.InsertUser(ctx, req.Email, string(hashPassword))
	if err != nil {
		return nil, err
	}

	return &proto.SignUpResponse{
		IsSuccess: true,
	}, nil
}

// GetTask ...
func (s *service) SignIn(ctx context.Context, req *proto.SignInRequest) (*proto.SignInResponse, error) {
	user, err := s.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashPassword), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	token, err := jwt.GenerateJWT(user.Uuid)
	if err != nil {
		return nil, err
	}

	user.HashPassword = ""

	return &proto.SignInResponse{
		User: user,
		Jwt:  token,
	}, nil
}
