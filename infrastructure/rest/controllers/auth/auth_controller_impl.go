package auth

import (
	authService "gin-gorm-microservice/application/service/auth"
	"gin-gorm-microservice/infrastructure/rest/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthControllerImpl struct {
	AuthService authService.AuthService
}

func NewAuthControllerImpl(service authService.AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: service,
	}
}

func (controller *AuthControllerImpl) Login(ctx *gin.Context) {
	var request LoginRequest

	if err := controllers.BindJSON(ctx, &request); err != nil {
		_ = ctx.Error(err)
		return
	}

	user := authService.LoginUser{
		Email:    request.Email,
		Password: request.Password,
	}

	authDataUser, err := controller.AuthService.Login(user)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, authDataUser)
}

func (controller *AuthControllerImpl) GetAccessTokenByRefreshToken(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
