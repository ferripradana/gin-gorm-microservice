package user

import (
	"gin-gorm-microservice/application/service/user"
	"gin-gorm-microservice/domain/errors"
	"gin-gorm-microservice/infrastructure/rest/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserControllerImpl struct {
	Service user.UserService
}

func NewUserControllerImpl(service user.UserService) UserController {
	return &UserControllerImpl{
		Service: service,
	}
}

func (u *UserControllerImpl) NewUser(ctx *gin.Context) {
	var request NewUserRequest
	if err := controllers.BindJSON(ctx, &request); err != nil {
		appError := errors.NewAppErrorImpl(err, errors.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	domainUser, err := u.Service.Create(toUserServiceMapper(&request))
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	userResponse := domainToResponseMapper(domainUser)
	ctx.JSON(http.StatusOK, userResponse)
}
