package rdb

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type Repository interface {
	Set(ctx context.Context, key, value string) error
	SetEX(ctx context.Context, key, value string, duration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) error
}

type RedisRepo struct {
	rds *redis.Client
}

func NewRedisRepo(rds *redis.Client) *RedisRepo {
	return &RedisRepo{
		rds: rds,
	}

}
func (r *RedisRepo) Set(ctx context.Context, key, value string) error {

	err := r.rds.Set(ctx, key, value, 0).Err()
	return err
}

// SetWithTTL

func (r *RedisRepo) Get(ctx context.Context, key string) (string, error) {

	val, err := r.rds.Get(ctx, key).Result()
	return val, err
}

func (r *RedisRepo) SetEX(ctx context.Context, key, value string, duration time.Duration) error {

	err := r.rds.SetEX(ctx, key, value, duration).Err()
	return err
}

func (r *RedisRepo) Del(ctx context.Context, key string) error {

	err := r.rds.Del(ctx, key).Err()
	return err
}
