package route

import (
	"gorm-template/api/controller"
	"gorm-template/bootstrap"
	"gorm-template/pkg/usecase"
	"time"

	//	"gorm-template/api/controller"

	"github.com/gin-gonic/gin"
)

func NewProductoRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	ec := &controller.ProductoController{
		ProductoRepository: &usecase.ProductoUseCase{},
	}
	ProductoRouter := group.Group("/producto")
	ProductoRouter.POST("/", ec.Create)
	ProductoRouter.GET("/", ec.Fetch)
	ProductoRouter.GET("/:id", ec.FetchById)
	ProductoRouter.PUT("/", ec.Update)
	ProductoRouter.DELETE("/:id", ec.Delete)
}
