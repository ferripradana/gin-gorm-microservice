package user

import (
	domainUser "gin-gorm-microservice/domain/user"
)

func (user *User) toDomainMapper() *domainUser.User {
	return &domainUser.User{
		ID:           user.ID,
		UserName:     user.UserName,
		Email:        user.Email,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Status:       user.Status,
		HashPassword: user.HashPassword,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}

func fromDomainMapper(user *domainUser.User) *User {
	return &User{
		ID:           user.ID,
		UserName:     user.UserName,
		Email:        user.Email,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Status:       user.Status,
		HashPassword: user.HashPassword,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}

func arrayToDomainMapper(users *[]User) *[]domainUser.User {
	userDomain := make([]domainUser.User, len(*users))
	for i, user := range *users {
		userDomain[i] = *user.toDomainMapper()
	}
	return &userDomain
}
