package handler

import (
	"robot-app/internal/repository"
	"robot-app/validation"
)

type Handler interface {
	DeviceHandler() DeviceHandler
}

type handlerImpl struct {
	deviceHandler DeviceHandler
}

func New(validator *validation.StructValidation, repo repository.Repository) Handler {
	return &handlerImpl{
		deviceHandler: NewDeviceHandler(validator, repo),
	}
}

func (h *handlerImpl) DeviceHandler() DeviceHandler {
	return h.deviceHandler
}
