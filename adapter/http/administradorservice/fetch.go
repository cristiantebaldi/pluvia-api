package administradorservice

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pluvia/pluvia-api/adapter/http/middleware"
)

// Fetch goDoc
// @Summary Fetch admins
// @Description Fetch admins
// @Tags administrador
// @Accept  json
// @Produce  json
// @Success 200 {object} domain.Administrador
// @Security ApiKeyAuth
// @Router /administrador [get]
func (service service) Fetch(c *gin.Context) {
	//params, err := dto.FromValuePaginationRequestParams(c.Request)

	// if err != nil {
	// 	if _, ok := err.(util.RequestError); ok {
	// 		c.JSON(http.StatusUnprocessableEntity, err.(util.RequestError))
	// 		return
	// 	}

	// 	c.JSON(http.StatusInternalServerError, middleware.ResponseMessage(err))
	// 	return
	// }

	administrador := middleware.GetContextData(c.Request)
	admins, err := service.usecase.Fetch(administrador)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, admins)
}
