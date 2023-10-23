package user

import domainUser "gin-gorm-microservice/domain/user"

type UserRepository interface {
	GetOneByMap(userMap map[string]interface{}) (*domainUser.User, error)
	Create(newUser *domainUser.User) (*domainUser.User, error)
	GetById(id int) (*domainUser.User, error)
}
