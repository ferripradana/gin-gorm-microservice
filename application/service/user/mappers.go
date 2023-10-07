package user

import (
	domainUser "gin-gorm-microservice/domain/user"
)

func (n *NewUser) toDomainMapper() *domainUser.User {
	return &domainUser.User{
		UserName:  n.UserName,
		Email:     n.Email,
		FirstName: n.FirstName,
		LastName:  n.LastName,
		Status:    n.Status,
	}
}
