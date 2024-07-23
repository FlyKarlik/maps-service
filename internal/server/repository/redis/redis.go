package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"maps-service/config"
)

func NewStyledMapsClient(ctx context.Context, config *config.Config) (*redis.Client, error) {
	r := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: "",
		DB:       config.RedisStyledMapsDB,
	})

	if err := r.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("r.Ping: %w", err)
	}
	return r, nil
}
