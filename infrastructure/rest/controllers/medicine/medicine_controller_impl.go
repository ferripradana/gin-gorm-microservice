package medicine

import (
	"gin-gorm-microservice/application/service/medicine"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type MedicineControllerImpl struct {
	Service medicine.MedicineService
}

func NewMedicineControllerImpl(service medicine.MedicineService) MedicineController {
	return &MedicineControllerImpl{
		Service: service,
	}
}

func (controller *MedicineControllerImpl) GetAllMedicines(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("limit", "10")

	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		_ = ctx.Error(err)
	}

	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		_ = ctx.Error(err)
	}

	medicines, err := controller.Service.GetAll(page, limit)
	if err != nil {
		_ = ctx.Error(err)
	}
	ctx.JSON(http.StatusOK, medicines)
}
