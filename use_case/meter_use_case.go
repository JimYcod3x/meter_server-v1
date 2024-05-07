package usecase

import (
	"github.com/JimYcod3x/meter_server/domain"
	"github.com/JimYcod3x/meter_server/models"
)

type meterUseCase struct {
	meterRepo domain.MeterRepo
}

func (n *meterUseCase) CreateMeter(createMeter models.Meter) error {
	err := n.meterRepo.CreateMeter(createMeter)
		return err
}

func NewMeterUseCase(meterRepo domain.MeterRepo) domain.MeterUseCase {
	return &meterUseCase{
		meterRepo: meterRepo,
	}
}

