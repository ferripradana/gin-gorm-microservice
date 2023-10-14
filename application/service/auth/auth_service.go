package auth

import domainUser "gin-gorm-microservice/domain/user"

type AuthService interface {
	Login(loginUser LoginUser) (*SecurityAuthenticatedUser, error)
	AccessTokenByRefreshToken(refreshToken string) (*SecurityAuthenticatedUser, error)
	Register(newUser *RegisterUser) (*domainUser.User, error)
}
