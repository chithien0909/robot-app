package model

import "time"

type Device struct {
	ID                 int       `json:"id" gorm:"primaryKey"`
	CreatedAt          time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt          time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	Model              string    `json:"model"`
	Version            string    `json:"version"`
	ProductionCost     int       `json:"production_cost"`
	ProductionDate     time.Time `json:"production_date"`
	Color              *string   `json:"color"`
	Camera             *string   `json:"camera"`
	BatterySize        *int      `json:"battery_size"`
	BatteryLife        *int      `json:"battery_life"`
	ChargingTime       *int      `json:"charging_time"` // min
	ChargingType       *string   `json:"charging_type"` // TypeC
	Rating             *float32  `json:"rating"`        // 0 - 5 star
	SensorX            *string   `json:"sensor_x"`
	Width              *float64  `json:"width"`
	Height             *float64  `json:"height"`
	Weight             *float64  `json:"weight"`
	Material           *string   `json:"material"`
	Brand              *string   `json:"brand"` // apple
	PowerConsumption   *int      `json:"power"` // watt 150,220
	Speed              *float64  `json:"speed"` // km/h
	SupportApp         *bool     `json:"support_app"`
	Temperature        *int      `json:"temperature"`
	CountryOfOrigin    *string   `json:"country_of_origin"`
	SupplyVoltage      *int      `json:"supply_voltage"` // 240
	DeepCamera         *string   `json:"deep_camera"`
	Lidar              *string   `json:"lidar"`        //
	StorageSize        *int      `json:"storage_size"` // GB
	Processor          *string   `json:"processor"`
	Bluetooth          *string   `json:"bluetooth"` // 5.0
	Ram                *int      `json:"ram"`
	Imei               *string   `json:"imei"`
	RefreshRate        *int      `json:"refresh_rate"` // 60
	DisplaySize        *float32  `json:"display_size"` // inch
	Gps                *bool     `json:"gps"`
	Microphone         *string   `json:"microphone"`
	NumberOfProcessors *int      `json:"number_of_processors"`
	Warranty           *int      `json:"warranty"`
	Speaker            *string   `json:"speaker"`
	Resolution         *int      `json:"resolution"` // Full HD
}

func (Device) TableName() string {
	return "devices"
}
