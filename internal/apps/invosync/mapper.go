package invosync

import "github.com/Okemwag/invosync/invosync/internal/pkg/model"


// Mapper is an interface that defines the methods that a struct must implement to be considered a Mapper

func DBBook(book *model.Book) *model.DBBook {
	return model.Book{
		Isbn:      book.Isbn,
		Name:      book.Name,
		Publisher: book.Publisher,
	}
}