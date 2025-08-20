package adminservice

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pluvia/pluvia-api/adapter/http/middleware"
)

// Delete goDoc
// @Summary Delete admin
// @Description Delete admin by id
// @Tags admin
// @Accept  json
// @Produce  json
// @Param id path int true "1"
// @Success 200 {object} domain.Admin
// @Security ApiKeyAuth
// @Router /admin/{id} [delete]
func (service service) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			middleware.ResponseMessage(fmt.Errorf("param route id is required and must be valid number")),
		)
		return
	}

	admin, err := service.usecase.Delete(int32(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, admin)
}
