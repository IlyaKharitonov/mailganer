package internal

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gocelery/gocelery"

	"mailganer/internal/entities"
	"mailganer/internal/handlers"
	"mailganer/internal/services/letterService"
	"mailganer/internal/services/subscriberService"
)

func Run(config *entities.Config) error{
	ctx := context.Background()
	//подключения к постгрессу, редису или рэббиту
	postgres, err := config.Postgres.ConnectPostgres(ctx)
	if err != nil {
		return fmt.Errorf("internal.Run() #1; Error: %s", err.Error())
	}
	log.Println("connected to pg")

	redisPool := config.Redis.CreatePool()
	log.Println("connected to redis")

	celery, err := gocelery.NewCeleryClient(
		gocelery.New)

	celery, err := gocelery.NewCeleryClient(
		gocelery.NewRedisBroker(redisPool),
		&gocelery.RedisCeleryBackend{Pool: redisPool},
		1,
	)
	if err != nil {
		return fmt.Errorf("internal.Run() #3; Error: %s", err.Error())
	}
	log.Println("connected to celery")

	err := handlers.RegisterHandlers(postgres, redis, celery)
	if err != nil {
		return fmt.Errorf("internal.Run() #4; Error: %s", err.Error())
	}
	log.Println("all handlers registered")

	fmt.Println("server started")
	err = http.ListenAndServe(fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port),nil)
	if err != nil{
		return fmt.Errorf("internal.Run() #5; Error: %s", err.Error())
	}
	return nil
}

