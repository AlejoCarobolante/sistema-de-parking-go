package route

import (
	"gorm-template/api/controller"
	"gorm-template/bootstrap"
	"gorm-template/pkg/usecase"
	"time"

	//	"gorm-template/api/controller"

	"github.com/gin-gonic/gin"
)

func NewPagoRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	ec := &controller.PagoController{
		PagoRepository: &usecase.PagoUseCase{},
	}
	PagoRouter := group.Group("/pago")
	PagoRouter.POST("/", ec.Create)
	PagoRouter.GET("/", ec.Fetch)
	PagoRouter.GET("/:id", ec.FetchById)
	PagoRouter.PUT("/", ec.Update)
	PagoRouter.DELETE("/:id", ec.Delete)
}
