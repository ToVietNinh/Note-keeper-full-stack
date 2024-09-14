package repository

import (
	"aedronef1_backend/entity"
	"aedronef1_backend/infrastucture/repository/define"
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(user *entity.User) error {
	err := r.db.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) FindByUsername(username string) (*entity.User, error) {
	user := &entity.User{}
	result := r.db.Where("username = ?", username).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New(define.DOCUMENT_NOT_FOUND)
	}

	return user, nil
}
