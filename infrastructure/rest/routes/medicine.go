package routes

import (
	"gin-gorm-microservice/infrastructure/rest/controllers/medicine"
	"github.com/gin-gonic/gin"
)

func MedicineRoutes(router *gin.RouterGroup, controller medicine.MedicineController) {
	routerMedicine := router.Group("/medicine")
	{
		routerMedicine.GET("/", controller.GetAllMedicines)
	}
}
