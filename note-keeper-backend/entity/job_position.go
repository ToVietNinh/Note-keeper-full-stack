package entity

import "time"

type JobPosition struct {
	JobId      int `gorm:"primaryKey"`
	PositionId int `gorm:"primaryKey"`

	Position *Position

	CreatedAt time.Time
	UpdatedAt time.Time
}
