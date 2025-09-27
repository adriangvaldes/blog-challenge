package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adriangvaldes/blog-challenge/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/posts", handlers.CreatePost).Methods(http.MethodPost)
	r.HandleFunc("/posts", handlers.GetPosts).Methods(http.MethodGet)
	r.HandleFunc("/posts/{id}", handlers.GetPost).Methods(http.MethodGet)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "API do Blog está no ar!")
	})

	fmt.Println("Servidor escutando na porta 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
