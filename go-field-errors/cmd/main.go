package main

import (
	"go-filed-errors/internals/handler"
	"go-filed-errors/internals/repository"
	"go-filed-errors/internals/service"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	repo := repository.NewPostRepository()
	service := service.NewPostservice(repo)
	handler := handler.NewPostHandler(service)

	r.Mount("/posts", handler.Routes())

	log.Println("server starting on port : 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
