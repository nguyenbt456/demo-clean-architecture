package repository

import (
	"github.com/nguyenbt456/demo-clean-architecture/model"
	"gorm.io/gorm"
)

type UserBookRepo struct {
	db *gorm.DB
}

func NewUserBookRepository(db *gorm.DB) model.UserBookRepository {
	return &UserBookRepo{db}
}

func (r *UserBookRepo) GetByUserID(userID string) ([]*model.UserBook, error) {
	userBooks := []*model.UserBook{}

	if err := r.db.Where("user_id = ?", userID).Find(&userBooks).Error; err != nil {
		return nil, err
	}

	return userBooks, nil
}
