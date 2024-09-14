package middleware

import (
	"aedronef1_backend/entity"
	"aedronef1_backend/infrastucture/repository/util"
	jwtUtil "aedronef1_backend/infrastucture/repository/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtVerifyMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := jwtUtil.GetToken(ctx)
		if err != nil {
			util.HandlerException(ctx, http.StatusUnauthorized, entity.ErrUnauthorized)
			return
		}

		claims := entity.TokenClaims{}
		if _, err := jwtUtil.ParseToken(token, &claims); err != nil {
			util.HandlerException(ctx, http.StatusUnauthorized, entity.ErrUnauthorized)
			return
		}
		ctx.Next()
	}
}
