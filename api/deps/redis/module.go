package redis

import (
	"os"
	"strconv"

	"flamingo.me/dingo"
	"github.com/redis/go-redis/v9"
)

type Module struct{}

func (*Module) Configure(injector *dingo.Injector) {
	injector.Bind(redis.NewClient(&redis.Options{})).ToProvider(func() (*redis.Client, error) {
		redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
		if err != nil {
			return &redis.Client{}, err
		}
		return redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
			Password: "",
			DB:       redisDB,
		}), nil

	}).In(dingo.Singleton)
}
