package handler

import (
	"aedronef1_backend/api/presenter"
	"aedronef1_backend/config"
	"aedronef1_backend/entity"
	"aedronef1_backend/infrastucture/repository/util"
	"aedronef1_backend/usecase/user"
	"fmt"
	"net/http"

	authPayload "aedronef1_backend/api/payload/auth"

	jwtUtil "aedronef1_backend/infrastucture/repository/util"

	"github.com/gin-gonic/gin"
)

// @Summary Login
// @Schemes
// @Description login
// @Tags auth
// @Param loginForm body authPayload.LoginForm true "Login"
// @Success 200 {object} presenter.LoginResponse
// @Failure 400 {object} presenter.Error400Response
// @Failure 404 {object} presenter.Error404Response
// @Failure 500 {object} presenter.Error500Response
// @Router /auth/login [post]
func login(ctx *gin.Context, userService user.UseCase) {
	var payload *authPayload.LoginForm
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		util.HandlerException(ctx, http.StatusBadRequest, entity.ErrBadRequest)
		return
	}

	tokenPair, user, err := userService.Login(payload.Username, payload.Password)
	if err != nil {
		switch err {
		case entity.ErrUsernameNotExists:
			util.HandlerException(ctx, http.StatusBadRequest, err)
			return

		case entity.ErrInvalidPassword:
			util.HandlerException(ctx, http.StatusBadRequest, err)
			return

		default:
			util.HandlerException(ctx, http.StatusInternalServerError, err)
			return
		}
	}

	response := presenter.LoginResponse{
		Status:  fmt.Sprint(http.StatusOK),
		Message: config.SUCCESS,
		Results: convertTokenPairUserToResponse(tokenPair, user),
	}
	ctx.JSON(http.StatusOK, response)
}

func signUp(ctx *gin.Context, userService user.UseCase) {
	var payload *authPayload.SignUpForm
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		util.HandlerException(ctx, http.StatusBadRequest, entity.ErrBadRequest)
		return
	}

	userEntity := convertPayloadToUserEntity(payload)

	err = userService.SignUp(userEntity)

	response := presenter.BasicResponse{
		Status:  fmt.Sprint(http.StatusOK),
		Message: config.SUCCESS,
		Results: nil,
	}
	ctx.JSON(http.StatusOK, response)
}

func logout(ctx *gin.Context) {
	// Get token
	token, err := jwtUtil.GetToken(ctx)
	if err != nil {
		response := presenter.BasicResponse{
			Status:  fmt.Sprint(http.StatusOK),
			Message: config.SUCCESS,
			Results: nil,
		}
		ctx.JSON(http.StatusOK, response)
		return
	}

	// Response in JSON
	response := presenter.BasicResponse{
		Status:  fmt.Sprint(http.StatusOK),
		Message: config.SUCCESS,
		Results: token,
	}
	ctx.JSON(http.StatusOK, response)
}
