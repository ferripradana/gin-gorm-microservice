package user

import (
	domainUser "gin-gorm-microservice/domain/user"
	"time"
)

// User is a struct that contains the user information
type User struct {
	ID           int       `json:"id" example:"1099" gorm:"primaryKey"`
	UserName     string    `json:"user_name" example:"UserName" gorm:"column:user_name;unique"`
	Email        string    `json:"email" example:"some@mail.com" gorm:"unique"`
	FirstName    string    `json:"first_name" example:"firstname"`
	LastName     string    `json:"last_name" example:"lastname"`
	Status       bool      `json:"status" example:"false"`
	HashPassword string    `json:"hash_password" example:"password"`
	CreatedAt    time.Time `json:"updated_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoCreateTime:mili"`
	UpdatedAt    time.Time `json:"updated_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoUpdateTime:mili"`
}

// TableName overrides the table name used by User to `users`
func (*User) name() string {
	return "users"
}

// PaginationResultUser is a struct that contains the pagination result for user
type PaginationResultUser struct {
	Data       *[]domainUser.User
	Total      int64
	Limit      int64
	Current    int64
	NextCursor uint
	PrevCursor uint
	NumPages   int64
}
