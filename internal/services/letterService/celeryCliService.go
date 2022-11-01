package letterService

import (
	"fmt"

	"github.com/gocelery/gocelery"
)

type celery struct {
	cli *gocelery.CeleryClient
}

func NewCeleryCliService(cli *gocelery.CeleryClient)*celery{
	return &celery{cli}
}

func(c *celery)Sand(templateID uint, recipientsID []uint)error{

	_, err := c.cli.Delay("Send", templateID, recipientsID)
	if err != nil{
		return fmt.Errorf("(c *celery)Sand #1; Error: %s", err.Error())
	}

	return nil
}