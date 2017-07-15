package redisClient

import (
    "log"

    "github.com/go-redis/redis"
)

func failOnError(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func NewClient() (*redis.Client) {
    client := redis.NewClient(&redis.Options{
        Addr:       "localhost:6379",
        Password:   "",
        DB:         0,
    })
    _, err := client.Ping().Result()
    failOnError(err)
    return client;
}
