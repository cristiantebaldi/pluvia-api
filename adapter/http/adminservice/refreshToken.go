package adminservice

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pluvia/pluvia-api/adapter/http/middleware"
)

func (service service) RefreshToken(c *gin.Context) {
	bearToken := c.Request.Header.Get("Authorization")
	refreshToken, typeToken, err := middleware.ParseBearerRefreshToken(bearToken)

	if err != nil {
		c.JSON(http.StatusUnauthorized, middleware.ResponseMessage(err))
		return
	}

	jwtTokenAuth, err := service.usecase.RefreshToken(*refreshToken, *typeToken)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, jwtTokenAuth)
}
