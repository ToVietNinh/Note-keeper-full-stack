package entity

import "time"

type Position struct {
	Id   int `gorm:"primaryKey"`
	Name string

	CreatedAt time.Time
	UpdatedAt time.Time
}
