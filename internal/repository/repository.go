package repository

import "gorm.io/gorm"

type Repository interface {
	Device() DeviceRepository
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{
		device: NewDeviceRepository(db),
	}
}

type repositoryImpl struct {
	device DeviceRepository
}

func (r *repositoryImpl) Device() DeviceRepository {
	return r.device
}
