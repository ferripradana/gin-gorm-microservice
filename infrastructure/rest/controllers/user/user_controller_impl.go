package user

import (
	"gin-gorm-microservice/application/service/user"
	"gin-gorm-microservice/domain/errors"
	"gin-gorm-microservice/infrastructure/rest/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func (u *UserControllerImpl) GetUserById(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		appError := errors.NewAppErrorImpl(err, errors.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	domainUser, err := u.Service.GetById(userId)
	if err != nil {
		appError := errors.NewAppErrorImpl(err, errors.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	ctx.JSON(http.StatusOK, domainToResponseMapper(domainUser))
}

func (u *UserControllerImpl) GetAllUsers(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("limit", "10")

	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		_ = ctx.Error(err)
	}

	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		_ = ctx.Error(err)
	}

	users, err := u.Service.GetAll(page, limit)
	usersResponse := &PaginationResultUser{
		Data:       mapFromDomainToResponse(users.Data),
		Total:      users.Total,
		Limit:      users.Limit,
		Current:    users.Current,
		NextCursor: users.NextCursor,
		PrevCursor: users.PrevCursor,
		NumPages:   users.NumPages,
	}
	if err != nil {
		_ = ctx.Error(err)
	}

	ctx.JSON(http.StatusOK, usersResponse)
}
