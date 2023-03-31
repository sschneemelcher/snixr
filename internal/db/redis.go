package db

import "github.com/redis/go-redis/v9"

func SetupDB() *redis.Client {
    // Setup redis connection
    rdb := redis.NewClient(&redis.Options{
        Addr:     "redis:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    return rdb
}
