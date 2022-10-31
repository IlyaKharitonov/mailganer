package subscriberHandlers

import "net/http"

func Register(h *handler){
	http.HandleFunc("/subscriber/getList", h.GetList)
	http.HandleFunc("/subscriber/add", h.Add)
}
