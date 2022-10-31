package letterHandlers

import "net/http"

func Register(h *handler){
	http.HandleFunc("/letter/send", h.Send)
}

