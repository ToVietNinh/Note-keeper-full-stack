package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        int `gorm:"primaryKey"`
	Username  string
	Password  string
	Email     string
	StatusId  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ClaimUser struct {
	UserId       int
	Username     string
	Token        string
	RefreshToken string
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) HashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return err
	}

	u.Password = string(hash)

	return nil
}
