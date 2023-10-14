package auth

import "github.com/gin-gonic/gin"

type AuthController interface {
	Login(ctx *gin.Context)
	GetAccessTokenByRefreshToken(ctx *gin.Context)
	Register(ctx *gin.Context)
}
