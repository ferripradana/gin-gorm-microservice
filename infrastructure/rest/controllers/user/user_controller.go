package user

import "github.com/gin-gonic/gin"

type UserController interface {
	NewUser(ctx *gin.Context)
}
