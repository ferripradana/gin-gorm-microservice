package medicine

import (
	"gin-gorm-microservice/domain/medicine"
)

func (n *NewMedicine) toDomainMapper() *medicine.Medicine {
	return &medicine.Medicine{
		Name:        n.Name,
		Description: n.Description,
		EANCode:     n.EANCode,
		Laboratory:  n.Laboratory,
	}
}
