package repository

import (
	"btaskee/libs/mongodb"
	"btaskee/model/user"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
)

type IRepository interface {
	InsertUser(ctx context.Context, email, hashPassword string) error
	GetUserByEmail(ctx context.Context, email string) (*user.User, error)
}

type repository struct {
	db mongodb.IMongoDb
}

func NewRepository(db mongodb.IMongoDb) IRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) InsertUser(ctx context.Context, email, hashPassword string) error {
	_, err := r.db.InsertOne(ctx, mongodb.UserCollection, user.User{
		Uuid:         uuid.NewString(),
		Email:        email,
		HashPassword: hashPassword,
		CreatedAt:    int64(time.Now().Unix()),
	})
	if err != nil {
		log.Err(err).Send()
		return err
	}

	return nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*user.User, error) {
	result := &user.User{}
	rs, err := r.db.GetOne(ctx, mongodb.UserCollection, bson.M{
		"email": email,
	})
	if err != nil {
		log.Err(err).Send()
		return nil, err
	}

	err = rs.Decode(result)
	if err != nil {
		log.Err(err).Send()
		return nil, err
	}

	return result, nil
}
