package repository

import (
	"aedronef1_backend/entity"
	"aedronef1_backend/infrastucture/repository/define"
	"errors"

	"gorm.io/gorm"
)

type StatusRepository struct {
	db *gorm.DB
}

func NewStatusRepository(db *gorm.DB) *StatusRepository {
	return &StatusRepository{
		db: db,
	}
}

func (r *StatusRepository) FindStatusByNameType(name string, typeStatus string) (*entity.Status, error) {
	status := &entity.Status{}
	err := r.db.Where("name = ? AND type = ?", name, typeStatus).Find(&status).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(define.DOCUMENT_NOT_FOUND)
		} else {
			return nil, err
		}
	}

	return status, nil
}
