package repository

import (
	"github.com/nguyenbt456/demo-clean-architecture/model"
	"gorm.io/gorm"
)

type BookRepo struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) model.BookRepository {
	return &BookRepo{db}
}

func (r *BookRepo) GetByID(id string) (*model.Book, error) {
	book := &model.Book{}

	if err := r.db.Where("id = ?", id).First(book).Error; err != nil {
		return nil, err
	}

	return book, nil
}
