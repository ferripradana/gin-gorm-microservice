package auth

import "gin-gorm-microservice/infrastructure/rest/controllers/user"

type ResponseRegistration struct {
	ResponseUser user.ResponseUser `json:"user"`
	Message      string            `json:"message"`
}
