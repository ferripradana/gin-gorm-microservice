package adapter

import (
	medicine3 "gin-gorm-microservice/application/service/medicine"
	"gin-gorm-microservice/infrastructure/repository/medicine"
	medicine2 "gin-gorm-microservice/infrastructure/rest/controllers/medicine"
	"gorm.io/gorm"
)

func MedicineAdapter(db *gorm.DB) medicine2.MedicineController {
	medicineRepository := medicine.NewMedicineRepositoryImpl(db)
	medicineService := medicine3.NewMedicineServiceImpl(medicineRepository)
	return medicine2.NewMedicineControllerImpl(medicineService)
}
