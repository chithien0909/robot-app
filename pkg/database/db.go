package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"robot-app/internal/model"
	"time"
)

func New(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Migrate(db *gorm.DB) {
	err := db.Migrator().AutoMigrate(&model.Device{})
	if err != nil {
		log.Printf("Failed to migrate: %s\n", err.Error())
	}
}

func Seed(db *gorm.DB) {
	var total int64
	db.Model(&model.Device{}).Count(&total)
	if total > 0 {
		return
	}

	colorRed := "Red"
	colorBlue := "Blue"
	camera := "12MP"
	sensorX := "X100"
	batterySize := 7000
	batteryLife := 5
	chargingTime := 3
	var rating float32 = 4.5
	var width = 100.0
	var height = 200.0
	var weight = 20.0
	var bluetooth = "5.0"

	devices := []*model.Device{
		{
			Name:           "Aga X",
			Description:    "Robot #1",
			Model:          "X",
			Version:        "b1.2",
			ProductionCost: 200,
			ProductionDate: time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC),
			Color:          &colorRed,
			Camera:         &camera,
			SensorX:        &sensorX,
			BatterySize:    &batterySize,
			BatteryLife:    &batteryLife,
			ChargingTime:   &chargingTime,
			Rating:         &rating,
			Width:          &width,
			Height:         &height,
			Weight:         &weight,
			Bluetooth:      &bluetooth,
		},
		{
			Name:           "Udo",
			Description:    "Robot #2",
			Model:          "Y",
			Version:        "a0.3",
			ProductionCost: 155,
			ProductionDate: time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC),
			Color:          &colorBlue,
			Camera:         &camera,
			SensorX:        &sensorX,
			BatterySize:    &batterySize,
			BatteryLife:    &batteryLife,
			ChargingTime:   &chargingTime,
			Rating:         &rating,
			Width:          &width,
			Height:         &height,
			Weight:         &weight,
			Bluetooth:      &bluetooth,
		},
		{
			Name:           "KimZi",
			Description:    "Robot #3",
			Model:          "Z",
			Version:        "1.5",
			ProductionCost: 155,
			ProductionDate: time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC),
			Color:          &colorBlue,
			Camera:         &camera,
			BatterySize:    &batterySize,
			BatteryLife:    &batteryLife,
			ChargingTime:   &chargingTime,
			Rating:         &rating,
			Width:          &width,
			Height:         &height,
			Weight:         &weight,
			Bluetooth:      &bluetooth,
		},
	}
	err := db.Model(&model.Device{}).Create(&devices)
	if err.Error != nil {
		log.Printf("Failed database.Seed: %s\n", err.Error)
	}
}
