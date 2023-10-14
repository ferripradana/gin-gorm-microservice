package auth

type LoginRequest struct {
	Email    string `json:"email" example:"mail@mail.com" gorm:"unique" binding:"required"`
	Password string `json:"password" example:"Password123" binding:"required"`
}

type RegisterRequest struct {
	UserName  string `json:"user" example:"someUser" gorm:"unique" binding:"required"`
	Email     string `json:"email" example:"mail@mail.com" gorm:"unique" binding:"required"`
	FirstName string `json:"first_name" example:"John" binding:"required"`
	LastName  string `json:"last_name" example:"Doe" binding:"required"`
	Password  string `json:"password" example:"Password123" binding:"required"`
}

type AccessTokenRequest struct {
	RefreshToken string `json:"refreshToken" example:"badbunybabybebe" binding:"required"`
}
