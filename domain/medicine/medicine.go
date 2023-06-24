package medicine

import "time"

type Medicine struct {
	ID          int       `json:"id" example:"123"`
	Name        string    `json:"name" example:"Paracetamol"`
	Description string    `json:"description" example:"Some Description"`
	EANCode     string    `json:"ean_code" example:"9900000124"`
	Laboratory  string    `json:"laboratory" example:"Roche"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
