package redis

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	instance *Redis
	once     sync.Once
)

type Redis struct {
	Client *redis.Client
}

func GetInstance() *Redis {
	once.Do(func() {
		instance = &Redis{}
		instance.init()
	})
	return instance
}

func (r *Redis) init() {
	r.Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("localhost:%s", os.Getenv("REDIS_PORT")), // Redis server address
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0, // Use default DB
	})

	// Test the connection
	ctx := context.Background()
	pong, err := r.Client.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Could not connect to Redis: ", err)
	}
	log.Println("Connected to Redis: ", pong)
}

func (r *Redis) Set(ctx context.Context, key string, value interface{}) error {
	return r.Client.Set(ctx, key, value, 0).Err()
}

func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}

func (r *Redis) AddToHash(ctx context.Context, hash string, field string, value string) error {
	return r.Client.HSet(ctx, hash, field, value).Err()
}

func (r *Redis) GetFromHash(ctx context.Context, hash string, field string) (string, error) {
	val, err := r.Client.HGet(ctx, hash, field).Result()

	if err == redis.Nil {
		return "", fmt.Errorf("field %s does not exist", field)
	} else if err != nil {
		return "", fmt.Errorf("failed to get field from HSET: %w", err)
	}

	return val, nil
}

func (r *Redis) GetAllFromHash(ctx context.Context, hash string) (map[string]string, error) {
	fields, err := r.Client.HGetAll(ctx, hash).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get all fields from HSET: %w", err)
	}

	return fields, nil
}

func (r *Redis) Close() error {
	return r.Client.Close()
}
