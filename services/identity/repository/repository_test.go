package repository

import (
	"btaskee/libs/logger"
	"btaskee/libs/mongodb"
	"context"
	"testing"

	"github.com/rs/zerolog/log"
)

func Init() IRepository {
	logger.Init()
	db := mongodb.Initialize("mongodb://admin:password@localhost:27017", "bTaskee")
	return NewRepository(db)
}

func TestInsertUser(t *testing.T) {
	repo := Init()

	err := repo.InsertUser(context.Background(), "test1@gmail.com", "abc")
	if err != nil {
		panic(err)
	}
}

func TestGetUserByEmail(t *testing.T) {
	repo := Init()

	user, err := repo.GetUserByEmail(context.Background(), "test1@gmail.com")
	if err != nil {
		panic(err)
	}
	log.Info().Msgf("%v", user)
}
