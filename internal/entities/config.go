package entities

import (
	"encoding/json"
	"io/ioutil"

	"mailganer/infrastructure/pg"
	"mailganer/infrastructure/rabbit"
	"mailganer/infrastructure/redis"
	"mailganer/infrastructure/smtp"
)

type Config struct {
	Server     Server             `json:"server"`
	Postgres   pg.Config		  `json:"postgres"`
	Redis 	   redis.Config       `json:"redis"`
	Rabbit     rabbit.Config 	  `json:"nats"`
	Smtp  	   smtp.Config 	      `json:"smtp"`
}

type Server struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func(c *Config)Parse(pathConfig string)error{
	file, err := ioutil.ReadFile(pathConfig)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &c)
}