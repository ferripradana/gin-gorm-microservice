package medicine

import (
	medicineDomain "gin-gorm-microservice/domain/medicine"
	"gin-gorm-microservice/infrastructure/repository/medicine"
)

type MedicineServiceImpl struct {
	MedicineRepository medicine.MedicineRepository
}

func NewMedicineServiceImpl(medicineRepository medicine.MedicineRepository) MedicineService {
	return &MedicineServiceImpl{
		MedicineRepository: medicineRepository,
	}
}

func (service *MedicineServiceImpl) GetAll(page int64, limit int64) (*PaginationResultMedicine, error) {
	all, err := service.MedicineRepository.GetAll(page, limit)
	if err != nil {
		return nil, err
	}
	return &PaginationResultMedicine{
		Data:       all.Data,
		Total:      all.Total,
		Limit:      all.Limit,
		Current:    all.Current,
		NextCursor: all.NextCursor,
		PrevCursor: all.PrevCursor,
		NumPages:   all.NumPages,
	}, nil
}

func (service *MedicineServiceImpl) Create(medicine *NewMedicine) (*medicineDomain.Medicine, error) {
	medicineToDomain := medicine.toDomainMapper()
	return service.MedicineRepository.Create(medicineToDomain)
}

func (service *MedicineServiceImpl) GetById(id int) (*medicineDomain.Medicine, error) {
	return service.MedicineRepository.GetById(id)
}

func (service *MedicineServiceImpl) Update(id int, medicineMap map[string]interface{}) (*medicineDomain.Medicine, error) {
	return service.MedicineRepository.Update(id, medicineMap)
}

func (service *MedicineServiceImpl) Delete(id int) error {
	return service.MedicineRepository.Delete(id)
}
