package entity

import "time"

type Location struct {
	Id       int `gorm:"primaryKey"`
	No       int
	Road     string
	Ward     string
	District string
	City     string
	Country  string

	CreatedAt time.Time
	UpdatedAt time.Time
}
