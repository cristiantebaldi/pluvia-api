package administradorservice

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pluvia/pluvia-api/adapter/http/middleware"
	"github.com/pluvia/pluvia-api/core/dto"
	"github.com/pluvia/pluvia-api/util"
)

// Create goDoc
// @Summary Create administrador
// @Description Create administrador
// @Tags administrador
// @Accept  json
// @Produce  json
// @Param administrador body dto.AdministradorRequestBody true "administrador"
// @Success 200 {object} domain.Administrador
// @Security ApiKeyAuth
// @Router /administrador [post]
func (service service) Create(c *gin.Context) {
	administradorRequest, err := dto.FromJsonToAdminRequestBody(c.Request.Body, service.validator)

	if err != nil {
		if _, ok := err.(util.RequestError); ok {
			c.JSON(http.StatusUnprocessableEntity, err.(util.RequestError))
			return
		}

		c.JSON(http.StatusInternalServerError, middleware.ResponseMessage(err))
		return
	}

	administrador, err := service.usecase.Create(administradorRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, administrador)
}
