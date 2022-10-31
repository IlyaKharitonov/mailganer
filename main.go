package main

import (
	"flag"
	"log"

	"mailganer/infrastructure/smtp"
	"mailganer/internal"
	"mailganer/internal/entities"
)

func main(){
	var (
		logMail = flag.String("l", "", "login for send email")
		passMail = flag.String("p", "", "password for send email")
		pathConfig = flag.String("config", "config/config.json", "")
		config = &entities.Config{Smtp: smtp.Config{Login: *logMail, Password: *passMail } }
	)
	flag.Parse()

	if err := config.Parse(*pathConfig); err != nil {
		log.Fatalf("main config.Parse() error: %s", err.Error())
	}

	if err := internal.Run(config); err != nil {
		log.Fatalf("main internal.StartServer() error: %s", err.Error())
	}
}
