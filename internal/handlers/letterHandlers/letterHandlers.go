package letterHandlers

import "net/http"

type LetterService interface {

}

type handler struct {
	service LetterService
}

func NewHandler(service LetterService) *handler {
	return &handler{service: service}
}

func(h *handler)Send(w http.ResponseWriter, req *http.Request){

}





