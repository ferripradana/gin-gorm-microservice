package adapter

import (
	authService "gin-gorm-microservice/application/service/auth"
	"gin-gorm-microservice/infrastructure/repository/user"
	"gin-gorm-microservice/infrastructure/rest/controllers/auth"
	"gorm.io/gorm"
)

func AuthAdapter(db *gorm.DB) auth.AuthController {
	userRepository := user.NewUserRepositoryImpl(db)
	authService := authService.NewAuthServiceImpl(userRepository)
	return auth.NewAuthControllerImpl(authService)
}
