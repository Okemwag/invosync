package service

import (
	"fmt"
	"github.com/Okemwag/invosync/internal/apps/invosync" 
	"github.com/Okemwag/invosync/internal/pkg/model"
	"github.com/Okemwag/invosync/internal/pkg/repo"
	"github.com/pkg/errors"
)

// BookService is responsible for managing book resources
type BookService struct {
	booksRepo *repo.BookRepo
}

func GetNewBookService(bookRepo *repo.BookRepo) BookService {
	return BookService{booksRepo: bookRepo}
}

func (bs *BookService) AddBook(book *model.Book) {
	dbBook := invosync.DBBook(book)
	bs.booksRepo.AddBook(dbBook)

}

func (bs *BookService) GetBook(isbn int) (*model.Book, error) {
	dbBook := bs.booksRepo.GetBook(isbn)
	if dbBook != nil {
		book := invosync.Book(dbBook)
		return book, nil
	}
	return nil, errors.New(fmt.Sprintf("book with isbn %d not found", isbn))
}

func (bs *BookService) UpdateBook(book *model.Book) error {
	dbBook := invosync.DBBook(book)
	err := bs.booksRepo.UpdateBook(dbBook)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to update book with isbn %d", book.Isbn))
	}
	return nil
}

func (bs *BookService) GetAllBooks() ([]*model.Book, error) {
	dbBooks, err := bs.booksRepo.GetAllBooks()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get all books")
	}
	books := make([]*model.Book, 0)
	for _, dbBook := range dbBooks {
		book := invosync.Book(dbBook)
		books = append(books, book)
	}
	return books, nil
}

func (bs *BookService) DeleteBook(isbn int) error {
	err := bs.booksRepo.DeleteBook(isbn)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to delete book with isbn %d", isbn))
	}
	return nil
}