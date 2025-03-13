package redis

import (
	"context"
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	redis := Initialize("172.26.48.1:6379", "")
	err := redis.Add(context.Background(), "abc", "a")
	if err != nil {
		panic(err)
	}

	is, err := redis.CheckList(context.Background(), "abc", "a")
	if err != nil {
		panic(err)
	}

	fmt.Println(is)
}
