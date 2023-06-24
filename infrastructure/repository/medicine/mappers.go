package medicine

import (
	"gin-gorm-microservice/domain/medicine"
)

func (m *Medicine) toDomainMapper() *medicine.Medicine {
	return &medicine.Medicine{
		ID:          m.ID,
		Name:        m.Name,
		Description: m.Description,
		EANCode:     m.EANCode,
		Laboratory:  m.Laboratory,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

func fromDomainMapper(m *medicine.Medicine) *Medicine {
	return &Medicine{
		ID:          m.ID,
		Name:        m.Name,
		Description: m.Description,
		EANCode:     m.EANCode,
		Laboratory:  m.Laboratory,
		CreatedAt:   m.CreatedAt,
	}
}

func arrayToDomainMapper(medicines *[]Medicine) *[]medicine.Medicine {
	medicinesDomain := make([]medicine.Medicine, len(*medicines))
	for i, medicine := range *medicines {
		medicinesDomain[i] = *medicine.toDomainMapper()
	}

	return &medicinesDomain
}
