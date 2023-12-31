package medicine

import (
	goError "errors"
	"gin-gorm-microservice/application/service/medicine"
	"gin-gorm-microservice/domain/errors"
	medicine2 "gin-gorm-microservice/domain/medicine"
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

func (controller *MedicineControllerImpl) GetMedicinesById(ctx *gin.Context) {
	medicineId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		appError := errors.NewAppErrorImpl(err, errors.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	domainMedicine, err := controller.Service.GetById(medicineId)
	if err != nil {
		appError := errors.NewAppErrorImpl(err, errors.ValidationError)
		_ = ctx.Error(appError)
		return
	}
	ctx.JSON(http.StatusOK, domainMedicine)
}

func (controller *MedicineControllerImpl) UpdateMedicine(ctx *gin.Context) {
	medicineId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		appError := errors.NewAppErrorImpl(goError.New("param id is necessary in the url"), errors.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	var requestMap map[string]interface{}
	err = controllers.BindJSON(ctx, &requestMap)
	if err != nil {
		appError := errors.NewAppErrorImpl(err, errors.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	err = UpdateValidation(requestMap)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	var _medicine *medicine2.Medicine
	_medicine, err = controller.Service.Update(medicineId, requestMap)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, _medicine)
}

func (controller *MedicineControllerImpl) DeleteMedicine(ctx *gin.Context) {
	medicineId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		appError := errors.NewAppErrorImpl(goError.New("param id is necessary in the url"), errors.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	err = controller.Service.Delete(medicineId)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Resource deleted successfully"})
}
