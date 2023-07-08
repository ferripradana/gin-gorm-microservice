package medicine

import "github.com/gin-gonic/gin"

type MedicineController interface {
	GetAllMedicines(ctx *gin.Context)
	NewMedicine(ctx *gin.Context)
}
