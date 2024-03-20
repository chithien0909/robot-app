package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"robot-app/dto"
	"robot-app/internal/helper"
	"robot-app/internal/repository"
	"robot-app/validation"
)

type DeviceHandler interface {
	Find(c *gin.Context)
}

type deviceHandlerImpl struct {
	validate *validation.StructValidation
	repo     repository.Repository
}

func NewDeviceHandler(v *validation.StructValidation, repo repository.Repository) DeviceHandler {
	return &deviceHandlerImpl{
		validate: v,
		repo:     repo,
	}
}

func (d *deviceHandlerImpl) Find(c *gin.Context) {

	body := dto.FindDeviceReq{}

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := d.validate.Validate(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	filters := body.GetFilters()

	pagination := helper.BuildPagination(body.Page, body.Limit)

	query := helper.BuildQuery(body.Q, filters, helper.BuildSortFromString(body.Sort), pagination)

	devices, total, err := d.repo.Device().Find(c, query)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    devices,
		"total":   total,
		"page":    pagination.Page,
		"limit":   pagination.Limit,
	})
}