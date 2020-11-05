package service

import (
	"errors"

	"github.com/nguyenbt456/demo-clean-architecture/model"
	"gorm.io/gorm"
)

type UserService struct {
	repo model.UserRepository
}

func NewUserService(repo model.UserRepository) model.UserService {
	return &UserService{repo}
}

func (s *UserService) GetByID(id string) (*model.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("UserID is invalid")
		}
		return nil, errors.New("Can't get user from database")
	}

	return user, nil
}
