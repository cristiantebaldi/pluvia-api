package administradorservice

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pluvia/pluvia-api/adapter/http/middleware"
	"github.com/pluvia/pluvia-api/core/dto"
	"github.com/pluvia/pluvia-api/util"
)

func (service service) GetByLoginPassword(c *gin.Context) {
	loginRequest, err := dto.FromJsonToAdminLoginRequestBody(c.Request.Body, service.validator)

	if err != nil {
		if _, ok := err.(util.RequestError); ok {
			c.JSON(http.StatusUnprocessableEntity, err.(util.RequestError))
			return
		}

		c.JSON(http.StatusInternalServerError, middleware.ResponseMessage(err))
		return
	}

	jwtTokenAuth, err := service.usecase.GetByLoginPassword(loginRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, jwtTokenAuth)
}
