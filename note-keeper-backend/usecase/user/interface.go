package user

import "aedronef1_backend/entity"

type UserRepository interface {
	CreateUser(user *entity.User) error
	FindByUsername(username string) (*entity.User, error)
}

type StatusRepository interface {
	FindStatusByNameType(name string, typeStatus string) (*entity.Status, error)
}

type UseCase interface {
	SignUp(user *entity.User) error
	Login(username string, password string) (*entity.TokenPair, *entity.User, error)
}
