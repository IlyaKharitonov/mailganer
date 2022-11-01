package handlers

import (
	"github.com/gocelery/gocelery"
	"github.com/jmoiron/sqlx"
	"gopkg.in/mail.v2"

	"mailganer/internal/entities"
	"mailganer/internal/handlers/letterHandlers"
	"mailganer/internal/handlers/subscriberHandlers"
	"mailganer/internal/services/celeryWorkerService"
	"mailganer/internal/services/letterService"
	"mailganer/internal/services/subscriberService"
)

func RegisterHandlers(postgres *sqlx.DB, celery *gocelery.CeleryClient, dialer *mail.Dialer, config *entities.Config)error {

	ss := subscriberService.NewSubscriberService(subscriberService.NewPg(postgres))
	sh := subscriberHandlers.NewHandler(ss)
	subscriberHandlers.Register(sh)

	ls := letterService.NewLetterService(
		letterService.NewPg(postgres),
		letterService.NewCeleryCliService(celery))
	lh := letterHandlers.NewHandler(ls)
	letterHandlers.Register(lh)

	//пуск воркера в горутине
	err := celeryWorkerService.NewCeleryWorker(celery,ss, dialer, letterService.NewPg(postgres), config).Run()
	if err != nil {
		return err
	}

	return nil
}