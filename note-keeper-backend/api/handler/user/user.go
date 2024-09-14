package handler

import (
	"backend_mobile_app/api/presenter"
	"backend_mobile_app/config"
	"backend_mobile_app/usecase/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUserInfo(ctx *gin.Context, userService user.UseCase) {

	response := presenter.BasicResponse{
		Status:  fmt.Sprint(http.StatusOK),
		Message: config.SUCCESS,
		Results: (nil),
	}
	ctx.JSON(http.StatusOK, response)
}
