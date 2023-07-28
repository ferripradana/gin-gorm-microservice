package medicine

import (
	"encoding/json"
	"gin-gorm-microservice/domain/errors"
	"gin-gorm-microservice/domain/medicine"
	"gorm.io/gorm"
)

type MedicineRepositoryImpl struct {
	DB *gorm.DB
}

func NewMedicineRepositoryImpl(db *gorm.DB) MedicineRepository {
	return &MedicineRepositoryImpl{DB: db}
}

func (m *MedicineRepositoryImpl) GetAll(page int64, limit int64) (*PaginationResultMedicine, error) {
	var medicines []Medicine
	var total int64

	err := m.DB.Model(&Medicine{}).Count(&total).Error
	if err != nil {
		return &PaginationResultMedicine{}, err
	}
	println(total)
	offset := (page - 1) * limit
	//println(limit)
	err = m.DB.Limit(int(limit)).Offset(int(offset)).Find(&medicines).Error
	if err != nil {
		return &PaginationResultMedicine{}, err
	}
	numPages := (total + limit - 1) / limit
	var nextCursor, prevCursor uint
	if page < numPages {
		nextCursor = uint(page + 1)
	}
	if page > 1 {
		prevCursor = uint(page - 1)
	}

	return &PaginationResultMedicine{
		Data:       arrayToDomainMapper(&medicines),
		Total:      total,
		Limit:      limit,
		Current:    page,
		NextCursor: nextCursor,
		PrevCursor: prevCursor,
		NumPages:   numPages,
	}, nil
}

func (m *MedicineRepositoryImpl) Create(newMedicine *medicine.Medicine) (createdMedicine *medicine.Medicine, err error) {
	medicine := fromDomainMapper(newMedicine)
	tx := m.DB.Create(medicine)

	if tx.Error != nil {
		byteErr, _ := json.Marshal(tx.Error)
		var newError errors.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {
			return createdMedicine, err
		}
		switch newError.Number {
		case 1062:
			err = errors.NewAppErrorWithType(errors.ResourceAlreadyExists)
		default:
			err = errors.NewAppErrorWithType(errors.UnknownError)
		}
		return createdMedicine, err
	}
	createdMedicine = medicine.toDomainMapper()
	return createdMedicine, nil
}

func (m *MedicineRepositoryImpl) GetById(id int) (*medicine.Medicine, error) {
	var _medicine Medicine
	err := m.DB.Where("id = ?", id).First(&_medicine).Error

	if err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			err = errors.NewAppErrorWithType(errors.NotFound)
		default:
			err = errors.NewAppErrorWithType(errors.UnknownError)
		}
		return &medicine.Medicine{}, err
	}
	return _medicine.toDomainMapper(), nil
}

func (m *MedicineRepositoryImpl) GetOneByMap(medicineMap map[string]interface{}) (*medicine.Medicine, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MedicineRepositoryImpl) Update(id int, medicineMap map[string]interface{}) (*medicine.Medicine, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MedicineRepositoryImpl) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}
