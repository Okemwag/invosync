package repo

import (
	"github.com/Okemwag/invosync/internal/pkg/model"
	"gorm.io/gorm"
)

//responsible for managing book resources
type BookRepo struct {
	db *gorm.DB
}

func GetNewBookRepo(db *gorm.DB) *BookRepo {
	return &BookRepo{db: db}
}

func (bk *BookRepo) AddBook(book *model.DBBook){
	bk.db.Create(book)
}

func (bk *BookRepo) UpdateBook(book *model.DBBook) error {
	bk.db.Model(&book).Where("isbn = ?", book.Isbn).Update("name", "publisher")
	return nil
}

func (bk *BookRepo) GetBook(isbn int) *model.DBBook {
	book := &model.DBBook{}
	bk.db.Where("isbn = ?", isbn).First(book)
	return book
}

func (bk *BookRepo) GetAllBooks() ([]*model.DBBook, error) {
	books := make([]*model.DBBook, 0)
	err := bk.db.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil	
}

func (bk *BookRepo) DeleteBook(isbn int) error {
	book := &model.DBBook{}
	bk.db.Where("isbn = ?", isbn).Delete(book)
	return nil
}