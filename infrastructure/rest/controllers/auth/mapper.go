package auth

import (
	"gin-gorm-microservice/application/service/auth"
	domainUser "gin-gorm-microservice/domain/user"
	authStructure "gin-gorm-microservice/infrastructure/rest/controllers/user"
)

func toAuthServiceMapper(request *RegisterRequest) *auth.RegisterUser {
	return &auth.RegisterUser{
		UserName:  request.UserName,
		Email:     request.Email,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Password:  request.Password,
	}
}

func domainToResponseMapper(user *domainUser.User) *ResponseRegistration {
	return &ResponseRegistration{
		ResponseUser: authStructure.ResponseUser{
			ID:        user.ID,
			UserName:  user.UserName,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Status:    user.Status,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}
}
