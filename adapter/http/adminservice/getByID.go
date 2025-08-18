package adminservice

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetByID goDoc
// @Summary Get admin
// @Description Get admin by id
// @Tags admin
// @Accept  json
// @Produce  json
// @Param id path int true "1"
// @Success 200 {object} domain.Admin
// @Security ApiKeyAuths
// @Router /admin/{id} [get]
func (service service) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			middleware.ResponseMessage(fmt.Errorf("param route id is required and must be valid number")),
		)
		return
	}

	admin, err := service.usecase.(int32(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, admin)
}
