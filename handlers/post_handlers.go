package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/adriangvaldes/blog-challenge/models"
)

var posts = []models.Post{
	{ID: 1, Title: "Meu Primeiro Post", Content: "Olá, mundo!", AuthorID: 1, CreatedAt: time.Now()},
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var newPost models.Post

	err := json.NewDecoder(r.Body).Decode(&newPost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	newPost.ID = int64(len(posts) + 1)
	newPost.CreatedAt = time.Now()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPost)
}
