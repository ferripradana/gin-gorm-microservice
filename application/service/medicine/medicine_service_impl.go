package medicine

import "gin-gorm-microservice/infrastructure/repository/medicine"

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
