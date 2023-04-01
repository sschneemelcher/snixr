package db

import (
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

func SetupDB() *redis.Client {
    // Setup redis connection
    redisAddr := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))

    rdb := redis.NewClient(&redis.Options{
        Addr:     redisAddr,
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    return rdb
}
