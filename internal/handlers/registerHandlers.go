package handlers

import (
	"github.com/go-redis/redis"
	"github.com/gocelery/gocelery"
	"github.com/jmoiron/sqlx"

	"mailganer/internal/handlers/letterHandlers"
	"mailganer/internal/handlers/subscriberHandlers"
	"mailganer/internal/services/letterService"
	"mailganer/internal/services/subscriberService"
)

func RegisterHandlers(postgres *sqlx.DB, redis *redis.Client, celery *gocelery.CeleryClient)error {

	ss := subscriberService.NewSubscriberService(subscriberService.NewPg(postgres))
	sh := subscriberHandlers.NewHandler(ss)
	subscriberHandlers.Register(sh)

	ls := letterService.NewLetterService(letterService.NewPg(postgres), redis, celery, *ss)
	lh := letterHandlers.NewHandler(ls)
	letterHandlers.Register(lh)



	return nil
}