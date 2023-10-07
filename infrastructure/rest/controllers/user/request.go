package user

type NewUserRequest struct {
	UserName  string `json:"user" example:"someUser" gorm:"unique" binding:"required"`
	Email     string `json:"email" example:"mail@mail.com" gorm:"unique" binding:"required"`
	FirstName string `json:"first_name" example:"John" binding:"required"`
	LastName  string `json:"last_name" example:"Doe" binding:"required"`
	Password  string `json:"password" example:"Password123" binding:"required"`
	Status    bool   `json:"status" example:"true" binding:"required"`
}
