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

func TestInsertTask(t *testing.T) {
	repo := Init()

	id, err := repo.InsertTask(context.Background(), "abc", "tasker")
	if err != nil {
		panic(err)
	}

	task, err := repo.GetTaskById(context.Background(), *id)
	log.Info().Msgf("%v", task)
}
