package services

import (
	"fmt"
	"reflect"

	"github.com/go-redis/redis/v8"
)

func RedisDB() *redis.Client{
	client := redis.NewClient(&redis.Options{
        Addr:     "redis:6379",
        Password: "example",
        DB:       0,
    })
    fmt.Println(reflect.ValueOf(client))
	// return *client

    return client
}