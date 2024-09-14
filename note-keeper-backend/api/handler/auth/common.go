package handler

import (
	authPayload "aedronef1_backend/api/payload/auth"
	"aedronef1_backend/entity"
	"fmt"

	"aedronef1_backend/api/presenter"
)

func convertPayloadToUserEntity(payload *authPayload.SignUpForm) *entity.User {
	return &entity.User{
		Username: payload.Username,
		Password: payload.Password,
		Email:    payload.Email,
	}
}

func convertTokenPairUserToResponse(tokenPair *entity.TokenPair, user *entity.User) *presenter.AuthResult {
	response := &presenter.AuthResult{
		Token:        tokenPair.Token,
		RefreshToken: tokenPair.RefreshToken,
		UserId:       fmt.Sprint(user.Id),
		Username:     user.Username,
		Email:        user.Email,
	}

	return response
}
