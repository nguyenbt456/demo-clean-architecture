package repository

import (
	"github.com/nguyenbt456/demo-clean-architecture/model"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) model.UserRepository {
	return &UserRepo{db}
}

func (r *UserRepo) GetByID(id string) (*model.User, error) {
	user := &model.User{}

	if err := r.db.Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
