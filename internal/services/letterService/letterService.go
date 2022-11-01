package letterService

import (
	"context"
	"fmt"

	"mailganer/internal/entities"
)

type (
	storage interface {
		GetTemplate(id uint)(string, error)
	}

	taskQueue interface {
		Send(templateID uint, recipientsID []uint)error
	}

	letterService struct {
		storage storage
		tq taskQueue
		//ss *subscriberService.SubscriberService
	}

)

func NewLetterService(s storage, tq *celery) *letterService {
	return &letterService{s,tq}
}

func(ms *letterService)Send(ctx context.Context, letter *entities.Letter)error{

	err := ms.tq.Send(letter.TemplateID, letter.RecipientsID)
	if err != nil {
		return fmt.Errorf("(ms *letterService)Send #1 %s", err.Error())
	}

	return nil
}

func (ms *letterService)MarkAsRead(id uint)error{


}
