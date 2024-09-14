package util

import (
	"aedronef1_backend/api/presenter"
	"fmt"

	"github.com/gin-gonic/gin"
)

func HandlerException(ctx *gin.Context, statusCode int, err error) {
	errorMessage := &presenter.BasicResponse{
		Status:  fmt.Sprint(statusCode),
		Message: err.Error(),
	}

	ctx.AbortWithStatusJSON(statusCode, errorMessage)
}
