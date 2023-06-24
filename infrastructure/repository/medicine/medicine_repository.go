package medicine

import "gin-gorm-microservice/domain/medicine"

type MedicineRepository interface {
	GetAll(page int64, limit int64) (*PaginationResultMedicine, error)
	Create(newMedicine *medicine.Medicine) (*medicine.Medicine, error)
	GetById(id int) (*medicine.Medicine, error)
	GetOneByMap(medicineMap map[string]interface{}) (*medicine.Medicine, error)
	Update(id int, medicineMap map[string]interface{}) (*medicine.Medicine, error)
	Delete(id int) error
}
