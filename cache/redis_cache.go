package cache

import (
	"context"
	"encoding/json"
	"redis/model"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisCache struct {
	host   string
	db     int
	expire time.Duration
	client *redis.Client
}

func NewRedisCache(
	host,
	password string,
	db int,
	expire time.Duration,
) Cache {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       db,
	})

	return &redisCache{
		host:   host,
		db:     db,
		expire: expire,
		client: redisClient,
	}
}

func (r *redisCache) Set(ctx context.Context, key string, human *model.Human) error {
	var err error
	humanByte, err := json.Marshal(human)
	if err != nil {
		return err
	}
	r.client.Set(ctx, key, humanByte, r.expire*time.Minute)
	return err
}
func (r *redisCache) Get(ctx context.Context, key string) (*model.Human, error) {
	var err error
	var human model.Human

	result, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return &human, err
	}
	err = json.Unmarshal([]byte(result), &human)
	if err != nil {
		return &human, err
	}
	return &human, err
}
