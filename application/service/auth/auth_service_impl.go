package auth

import (
	"errors"
	"gin-gorm-microservice/application/security/jwt"
	domainUser "gin-gorm-microservice/domain/user"
	"gin-gorm-microservice/infrastructure/repository/user"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthServiceImpl struct {
	UserRepository user.UserRepository
}

func NewAuthServiceImpl(userRepository user.UserRepository) AuthService {
	return &AuthServiceImpl{
		UserRepository: userRepository,
	}
}

func (authServ *AuthServiceImpl) Login(loginUser LoginUser) (*SecurityAuthenticatedUser, error) {
	userMap := make(map[string]interface{})
	userMap["email"] = loginUser.Email
	domainUser, err := authServ.UserRepository.GetOneByMap(userMap)
	if err != nil {
		return &SecurityAuthenticatedUser{}, err
	}
	if domainUser.ID == 0 {
		return &SecurityAuthenticatedUser{}, errors.New("email or password does not match")
	}

	isAuthenticated := CheckPasswordHash(loginUser.Password, domainUser.HashPassword)
	if !isAuthenticated {
		return &SecurityAuthenticatedUser{}, errors.New("email or password does not match")
	}

	accessTokenClaims, err := jwt.GenerateJWTToken(domainUser.ID, "access")
	if err != nil {
		return &SecurityAuthenticatedUser{}, err
	}

	refreshTokenClaims, err := jwt.GenerateJWTToken(domainUser.ID, "refresh")
	if err != nil {
		return &SecurityAuthenticatedUser{}, err
	}

	return secAuthUserMapper(domainUser, &Auth{
		AccessToken:               accessTokenClaims.Token,
		RefreshToken:              refreshTokenClaims.Token,
		ExpirationAccessDateTime:  accessTokenClaims.ExpirationTime,
		ExpirationRefreshDateTime: refreshTokenClaims.ExpirationTime,
	}), err
}

func (authServ *AuthServiceImpl) AccessTokenByRefreshToken(refreshToken string) (*SecurityAuthenticatedUser, error) {
	claimsMap, err := jwt.GetClaimsAndVerifyToken(refreshToken, "refresh")
	if err != nil {
		return nil, err
	}

	userMap := map[string]interface{}{"id": claimsMap["id"]}
	domainUser, err := authServ.UserRepository.GetOneByMap(userMap)
	if err != nil {
		return nil, err
	}

	accessTokenClaims, err := jwt.GenerateJWTToken(domainUser.ID, "access")
	if err != nil {
		return &SecurityAuthenticatedUser{}, err
	}

	var expTime = int64(claimsMap["exp"].(float64))
	return secAuthUserMapper(domainUser, &Auth{
		AccessToken:               accessTokenClaims.Token,
		RefreshToken:              refreshToken,
		ExpirationAccessDateTime:  accessTokenClaims.ExpirationTime,
		ExpirationRefreshDateTime: time.Unix(expTime, 0),
	}), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (authServ *AuthServiceImpl) Register(newUser *RegisterUser) (*domainUser.User, error) {
	newUserDomain := newUser.toDomainMapper()
	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return &domainUser.User{}, err
	}
	newUserDomain.HashPassword = string(hash)
	return authServ.UserRepository.Create(newUserDomain)
}
