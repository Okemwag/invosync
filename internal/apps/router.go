package handlers

import (
	"github.com/Okemwag/invosync/internal/pkg/service"
	"github.com/gorilla/mux"
)

// Router is responsible for routing requests to the appropriate handler

func ProvideRouter(bookService service.BookService) *mux.Router {
	router := mux.NewRouter()
	bookHandler := GetNewBookHandler(bookService)

	router.HandleFunc("/books", bookHandler.GetBookList).Methods("GET")
	router.HandleFunc("/books/{isbn:[0-9]+}", bookHandler.GetOrRemoveBookHandler).Methods("GET", "DELETE")
	router.HandleFunc("/books", bookHandler.AddOrUpdateBookHandler).Methods("POST", "PUT")

	return router
}

