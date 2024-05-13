package models

import "time"

type Meter struct {
	MeterID string
	MeterType string
	MasterKey string
	DataKey string
	CreatedAt time.Time
	UpdatedAt time.Time
}

