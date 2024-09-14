package entity

import "time"

type Job struct {
	Id          int `gorm:"primaryKey"`
	Title       string
	Description string
	Address     []*Address `gorm:"foreignKey:JobId"`

	CompanyId *int
	Company   *Company

	JobPosition []*JobPosition `gorm:"foreignKey:JobId"`
	JobSkill    []*JobSkill    `gorm:"foreignKey:JobId"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
