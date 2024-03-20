package dto

import (
	"fmt"
	"robot-app/internal/helper"
	"robot-app/internal/model"
	"strings"
)

type FindDeviceReq struct {
	Q                  string  `form:"q"`
	Sort               string  `form:"sort" validate:"omitempty,enum=id;created_at;updated_at;name;description;model;version;production_cost;production_date;color;camera;battery_size;battery_life;charging_time;charging_type;rating;sensor_x;width;height;weight;material;brand;display_size;gps;microphone;number_of_processors;warranty;speaker;resolution"`
	Limit              int     `form:"limit"`
	Page               int     `form:"page"`
	Name               *string `form:"name"`
	Description        *string `form:"description"`
	Model              *string `form:"model"`
	Version            *string `form:"version"`
	ProductionCost     *string `form:"production_cost"`
	ProductionDate     *string `json:"production_date" form:"production_date" validate:"omitempty,dateRange"`
	Color              *string `form:"color"`
	Camera             *string `form:"camera"`
	BatterySize        *string `form:"battery_size"`
	BatteryLife        *string `form:"battery_life"`
	ChargingTime       *string `form:"charging_time"` // min
	ChargingType       *string `form:"charging_type"` // TypeC
	Rating             *string `form:"rating"`        // 0 - 5 star
	SensorX            *string `form:"sensor_x"`
	Width              *string `form:"width"`
	Height             *string `form:"height"`
	Weight             *string `form:"weight"`
	Material           *string `form:"material"`
	Brand              *string `form:"brand"`             // apple
	PowerConsumption   *string `form:"power_consumption"` // watt 150,220
	Speed              *string `form:"speed"`             // km/h
	SupportApp         *bool   `form:"support_app"`
	Temperature        *string `form:"temperature"`
	CountryOfOrigin    *string `form:"country_of_origin"`
	SupplyVoltage      *string `form:"supply_voltage"` // 240
	DeepCamera         *string `form:"deep_camera"`
	Lidar              *string `form:"lidar"`        //
	StorageSize        *string `form:"storage_size"` // GB
	Processor          *string `form:"processor"`
	Bluetooth          *string `form:"bluetooth"` // 5.0
	Ram                *string `form:"ram"`
	Imei               *string `form:"imei"`
	RefreshRate        *string `form:"refresh_rate"` // 60
	DisplaySize        *string `form:"display_size"` // inch
	Gps                *bool   `form:"gps"`
	Microphone         *string `form:"microphone"`
	NumberOfProcessors *string `form:"number_of_processors"`
	Warranty           *string `form:"warranty"`
	Speaker            *string `form:"speaker"`
	Resolution         *string `form:"resolution"` // Full HD
}

