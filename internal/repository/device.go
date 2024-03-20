package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"robot-app/internal/helper"
	"robot-app/internal/model"
	"strings"
)

type DeviceRepository interface {
	Find(ctx context.Context, query *model.Query) (devices []*model.Device, total int64, err error)
}

type deviceRepositoryImpl struct {
	db *gorm.DB
}

func NewDeviceRepository(db *gorm.DB) DeviceRepository {
	return &deviceRepositoryImpl{
		db: db,
	}
}

func (d *deviceRepositoryImpl) Find(ctx context.Context, query *model.Query) (devices []*model.Device, total int64, err error) {

	fields, values := helper.BuildFilters(query.Filters)

	tx := d.db.Model(&model.Device{}).Debug().
		WithContext(ctx).
		Where(strings.Join(fields, " AND "), values...)

	if query.Sort != nil {
		tx.Order(fmt.Sprintf("%s %s", query.Sort.Key, query.Sort.SortBy))
	}

	if err := tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	pagin := query.GetPagination()

	tx = tx.
		Offset(pagin.Offset).
		Limit(pagin.Limit)

	if err := tx.Find(&devices).Error; err != nil {
		return nil, 0, err
	}

	return devices, total, nil
}
