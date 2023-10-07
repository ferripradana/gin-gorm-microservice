package routes

import (
	"gin-gorm-microservice/infrastructure/rest/controllers/user"
	"gin-gorm-microservice/infrastructure/rest/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup, controller user.UserController) {
	routerUser := router.Group("/user")
	routerUser.Use(middlewares.AuthJWTMiddleware())
	{
		routerUser.POST("/", controller.NewUser)
	}
}
