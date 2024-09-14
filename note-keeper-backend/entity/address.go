package entity

import "time"

type Address struct {
	Id       int `gorm:"primaryKey"`
	No       int
	Road     string
	Ward     string
	District string
	City     string
	Country  string
	JobId    *int

	CreatedAt time.Time
	UpdatedAt time.Time
}

type ListJobIdByAddress struct {
	JobId int
}
