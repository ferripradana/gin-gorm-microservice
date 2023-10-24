package user

import (
	"gin-gorm-microservice/application/service/user"
	userDomain "gin-gorm-microservice/domain/user"
)

func toUserServiceMapper(request *NewUserRequest) *user.NewUser {
	return &user.NewUser{
		UserName:  request.UserName,
		Email:     request.Email,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Password:  request.Password,
		Status:    request.Status,
	}
}

func domainToResponseMapper(user *userDomain.User) *ResponseUser {
	return &ResponseUser{
		ID:        user.ID,
		UserName:  user.UserName,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func mapFromDomainToResponse(users *[]userDomain.User) *[]ResponseUser {
	usersResponse := make([]ResponseUser, len(*users))
	for i, user := range *users {
		usersResponse[i] = *domainToResponseMapper(&user)
	}
	return &usersResponse
}
