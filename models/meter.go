package models

import "gorm.io/gorm"

type Meter struct {
	ID string `gorm:"unique;primaryKey;autoIncrement"`
	MeterID string `gorm:"type:varchar(50);unique" json:"meterId"`
	MeterType string `gorm:"type:varchar(50);not null" json:"meterType"`
	MasterKey string `gorm:"type:varchar(50)"`
	DataKey string `gorm:"type:varchar(50)"`
	*gorm.Model
}

func (m *Meter) TableName() string{
	return "meter"
}