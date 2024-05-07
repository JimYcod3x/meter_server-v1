package domain

import "github.com/JimYcod3x/meter_server/models"


type MeterRepo interface {
	CreateMeter(createMeter models.Meter) error
	// GetMeterById(id int) (models.Meter, error)
}

type MeterUseCase interface{
	CreateMeter(createMeter models.Meter) error
	// GetMeterById(id int) (models.Meter, error)
}
