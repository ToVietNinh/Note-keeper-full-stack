package entity

import (
	"aedronef1_backend/config"
	"errors"
)

var ErrLock = errors.New(config.LOCK)
var ErrUnauthorized = errors.New(config.UNAUTHORIZED)
var ErrInternalServerError = errors.New(config.INTERNAL_SERVER_ERROR)
var ErrBadRequest = errors.New(config.BAD_REQUEST)
var ErrUsernameNotExists = errors.New(config.USERNAME_NOT_EXISTS)
var ErrUsernameOrEmailNotExists = errors.New(config.USERNAME_OR_EMAIL_NOT_EXISTS)
var ErrInvalidPassword = errors.New(config.INVALID_PASSWORD)
var ErrForbidden = errors.New(config.FORBIDDEN)
var ErrEmailNotExists = errors.New(config.EMAIL_NOT_EXISTS)
var InvalidOTPCode = errors.New(config.INVALID_OTP_CODE)
