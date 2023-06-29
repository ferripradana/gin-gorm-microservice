package auth

import "time"

type LoginUser struct {
	Email    string
	Password string
}

type Auth struct {
	AccessToken               string
	RefreshToken              string
	ExpirationAccessDateTime  time.Time
	ExpirationRefreshDateTime time.Time
}

type DataUserAuthenticated struct {
	ID        int    `json:"id" example:"123"`
	UserName  string `json:"user_name" example:"UserName" gorm:"unique"`
	Email     string `json:"email" example:"some@mail.com" gorm:"unique"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Status    bool   `json:"status"`
}

type DataSecurityAuthenticated struct {
	JWTAccessToken            string    `json:"jwtAccessToken" example:"SomeAccessToken"`
	JWTRefreshToken           string    `json:"jwtRefreshToken" example:"SomeRefreshToken"`
	ExpirationAccessDateTime  time.Time `json:"expirationAccessDateTime" example:"2023-02-02T21:03:53.196419-06:00"`
	ExpirationRefreshDateTime time.Time `json:"expirationRefreshDateTime" example:"2023-02-03T06:53:53.196419-06:00"`
}

type SecurityAuthenticatedUser struct {
	Data     DataUserAuthenticated     `json:"data"`
	Security DataSecurityAuthenticated `json:"security"`
}
