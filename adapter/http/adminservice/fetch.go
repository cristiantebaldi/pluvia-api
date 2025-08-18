package adminservice

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Fetch goDoc
// @Summary Fetch admins
// @Description Fetch admins
// @Tags admin
// @Accept  json
// @Produce  json
// @Success 200 {object} domain.Admin
// @Security ApiKeyAuth
// @Router /admin [get]
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

	//admin := middleware.GetContextData(c.Request)
	admins, err := service.usecase.Fetch()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, admins)
}
