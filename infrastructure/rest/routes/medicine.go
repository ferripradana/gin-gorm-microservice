package routes

import (
	"gin-gorm-microservice/infrastructure/rest/controllers/medicine"
	"gin-gorm-microservice/infrastructure/rest/middlewares"
	"github.com/gin-gonic/gin"
)

func MedicineRoutes(router *gin.RouterGroup, controller medicine.MedicineController) {
	routerMedicine := router.Group("/medicine")
	routerMedicine.Use(middlewares.AuthJWTMiddleware())
	{
		routerMedicine.GET("/", controller.GetAllMedicines)
		routerMedicine.POST("/", controller.NewMedicine)
		routerMedicine.GET("/:id", controller.GetMedicinesById)
	}
}
