package celery

import (
	"fmt"
	"log"

	"github.com/gocelery/gocelery"

	"mailganer/internal/entities"
)

func Connect(config *entities.Config)(*gocelery.CeleryClient, error){
	redisPool := config.Redis.CreatePool()
	log.Println("connected to redis")

	celery, err := gocelery.NewCeleryClient(
		gocelery.NewRedisBroker(redisPool),
		&gocelery.RedisCeleryBackend{Pool: redisPool},
		1,
	)
	if err != nil {
		return nil, fmt.Errorf("internal.Connect() #1; Error: %s", err.Error())
	}

	return celery,nil
}
