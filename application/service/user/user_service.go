package user

import domainUser "gin-gorm-microservice/domain/user"

type UserService interface {
	Create(newUser *NewUser) (*domainUser.User, error)
	GetById(id int) (*domainUser.User, error)
	GetAll(page int64, limit int64) (*PaginationResultUser, error)
	Delete(id int) error
}
