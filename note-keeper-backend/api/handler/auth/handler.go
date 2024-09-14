package handler

import (
	"aedronef1_backend/usecase/user"

	"github.com/gin-gonic/gin"
)

func MakeHandlers(app *gin.Engine, userService user.UseCase) {
	authGroup := app.Group("/api/auth")
	{
		authGroup.POST("/login", func(ctx *gin.Context) {
			login(ctx, userService)
		})

		authGroup.POST("/sign-up", func(ctx *gin.Context) {
			signUp(ctx, userService)
		})

		authGroup.POST("/logout", func(ctx *gin.Context) {
			logout(ctx)
		})
	}
}
