package redis

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type IRedis interface {
	AcquireLock(ctx context.Context, key, value string, expiration time.Duration) (bool, error)
	ReleaseLock(ctx context.Context, key string, value string) error
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}) error
	Add(ctx context.Context, key string, values ...interface{}) error
	CheckList(ctx context.Context, key string, value interface{}) (bool, error)
}

type myRedis struct {
	client *redis.Client
}

func Initialize(address, pass string) IRedis {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     address, // Địa chỉ Redis server
		Password: pass,    // Mặc định không có password
		DB:       0,       // Sử dụng DB số 0
	})

	// Kiểm tra kết nối
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Không thể kết nối đến Redis: %v", err)
	}
	fmt.Println("Kết nối Redis thành công!")

	return &myRedis{
		client: redisClient,
	}
}

func (re *myRedis) AcquireLock(ctx context.Context, key, value string, expiration time.Duration) (bool, error) {
	// Sử dụng SetNX để kiểm tra và đặt khóa
	return re.client.SetNX(ctx, key, value, expiration).Result()
}

func (re *myRedis) ReleaseLock(ctx context.Context, key string, value string) error {
	// Giải phóng khóa
	val, err := re.client.Get(ctx, key).Result()
	if err == nil && val == value {
		re.client.Del(ctx, key)
	}

	return err
}

func (re *myRedis) Get(ctx context.Context, key string) (string, error) {
	return re.client.Get(ctx, key).Result()
}

func (re *myRedis) Set(ctx context.Context, key string, value interface{}) error {
	rs := re.client.Set(ctx, key, value, 10*time.Minute)
	return rs.Err()
}

func (re *myRedis) Add(ctx context.Context, key string, values ...interface{}) error {
	rs := re.client.SAdd(ctx, key, values...)
	return rs.Err()
}

func (re *myRedis) CheckList(ctx context.Context, key string, value interface{}) (bool, error) {
	return re.client.SIsMember(ctx, key, value).Result()
}
