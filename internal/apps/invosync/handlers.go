package invosync

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/Okemwag/invosync/internal/pkg/model"
	"github.com/Okemwag/invosync/internal/pkg/service"
	"github.com/gorilla/mux"
)

const SuccessResponse = "success"
const ErrorResponse = "error"

type BookHandler struct {
	bookService service.BookService
}

func GetNewBookHandler(bookService service.BookService) *BookHandler {
	return &BookHandler{bookService: bookService}
}

func (bh *BookHandler) GetBookList(w http.ResponseWriter, r *http.Request) {
	books, err := bh.bookService.GetAllBooks()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, books)
}

func (bh *BookHandler) GetOrRemoveBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	isbn, err := strconv.Atoi(vars["isbn"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid book ISBN")
		return
	}
	switch r.Method {
	case http.MethodGet:
		book, err := bh.bookService.GetBook(isbn)
		if err != nil {
			respondWithError(w, http.StatusNotFound, err.Error())
			return
		}
		respondWithJSON(w, http.StatusOK, book)
	case http.MethodDelete:
		err := bh.bookService.DeleteBook(isbn)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJSON(w, http.StatusOK, SuccessResponse)
	}
}

func (bh *BookHandler) AddOrUpdateBookHandler(w http.ResponseWriter, r *http.Request) {
	book := &model.Book{}
	err := json.NewDecoder(r.Body).Decode(book)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	switch r.Method {
	case http.MethodPost:
		bh.bookService.AddBook(book)
		respondWithJSON(w, http.StatusCreated, SuccessResponse)
	case http.MethodPut:
		err := bh.bookService.UpdateBook(book)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJSON(w, http.StatusOK, SuccessResponse)
	}
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}