package medicine

import "gin-gorm-microservice/domain/medicine"

type MedicineService interface {
	GetAll(page int64, limit int64) (*PaginationResultMedicine, error)
	Create(medicine *NewMedicine) (*medicine.Medicine, error)
}
