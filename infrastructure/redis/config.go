package redis

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	redigo "github.com/gomodule/redigo/redis"
)

type Config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
	NameDB   int    `json:"name"`
}


func (r *Config) CreatePool()*redigo.Pool{
	return &redigo.Pool{
		MaxIdle:     3,
		MaxActive:   0,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redigo.Conn, error) {
			c, err := redigo.DialURL( fmt.Sprintf("redis://%s:%s", r.Host, r.Port))
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redigo.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}


func (r *Config) Connect(ctx context.Context) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", r.Host, r.Port),
		Password: r.Password,
		DB:       r.NameDB,
	})

	pong, err := client.Ping(ctx).Result()
	//if err != nil{
	//	return nil, err
	//}
	if err == nil && pong != "PONG"{
		return nil, errors.New("No connect redis")
	}

	return client, nil
}