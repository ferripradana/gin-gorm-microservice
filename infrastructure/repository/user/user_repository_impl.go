package user

import (
	"encoding/json"
	"gin-gorm-microservice/domain/errors"
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

func (u *UserRepositoryImpl) Create(newUser *domainUser.User) (createdUser *domainUser.User, err error) {
	userRepo := fromDomainMapper(newUser)
	txDb := u.DB.Create(userRepo)

	if txDb.Error != nil {
		byteErr, _ := json.Marshal(err)
		var newError errors.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {
			return createdUser, err
		}
		switch newError.Number {
		case 1062:
			err = errors.NewAppErrorWithType(errors.ResourceAlreadyExists)
		case 0:
			err = errors.NewAppErrorWithType(errors.ResourceAlreadyExists)
		default:
			err = errors.NewAppErrorWithType(errors.UnknownError)
		}
		return createdUser, err
	}
	createdUser = userRepo.toDomainMapper()
	return createdUser, err
}
