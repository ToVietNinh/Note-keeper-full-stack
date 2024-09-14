package entity

import "time"

type Skill struct {
	Id   int `gorm:"primaryKey"`
	Name string

	CreatedAt time.Time
	UpdatedAt time.Time
}
