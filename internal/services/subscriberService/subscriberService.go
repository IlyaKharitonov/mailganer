package subscriberService

import (
	"mailganer/internal/entities"
)

type (

	storage interface {
		GetList(IDs[]uint)([]*entities.Subscriber, error)
		Add([]*entities.Subscriber)error
	}

	SubscriberService struct {
		storage storage
	}
)

func NewSubscriberService(s *postgres) *SubscriberService {
	return &SubscriberService{storage: s}
}

func (ss *SubscriberService)GetList()([]*entities.Subscriber,error) {
	return nil, nil
}

func (ss *SubscriberService)Add([]*entities.Subscriber)error{
	return nil
}

func(ss *SubscriberService)GetStorage()storage{
	return ss.storage
}
