package service

import (
	"errors"

	"gorm.io/gorm"

	"github.com/nguyenbt456/demo-clean-architecture/model"
)

type BookService struct {
	bookRepo     model.BookRepository
	userBookRepo model.UserBookRepository
	userRepo     model.UserRepository
}

func NewBookService(bRepo model.BookRepository, ubRepo model.UserBookRepository, uRepo model.UserRepository) model.BookService {
	return &BookService{
		bookRepo:     bRepo,
		userBookRepo: ubRepo,
		userRepo:     uRepo,
	}
}

func (s *BookService) GetByID(id string) (*model.Book, error) {
	book, err := s.bookRepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("BookID is invalid")
		}
		return nil, errors.New("Can't get book from database")
	}

	return book, nil
}

func (s *BookService) GetByUserID(userID string) ([]*model.Book, error) {
	if userID == "" {
		return nil, errors.New("UserID is null")
	}

	_, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}

	userBooks, err := s.userBookRepo.GetByUserID(userID)
	if err != nil {
		return nil, errors.New("Can't get user_book from database")
	}

	books := []*model.Book{}
	for _, v := range userBooks {
		book, err := s.bookRepo.GetByID(v.BookID)
		if err != nil {
			return nil, errors.New("Can't get book from database")
		}

		books = append(books, book)
	}

	return books, nil
}
