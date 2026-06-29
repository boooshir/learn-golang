package handler

import (
	"encoding/json"
	"go-filed-errors/internals/model"
	"go-filed-errors/internals/service"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type PostHandler struct {
	service service.PostService
}

func NewPostHandler(s service.PostService) *PostHandler {
	return &PostHandler{service: s}
}

func (handler *PostHandler) Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", handler.CreatePost)
	return r
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post model.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("%s", &post)
	createPost, err := h.service.CreatePost(&post)
	if err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			errors := make(map[string]string)
			for _, fieldErr := range validationErrs {
				errors[fieldErr.Field()] = fieldErr.Error()
			}
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(model.ErrorResponse{Errors: errors})
		} else {
			http.Error(w, "Couldnt create a post", http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.PostResponse{Data: *createPost})
}
