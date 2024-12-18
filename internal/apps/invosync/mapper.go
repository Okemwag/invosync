package invosync

import "github.com/Okemwag/invosync/internal/pkg/model"

// DBBook maps a model.Book to a model.DBBook
func DBBook(book *model.Book) *model.DBBook {
	dbBook := &model.DBBook{
		Isbn: book.Isbn,
		Name:  book.Name,
		Publisher: book.Publisher,
	}
	return dbBook
}

func Book(dbBook *model.DBBook) *model.Book {
	book := &model.Book{
		Isbn: dbBook.Isbn,
		Name: dbBook.Name,
		Publisher: dbBook.Publisher,
	}
	return book
}