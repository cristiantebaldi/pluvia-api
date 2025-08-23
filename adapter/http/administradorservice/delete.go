package administradorservice

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pluvia/pluvia-api/adapter/http/middleware"
)

// Delete goDoc
// @Summary Delete administrador
// @Description Delete administrador by id
// @Tags administrador
// @Accept  json
// @Produce  json
// @Param id path int true "1"
// @Success 200 {object} domain.Administrador
// @Security ApiKeyAuth
// @Router /administrador/{id} [delete]
func (service service) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			middleware.ResponseMessage(fmt.Errorf("param route id is required and must be valid number")),
		)
		return
	}

	administrador, err := service.usecase.Delete(int32(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, administrador)
}
