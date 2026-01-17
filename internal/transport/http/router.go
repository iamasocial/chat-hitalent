package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(h *Handler) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/chats/", h.chat.CreateChat).Methods("POST")
	r.HandleFunc("/chats/{id:[0-9]+}", h.chat.GetChatByID).Methods("GET")
	r.HandleFunc("/chats/{id:[0-9]+}", h.chat.Delete).Methods("DELETE")

	r.HandleFunc("/chats/{id:[0-9]+}/messages/", h.message.SendMessage).Methods("POST")

	return r
}
