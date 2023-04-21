package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)



type Handler struct {
	Router *mux.Router
	Service CommentService
	Server *http.Server
}

func NewHandler(service CommentService) *Handler {
	h := &Handler{
		Service: service,
	}

	h.Router = mux.NewRouter()
	h.mapRoutes()

	h.Server = &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: h.Router,
	}

	return h
}

func (h *Handler) mapRoutes()  {
	h.Router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	h.Router.HandleFunc("api/v1/comment", h.PostComment).Methods("POST")
	h.Router.HandleFunc("api/v1/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("api/v1/comment/{id}", h.UpdateComment).Methods("PUT")
	h.Router.HandleFunc("api/v1/comment/{id}", h.DeleteComment).Methods("DELETE")
}

func (h *Handler) Serve() error {
	if err := h.Server.ListenAndServe(); err != nil {
		log.Println(err.Error())
	}
	return nil
}