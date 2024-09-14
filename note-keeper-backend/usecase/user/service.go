package user

import (
	"aedronef1_backend/entity"
	"aedronef1_backend/infrastucture/repository/define"

	jwtUtil "aedronef1_backend/infrastucture/repository/util"
)

type Service struct {
	userRepo   UserRepository
	statusRepo StatusRepository
}

func NewService(userRepo UserRepository, statusRepo StatusRepository) *Service {
	return &Service{
		userRepo:   userRepo,
		statusRepo: statusRepo,
	}
}

func (s *Service) SignUp(user *entity.User) error {
	err := user.HashPassword()
	if err != nil {
		return err
	}

	userDefault, err := s.statusRepo.FindStatusByNameType(define.ACTIVE, define.TYPE_USER)
	if err != nil {
		return err
	}

	user.StatusId = userDefault.Id

	err = s.userRepo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Login(username string, password string) (*entity.TokenPair, *entity.User, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return nil, nil, entity.ErrInternalServerError
	}

	if user == nil {
		return nil, nil, entity.ErrUsernameNotExists
	}

	validatePassword := user.ValidatePassword(password)
	if !validatePassword {
		return nil, nil, entity.ErrInvalidPassword
	}

	token, err := jwtUtil.GenerateAccessToken(user)
	if err != nil {
		return nil, nil, entity.ErrInternalServerError
	}

	refreshToken, err := jwtUtil.GenerateRefreshToken(user)
	if err != nil {
		return nil, nil, entity.ErrInternalServerError
	}

	tokenPair := entity.TokenPair{
		Token:        token,
		RefreshToken: refreshToken,
	}

	return &tokenPair, user, nil
}
