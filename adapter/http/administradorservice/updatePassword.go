package administradorservice

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pluvia/pluvia-api/adapter/http/middleware"
	"github.com/pluvia/pluvia-api/core/dto"
	"github.com/pluvia/pluvia-api/util"
)

// UpdatePassword goDoc
// @Summary UpdatePassword administrador
// @Description UpdatePassword administrador by id
// @Tags administrador
// @Accept  json
// @Produce  json
// @Param administrador body dto.AdministradorRequestBody true "administrador"
// @Param id path int true "1"
// @Success 200 {object} domain.Administrador
// @Security ApiKeyAuth
// @Router /administrador-pass/{id} [put]
func (service service) UpdatePassword(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			middleware.ResponseMessage(fmt.Errorf("param route id is required and must be valid number")),
		)
		return
	}

	accountRequest, err := dto.FromJsonToAdminRequestBody(c.Request.Body, service.validator)

	if err != nil {
		if _, ok := err.(util.RequestError); ok {
			c.JSON(http.StatusUnprocessableEntity, err.(util.RequestError))
			return
		}

		c.JSON(http.StatusInternalServerError, middleware.ResponseMessage(err))
		return
	}

	account, err := service.usecase.UpdatePassword(int32(id), accountRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, account)
}
