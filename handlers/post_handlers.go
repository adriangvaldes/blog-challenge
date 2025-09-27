package handlers

import (
	"net/http"
	"time"

	"github.com/adriangvaldes/blog-challenge/models"
)

var posts = []models.Post{
	{ID: 1, Title: "Meu Primeiro Post", Content: "Olá, mundo!", AuthorID: 1, CreatedAt: time.Now()},
}

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	// ...
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	// ...
}

func createPost(w http.ResponseWriter, r *http.Request) {
	// ...
}
