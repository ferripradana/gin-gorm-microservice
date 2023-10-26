package routes

import (
	_ "gin-gorm-microservice/docs"
	"gin-gorm-microservice/infrastructure/rest/adapter"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func ApplicationV1Router(router *gin.Engine, db *gorm.DB) {
	routerV1 := router.Group("/v1")
	{
		{
			routerV1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}
		MedicineRoutes(routerV1, adapter.MedicineAdapter(db))
		AuthRoutes(routerV1, adapter.AuthAdapter(db))
		UserRoutes(routerV1, adapter.UserAdapter(db))
	}
}
