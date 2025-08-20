package adminservice

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pluvia/pluvia-api/adapter/http/middleware"
	"github.com/pluvia/pluvia-api/core/dto"
	"github.com/pluvia/pluvia-api/util"
)

// Update goDoc
// @Summary Update admin
// @Description Update admin by id
// @Tags admin
// @Accept  json
// @Produce  json
// @Param admin body dto.AdminRequestBody true "admin"
// @Param id path int true "1"
// @Success 200 {object} domain.Admin
// @Security ApiKeyAuth
// @Router /admin/{id} [put]
func (service service) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			middleware.ResponseMessage(fmt.Errorf("param route id is required and must be valid number")),
		)
		return
	}

	adminRequest, err := dto.FromJsonToAdminRequestBody(c.Request.Body, service.validator)

	if err != nil {
		if _, ok := err.(util.RequestError); ok {
			c.JSON(http.StatusUnprocessableEntity, err.(util.RequestError))
			return
		}

		c.JSON(http.StatusInternalServerError, middleware.ResponseMessage(err))
		return
	}

	admin, err := service.usecase.Update(int32(id), adminRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, admin)
}
