package medicine

import (
	"gin-gorm-microservice/application/service/medicine"
	"gin-gorm-microservice/domain/errors"
	"gin-gorm-microservice/infrastructure/rest/controllers"
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

// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
func (controller *MedicineControllerImpl) NewMedicine(ctx *gin.Context) {
	var request NewMedicineRequest
	if err := controllers.BindJSON(ctx, &request); err != nil {
		appError := errors.NewAppErrorImpl(err, errors.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	newMedicine := medicine.NewMedicine{
		Name:        request.Name,
		Description: request.Description,
		EANCode:     request.EanCode,
		Laboratory:  request.Laboratory,
	}

	domainMedicine, err := controller.Service.Create(&newMedicine)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, domainMedicine)
}
