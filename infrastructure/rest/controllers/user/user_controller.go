package user

import "github.com/gin-gonic/gin"

type UserController interface {
	NewUser(ctx *gin.Context)
	GetUserById(ctx *gin.Context)
	GetAllUsers(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
}
