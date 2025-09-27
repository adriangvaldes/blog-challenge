package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/adriangvaldes/blog-challenge/models"
	"github.com/gorilla/mux"
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

	posts = append(posts, newPost)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPost)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, post := range posts {
		if post.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(post)
			return
		}
	}

	http.Error(w, "Post não encontrado", http.StatusNotFound)
}
