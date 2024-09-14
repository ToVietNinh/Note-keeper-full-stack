package repository

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func ConnectRedis(redisHost, redisPort string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%v:%v", redisHost, redisPort),
	})

	_, err := client.Ping(context.TODO()).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
