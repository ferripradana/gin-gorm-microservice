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

	tx := u.DB.Where(userMap).Where("status = ?", true).Limit(1).Find(&userRepo)
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

func (u *UserRepositoryImpl) GetById(id int) (*domainUser.User, error) {
	var user User
	err := u.DB.Where("id = ?", id).First(&user).Error

	if err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			err = errors.NewAppErrorWithType(errors.NotFound)
		default:
			err = errors.NewAppErrorWithType(errors.UnknownError)
		}
		return &domainUser.User{}, err
	}

	return user.toDomainMapper(), nil
}

func (u *UserRepositoryImpl) GetAll(page int64, limit int64) (*PaginationResultUser, error) {
	var users []User
	var total int64

	err := u.DB.Model(&User{}).Count(&total).Error
	if err != nil {
		return &PaginationResultUser{}, err
	}

	offset := (page - 1) * limit
	err = u.DB.Limit(int(limit)).Offset(int(offset)).Find(&users).Error
	if err != nil {
		return &PaginationResultUser{}, err
	}
	numPages := (total + limit - 1) / limit
	var nextCursor, prevCursor uint
	if page < numPages {
		nextCursor = uint(page + 1)
	}

	if page > 1 {
		prevCursor = uint(page - 1)
	}

	return &PaginationResultUser{
		Data:       arrayToDomainMapper(&users),
		Total:      total,
		Limit:      limit,
		Current:    page,
		NextCursor: nextCursor,
		PrevCursor: prevCursor,
		NumPages:   numPages,
	}, nil
}

func (u *UserRepositoryImpl) Delete(id int) (err error) {
	tx := u.DB.Delete(&User{}, id)
	if tx.Error != nil {
		err = errors.NewAppErrorImpl(err, errors.UnknownError)
		return err
	}

	if tx.RowsAffected == 0 {
		err = errors.NewAppErrorWithType(errors.NotFound)
		return err
	}

	return nil
}

func (u *UserRepositoryImpl) Update(id int, userMap map[string]interface{}) (*domainUser.User, error) {
	var _user User
	_user.ID = id
	userMap["user_name"] = userMap["user"]

	err := u.DB.Where("id = ?", id).First(&_user).Error
	if err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			err = errors.NewAppErrorWithType(errors.NotFound)
		default:
			err = errors.NewAppErrorWithType(errors.UnknownError)
		}
		return &domainUser.User{}, err
	}

	err = u.DB.Model(&_user).
		Select("first_name", "last_name", "status", "user_name", "email").
		Updates(userMap).Error

	if err != nil {
		byteErr, _ := json.Marshal(err)
		var newError errors.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {
			err = errors.NewAppErrorImpl(err, errors.UnknownError)
			return &domainUser.User{}, err
		}
		switch newError.Number {
		case 1062:
			err = errors.NewAppErrorWithType(errors.ResourceAlreadyExists)
			return &domainUser.User{}, err
		default:
			err = errors.NewAppErrorWithType(errors.UnknownError)
			return &domainUser.User{}, err
		}
	}
	err = u.DB.Where("id = ?", id).First(&_user).Error
	if err != nil {
		err = errors.NewAppErrorWithType(errors.NotFound)
	}
	return _user.toDomainMapper(), nil
}
