package controller

import (
	"net/http"
	"strconv"

	"gorm-template/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PagoController struct {
	PagoRepository domain.PagoRepository
}

func (te *PagoController) Create(c *gin.Context) { //Hay que ingresar todos los datos necesarios para crear
	var Pago domain.Pago

	err := c.ShouldBind(&Pago)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	Pago.ID = uuid.New()

	err = te.PagoRepository.Create(c, Pago)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Pago created successfully",
	})
}

func (te *PagoController) Fetch(c *gin.Context) {
	Pagos, err := te.PagoRepository.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Pagos)
}

func (te *PagoController) FetchById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	Pagos, err := te.PagoRepository.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Pagos)
}

func (te *PagoController) Update(c *gin.Context) {
	updatedPago := &domain.Pago{}

	err := c.ShouldBind(updatedPago)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if updatedPago.ID == uuid.Nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "ID Pago is requiered to update"})
		return
	}

	err = te.PagoRepository.Update(c, *updatedPago)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Pago updated succesfully"})
}

func (te *PagoController) Delete(c *gin.Context) {
	PagoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	err = te.PagoRepository.Delete(c, PagoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Pago delete succesfully"})
}
