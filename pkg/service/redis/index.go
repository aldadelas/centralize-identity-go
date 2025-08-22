package redis_service

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisService struct {
	Client *redis.Client
}

var (
	instance *RedisService
	once     sync.Once
)

func NewRedisService() *RedisService {
	once.Do(func() {
		if os.Getenv("REDIS_HOST") == "" {
			log.Fatal("REDIS_HOST is not defined")
		}

		host := os.Getenv("REDIS_HOST")
		port := "6379"
		password := os.Getenv("REDIS_PASSWORD")

		if os.Getenv("REDIS_PORT") != "" {
			port = os.Getenv("REDIS_PORT")
		}

		instance = &RedisService{
			Client: redis.NewClient(&redis.Options{
				Addr:     fmt.Sprintf("%s:%s", host, port),
				Password: password,
				DB:       0,
			}),
		}

	})

	return instance
}

func (s *RedisService) SetKey(key string, value string, ttl time.Duration) error {
	result := s.Client.Set(context.Background(), key, value, ttl)

	if result.Err() != nil {
		return result.Err()
	}

	return nil
}

func (s *RedisService) GetKey(key string) (string, error) {
	result := s.Client.Get(context.Background(), key)

	if result.Err() != nil {
		return "", result.Err()
	}

	return result.Val(), nil
}

func (s *RedisService) DeleteKey(key string) error {
	result := s.Client.Del(context.Background(), key)

	if result.Err() != nil {
		return result.Err()
	}

	return nil
}
