package subscriberHandlers

import "net/http"

type SubscribeService interface {

}

type handler struct {
	service SubscribeService
}

func NewHandler(service SubscribeService) *handler {
	return &handler{service: service}
}

func(h *handler)GetList(w http.ResponseWriter, req *http.Request){

}

func(h *handler)Add(w http.ResponseWriter, req *http.Request){

}