package routes

import (
	"gin-gorm-microservice/infrastructure/rest/adapter"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ApplicationV1Router(router *gin.Engine, db *gorm.DB) {
	routerV1 := router.Group("/v1")
	MedicineRoutes(routerV1, adapter.MedicineAdapter(db))
	AuthRoutes(routerV1, adapter.AuthAdapter(db))
	UserRoutes(routerV1, adapter.UserAdapter(db))
}
