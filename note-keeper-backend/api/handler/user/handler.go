package handler

import (
	"backend_mobile_app/usecase/user"

	"github.com/gin-gonic/gin"

	authMiddleware "backend_mobile_app/api/middleware"
)

func MakeHandlers(app *gin.Engine, userService user.UseCase) {
	userGroup := app.Group("/api/user")
	{
		userGroup.GET("/information", authMiddleware.JwtVerifyMiddleware(), func(ctx *gin.Context) {
			getUserInfo(ctx, userService)
		})
	}
}
