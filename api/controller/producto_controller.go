package controller

import (
	"net/http"
	"strconv"

	"gorm-template/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductoController struct {
	ProductoRepository domain.ProductoRepository
}

func (te *ProductoController) Create(c *gin.Context) { //Hay que ingresar todos los datos necesarios para crear
	var Producto domain.Producto

	err := c.ShouldBind(&Producto)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if Producto.NombreProducto == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Name is required"})
		return
	}

	Producto.ID = uuid.New()

	err = te.ProductoRepository.Create(c, Producto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Producto created successfully",
	})
}

func (te *ProductoController) Fetch(c *gin.Context) {
	Productos, err := te.ProductoRepository.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Productos)
}

func (te *ProductoController) FetchById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	Productos, err := te.ProductoRepository.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Productos)
}

func (te *ProductoController) Update(c *gin.Context) {
	updatedProducto := &domain.Producto{}

	err := c.ShouldBind(updatedProducto)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if updatedProducto.ID == uuid.Nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "ID Producto is requiered to update"})
		return
	}

	err = te.ProductoRepository.Update(c, *updatedProducto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Producto updated succesfully"})
}

func (te *ProductoController) Delete(c *gin.Context) {
	ProductoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	err = te.ProductoRepository.Delete(c, ProductoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Producto delete succesfully"})
}
