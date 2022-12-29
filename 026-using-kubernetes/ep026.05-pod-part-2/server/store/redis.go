package store

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type RedisStore struct {
	client *redis.Client
}

func NewRedisStore(redisURL string) *RedisStore {
	rs := new(RedisStore)

	rdbConfig := &redis.Options{
		Addr:     redisURL,
		Password: "",
		DB:       0,
	}

	client := redis.NewClient(rdbConfig)
	rs.client = client

	return rs
}

func (rs *RedisStore) Ping() error {
	cmd := rs.client.Ping(context.Background())
	return cmd.Err()
}

func (rs *RedisStore) Put(key string, value []byte) error {
	cmd := rs.client.Set(context.Background(), key, value, 0)
	return cmd.Err()
}

func (rs *RedisStore) Get(key string) ([]byte, error) {
	cmd := rs.client.Get(context.Background(), key)
	return cmd.Bytes()
}
