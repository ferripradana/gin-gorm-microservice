package medicine

import "gin-gorm-microservice/domain/medicine"

type MedicineService interface {
	GetAll(page int64, limit int64) (*PaginationResultMedicine, error)
	Create(medicine *NewMedicine) (*medicine.Medicine, error)
	GetById(id int) (*medicine.Medicine, error)
	Update(id int, medicineMap map[string]interface{}) (*medicine.Medicine, error)
}
