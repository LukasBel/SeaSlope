package models

import "gorm.io/gorm"

type SeaSlopeData struct {
	ID          uint8       `gorm:"primaryKey" json:"id"`
	Conditions  Conditions  `json:"conditions"`
	WeatherData WeatherData `json:"weatherData"`
	Forecast    Forecast    `json:"forecast"`
}

func MigrateSpots(db *gorm.DB) error {
	err := db.AutoMigrate(&SeaSlopeData{})
	if err != nil {
		return err
	}
	return nil
}
