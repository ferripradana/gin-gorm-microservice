package user

import (
	domainUser "gin-gorm-microservice/domain/user"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (u *UserRepositoryImpl) GetOneByMap(userMap map[string]interface{}) (*domainUser.User, error) {
	var userRepo User

	tx := u.DB.Where(userMap).Limit(1).Find(&userRepo)
	if tx.Error != nil {
		err := tx.Error
		return &domainUser.User{}, err
	}

	return userRepo.toDomainMapper(), nil
}
