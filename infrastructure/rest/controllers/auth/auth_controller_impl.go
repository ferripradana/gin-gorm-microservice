package auth

import (
	authService "gin-gorm-microservice/application/service/auth"
	"gin-gorm-microservice/domain/errors"
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
		appError := errors.NewAppErrorImpl(err, errors.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	user := authService.LoginUser{
		Email:    request.Email,
		Password: request.Password,
	}

	authDataUser, err := controller.AuthService.Login(user)
	if err != nil {
		appError := errors.NewAppErrorImpl(err, errors.NotAuthorized)
		_ = ctx.Error(appError)
		return
	}
	ctx.JSON(http.StatusOK, authDataUser)
}

func (controller *AuthControllerImpl) GetAccessTokenByRefreshToken(ctx *gin.Context) {
	var request AccessTokenRequest

	if err := controllers.BindJSON(ctx, &request); err != nil {
		_ = ctx.Error(err)
		return
	}

	authDataUser, err := controller.AuthService.AccessTokenByRefreshToken(request.RefreshToken)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, authDataUser)
}

func (controller *AuthControllerImpl) Register(ctx *gin.Context) {
	var request RegisterRequest
	if err := controllers.BindJSON(ctx, &request); err != nil {
		appError := errors.NewAppErrorImpl(err, errors.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	domainUser, err := controller.AuthService.Register(toAuthServiceMapper(&request))
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	registerResponse := domainToResponseMapper(domainUser)
	registerResponse.Message = "Waiting Approval"
	ctx.JSON(http.StatusOK, registerResponse)
}
