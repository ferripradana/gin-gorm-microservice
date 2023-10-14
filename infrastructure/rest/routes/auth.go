package routes

import (
	"gin-gorm-microservice/infrastructure/rest/controllers/auth"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup, controller auth.AuthController) {
	routerAuth := router.Group("/auth")
	{
		routerAuth.POST("/login", controller.Login)
		routerAuth.POST("/access-token", controller.GetAccessTokenByRefreshToken)
		routerAuth.POST("/register", controller.Register)
	}
}
