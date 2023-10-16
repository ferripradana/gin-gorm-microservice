package user

import domainUser "gin-gorm-microservice/domain/user"

type UserService interface {
	Create(newUser *NewUser) (*domainUser.User, error)
}