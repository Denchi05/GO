package main

import (
	"cinema-friend/internal/handlers"
	"cinema-friend/internal/repository"
	"fmt"
	"log"
	"net/http"
)

func main() {
	repo := repository.NewStorage()
	handler := &handlers.Handler{Repo: repo}

	http.HandleFunc("/api/v1/movies", handler.GetMoviesHandler)
	http.HandleFunc("/api/v1/book", handler.BookSeatHandler)

	fmt.Println("🚀 Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
