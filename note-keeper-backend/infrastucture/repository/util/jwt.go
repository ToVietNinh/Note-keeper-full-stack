package util

import (
	"aedronef1_backend/config"
	"aedronef1_backend/entity"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func keyFunc(token *jwt.Token) (interface{}, error) {
	if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
		return nil, errors.New(config.UNAUTHORIZED)
	}
	return []byte(config.GetString("jwt.key")), nil
}

func GetToken(ctx *gin.Context) (string, error) {
	authHeader := ctx.GetHeader(config.GetString("jwt.header"))
	if len(authHeader) <= len(config.GetString("jwt.schema"))+1 {
		return "", errors.New(config.UNAUTHORIZED)
	}

	tokenString := authHeader[len(config.GetString("jwt.schema"))+1:]

	return tokenString, nil
}

func ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, keyFunc)
}

func ParseToken(token string, claims jwt.Claims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, claims, keyFunc)
}

func ParseUnverifiedToken(token string, claims jwt.Claims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, claims, nil)
}

func GenerateAccessToken(user *entity.User) (string, error) {
	// Generate random
	random, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	// List function

	// Generate claims
	claims := entity.TokenClaims{
		UserId: user.Id,
		Jti:    random.String(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(config.GetInt("jwt.accessMaxAge"))).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign with secret
	signedToken, err := token.SignedString([]byte(config.GetString("jwt.key")))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func GenerateRefreshToken(user *entity.User) (string, error) {
	claims := entity.RefreshTokenClaims{
		UserId: user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(config.GetInt("jwt.refreshMaxAge"))).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign with secret
	signedRefreshToken, err := refreshToken.SignedString([]byte(config.GetString("jwt.key")))
	if err != nil {
		return "", err
	}

	return signedRefreshToken, nil
}

func GenerateForgotPasswordToken(user *entity.User) (string, error) {
	// Generate claims
	claims := entity.ForgotPasswordToken{
		UserId: user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(config.GetInt("jwt.otpMaxAge"))).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign with secret
	signedToken, err := token.SignedString([]byte(config.GetString("jwt.key")))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
