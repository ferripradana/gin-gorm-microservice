package medicine

import "gin-gorm-microservice/domain/medicine"

type NewMedicine struct {
	Name        string `json:"name" example:"Paracetamol"`
	Description string `json:"description" example:"Paracetamol"`
	EANCode     string `json:"ean_code" example:"7837373"`
	Laboratory  string `json:"laboratory" example:"Bayer"`
}

type PaginationResultMedicine struct {
	Data       *[]medicine.Medicine
	Total      int64
	Limit      int64
	Current    int64
	NextCursor uint
	PrevCursor uint
	NumPages   int64
}