func (f FindDeviceReq) GetFilters() (filters []*model.Filter) {
	if f.Name != nil {
		filters = append(filters, &model.Filter{
			Key:    "LOWER(name)",
			Method: "LIKE",
			Value:  fmt.Sprintf("%%%s%%", strings.ToLower(*f.Name)),
		})
	}

	if f.Description != nil {
		filters = append(filters, &model.Filter{
			Key:    "LOWER(description)",
			Method: "LIKE",
			Value:  fmt.Sprintf("%%%s%%", strings.ToLower(*f.Description)),
		})
	}

	if f.Model != nil {
		filters = append(filters, helper.FilterInString(*f.Model, "model"))
	}

	if f.Version != nil {
		filters = append(filters, helper.FilterInString(*f.Version, "version"))
	}

	if f.ProductionCost != nil {
		filters = append(filters, helper.FilterInInt(*f.ProductionCost, "production_cost"))
	}

	if f.Color != nil {
		filters = append(filters, helper.FilterInString(*f.Color, "color"))
	}

	if f.Camera != nil {
		filters = append(filters, helper.FilterInString(*f.Camera, "camera"))
	}

	if f.ProductionDate != nil {
		filters = append(filters, helper.FilterDateRange(*f.ProductionDate, "production_date")...)
	}

	if f.BatterySize != nil {
		filters = append(filters, helper.FilterInInt(*f.BatterySize, "battery_size"))
	}

	if f.BatteryLife != nil {
		filters = append(filters, helper.FilterInInt(*f.BatteryLife, "battery_life"))
	}

	if f.ChargingTime != nil {
		filters = append(filters, helper.FilterInInt(*f.ChargingTime, "charging_time"))
	}
	if f.ChargingType != nil {
		filters = append(filters, helper.FilterInString(*f.ChargingType, "charging_type"))
	}
	if f.Rating != nil {
		filters = append(filters, helper.FilterInFloat(*f.Rating, "rating"))
	}

	if f.SensorX != nil {
		filters = append(filters, helper.FilterInString(*f.SensorX, "sensor_x"))
	}

	if f.Width != nil {
		filters = append(filters, helper.FilterInFloat(*f.Width, "width"))
	}

	if f.Height != nil {
		filters = append(filters, helper.FilterInFloat(*f.Height, "height"))
	}

	if f.Weight != nil {
		filters = append(filters, helper.FilterInFloat(*f.Weight, "weight"))
	}
	if f.Material != nil {
		filters = append(filters, helper.FilterInString(*f.Material, "material"))
	}

	if f.Brand != nil {
		filters = append(filters, helper.FilterInString(*f.Brand, "brand"))
	}

	if f.PowerConsumption != nil {
		filters = append(filters, helper.FilterInFloat(*f.PowerConsumption, "power_consumption"))
	}

	if f.Speed != nil {
		filters = append(filters, helper.FilterInFloat(*f.Speed, "speed"))
	}

	if f.SupportApp != nil {
		filters = append(filters, helper.FilterInBool(*f.SupportApp, "support_app"))
	}

	if f.Temperature != nil {
		filters = append(filters, helper.FilterInInt(*f.Temperature, "temperature"))
	}

	if f.CountryOfOrigin != nil {
		filters = append(filters, helper.FilterInString(*f.CountryOfOrigin, "country_of_origin"))
	}

	if f.SupplyVoltage != nil {
		filters = append(filters, helper.FilterInInt(*f.SupplyVoltage, "supply_voltage"))
	}

	if f.DeepCamera != nil {
		filters = append(filters, helper.FilterInString(*f.DeepCamera, "deep_camera"))
	}

	if f.Lidar != nil {
		filters = append(filters, helper.FilterInString(*f.Lidar, "lidar"))
	}

	if f.StorageSize != nil {
		filters = append(filters, helper.FilterInInt(*f.StorageSize, "storage_size"))
	}

	if f.Processor != nil {
		filters = append(filters, helper.FilterInString(*f.Processor, "processor"))
	}

	if f.Bluetooth != nil {
		filters = append(filters, helper.FilterInString(*f.Bluetooth, "bluetooth"))
	}

	if f.Ram != nil {
		filters = append(filters, helper.FilterInInt(*f.Ram, "ram"))
	}

	if f.Imei != nil {
		filters = append(filters, helper.FilterInString(*f.Imei, "imei"))
	}

	if f.RefreshRate != nil {
		filters = append(filters, helper.FilterInInt(*f.RefreshRate, "refresh_rate"))
	}

	if f.DisplaySize != nil {
		filters = append(filters, helper.FilterInFloat(*f.DisplaySize, "display_size"))
	}

	if f.Gps != nil {
		filters = append(filters, helper.FilterInBool(*f.Gps, "gps"))
	}

	if f.Microphone != nil {
		filters = append(filters, helper.FilterInString(*f.Microphone, "microphone"))
	}

	if f.NumberOfProcessors != nil {
		filters = append(filters, helper.FilterInInt(*f.NumberOfProcessors, "number_of_processors"))
	}

	if f.Warranty != nil {
		filters = append(filters, helper.FilterInInt(*f.Warranty, "warranty"))
	}

	if f.Speaker != nil {
		filters = append(filters, helper.FilterInString(*f.Speaker, "speaker"))
	}

	if f.Resolution != nil {
		filters = append(filters, helper.FilterInString(*f.Resolution, "resolution"))
	}

	return filters
}
