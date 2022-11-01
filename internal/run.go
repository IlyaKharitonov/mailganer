package internal

import (
	"context"
	"fmt"
	"log"
	"net/http"

	gomail "gopkg.in/mail.v2"

	"mailganer/infrastructure/celery"
	"mailganer/internal/entities"
	"mailganer/internal/handlers"
	//"mailganer/internal/services/letterService"
	//"mailganer/internal/services/subscriberService"
)

func Run(config *entities.Config) error{
	ctx := context.Background()
	//подключение к постгрессу
	postgres, err := config.Postgres.ConnectPostgres(ctx)
	if err != nil {
		return fmt.Errorf("internal.Run() #1; Error: %s", err.Error())
	}
	log.Println("connected to pg")

	//подкл к celery
	celery, err := celery.Connect(config)
	if err != nil {
		return fmt.Errorf("internal.Run() #3; Error: %s", err.Error())
	}
	log.Println("connected to celery")

	//подключение к smtp серверу
	dialer := gomail.NewDialer(config.Smtp.Host, config.Smtp.Port, config.Smtp.Login, config.Smtp.Password)
	log.Println("connected to smtp")

	//регистрация хэндлеров и воркеров
	err = handlers.RegisterHandlers(postgres, celery, dialer, config)
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

