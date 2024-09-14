package entity

import "time"

type Status struct {
	Id        int `gorm:"primaryKey"`
	Name      string
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
