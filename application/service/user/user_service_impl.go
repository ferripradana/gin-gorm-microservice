package user

import (
	domainUser "gin-gorm-microservice/domain/user"
	"gin-gorm-microservice/infrastructure/repository/user"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository user.UserRepository
}

func NewUserServiceImpl(userRepository user.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
	}
}

func (userService *UserServiceImpl) Create(newUser *NewUser) (*domainUser.User, error) {
	domain := newUser.toDomainMapper()

	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return &domainUser.User{}, err
	}

	domain.HashPassword = string(hash)
	return userService.UserRepository.Create(domain)
}

func (userService *UserServiceImpl) GetById(id int) (*domainUser.User, error) {
	return userService.UserRepository.GetById(id)
}

func (userService *UserServiceImpl) GetAll(page int64, limit int64) (*PaginationResultUser, error) {
	all, err := userService.UserRepository.GetAll(page, limit)
	if err != nil {
		return nil, err
	}
	return &PaginationResultUser{
		Data:       all.Data,
		Total:      all.Total,
		Limit:      all.Limit,
		Current:    all.Current,
		NextCursor: all.NextCursor,
		PrevCursor: all.PrevCursor,
		NumPages:   all.NumPages,
	}, nil
}

func (userService *UserServiceImpl) Delete(id int) error {
	return userService.UserRepository.Delete(id)
}

func (userService *UserServiceImpl) Update(id int, userMap map[string]interface{}) (*domainUser.User, error) {
	return userService.UserRepository.Update(id, userMap)
}
