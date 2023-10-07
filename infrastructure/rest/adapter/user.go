package adapter

import (
	user3 "gin-gorm-microservice/application/service/user"
	user2 "gin-gorm-microservice/infrastructure/repository/user"
	"gin-gorm-microservice/infrastructure/rest/controllers/user"
	"gorm.io/gorm"
)

func UserAdapter(db *gorm.DB) user.UserController {
	userRepository := user2.NewUserRepositoryImpl(db)
	userService := user3.NewUserServiceImpl(userRepository)
	return user.NewUserControllerImpl(userService)
}
