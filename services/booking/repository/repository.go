package repository

import (
	"btaskee/libs/mongodb"
	"btaskee/model/task"
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type IRepository interface {
	InsertTask(ctx context.Context, detail, asker string) (*bson.ObjectID, error)
	GetTaskById(ctx context.Context, id bson.ObjectID) (*task.Task, error)
	GetTaskByUuid(ctx context.Context, uuid string) (*task.Task, error)
	UpdateTask(ctx context.Context, task *task.Task) error
}

type repository struct {
	db mongodb.IMongoDb
}

func NewRepository(db mongodb.IMongoDb) IRepository {
	return &repository{
		db: db,
	}
}

func (repo *repository) InsertTask(ctx context.Context, detail, asker string) (*bson.ObjectID, error) {
	rs, err := repo.db.InsertOne(ctx, mongodb.TaskCollection, task.Task{
		Uuid:      uuid.NewString(),
		Detail:    detail,
		AskerId:   asker,
		Status:    task.TaskStatus_PENDING,
		CreatedAt: int64(time.Now().Unix()),
	})
	if err != nil {
		return nil, err
	}

	id := rs.InsertedID.(bson.ObjectID)
	return &id, nil
}

func (repo *repository) UpdateTask(ctx context.Context, task *task.Task) error {
	rs, err := repo.db.UpdateOne(ctx, mongodb.TaskCollection, bson.M{"uuid": task.Uuid}, bson.M{"$set": task})
	if err != nil {
		log.Err(err).Send()
		return err
	}

	log.Debug().Msgf("%v", rs)
	if rs.MatchedCount == 0 {
		return errors.New("Data not found")
	}

	return nil
}

func (repo *repository) GetTaskById(ctx context.Context, id bson.ObjectID) (*task.Task, error) {
	result := &task.Task{}
	rs, err := repo.db.GetOne(ctx, mongodb.TaskCollection, bson.M{
		"_id": id,
	})
	if err != nil {
		return nil, err
	}

	err = rs.Decode(result)
	if err != nil {
		log.Err(err).Send()
		return nil, err
	}

	return result, nil
}

func (repo *repository) GetTaskByUuid(ctx context.Context, uuid string) (*task.Task, error) {
	result := &task.Task{}
	rs, err := repo.db.GetOne(ctx, mongodb.TaskCollection, bson.M{
		"uuid": uuid,
	})
	if err != nil {
		return nil, err
	}

	err = rs.Decode(result)
	if err != nil {
		log.Err(err).Send()
		return nil, err
	}

	return result, nil
}
