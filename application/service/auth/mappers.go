package auth

import (
	userDomain "gin-gorm-microservice/domain/user"
)

func secAuthUserMapper(domainUser *userDomain.User, authInfo *Auth) *SecurityAuthenticatedUser {
	return &SecurityAuthenticatedUser{
		Data: DataUserAuthenticated{
			ID:        domainUser.ID,
			UserName:  domainUser.UserName,
			Email:     domainUser.Email,
			FirstName: domainUser.FirstName,
			LastName:  domainUser.LastName,
			Status:    domainUser.Status,
		},
		Security: DataSecurityAuthenticated{
			JWTAccessToken:            authInfo.AccessToken,
			JWTRefreshToken:           authInfo.RefreshToken,
			ExpirationAccessDateTime:  authInfo.ExpirationAccessDateTime,
			ExpirationRefreshDateTime: authInfo.ExpirationRefreshDateTime,
		},
	}
}

func (registerUser *RegisterUser) toDomainMapper() *userDomain.User {
	return &userDomain.User{
		UserName:  registerUser.UserName,
		Email:     registerUser.Email,
		FirstName: registerUser.FirstName,
		LastName:  registerUser.LastName,
		Status:    false,
	}
}
