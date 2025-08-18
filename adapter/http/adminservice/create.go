package adminservice

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pluvia/pluvia-api/core/dto"
	"github.com/pluvia/pluvia-api/util"
)

// Create goDoc
// @Summary Create admin
// @Description Create admin
// @Tags admin
// @Accept  json
// @Produce  json
// @Param admin body dto.AdminRequestBody true "admin"
// @Success 200 {object} domain.Admin
// @Security ApiKeyAuth
// @Router /admin [post]
func (service service) Create(c *gin.Context) {
	adminRequest, err := dto.FromJsonToAdminRequestBody(c.Request.Body, service.validator)

	if err != nil {
		if _, ok := err.(util.RequestError); ok {
			c.JSON(http.StatusUnprocessableEntity, err.(util.RequestError))
			return
		}

		c.JSON(http.StatusInternalServerError, middleware.ResponseMessage(err))
		return
	}

	adminContext := middleware.GetContextData(c.Request)

	admin, err := service.usecase.Create(adminRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, admin)
}
