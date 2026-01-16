package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(h *Handler) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/chats/", h.chat.CreateChat).Methods("POST")

	return r
}
