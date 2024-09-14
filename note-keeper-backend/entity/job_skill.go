package entity

import "time"

type JobSkill struct {
	JobId   int `gorm:"primaryKey"`
	SkillId int `gorm:"primaryKey"`

	Skill *Skill

	CreatedAt time.Time
	UpdatedAt time.Time
}
