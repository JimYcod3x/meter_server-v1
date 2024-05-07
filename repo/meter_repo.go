package repo

import (
	"errors"

	"github.com/JimYcod3x/meter_server/domain"
	"github.com/JimYcod3x/meter_server/models"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type meterRepo struct {
	db *gorm.DB
	rdb *redis.Client
}

func (n *meterRepo) CreateMeter(createMeter models.Meter) error {
	if err := n.db.Create(&createMeter).Error; err != nil {
		return errors.New("internal server error: cannot create meter")
	}
	return nil
}

func NewMeterRepo(db *gorm.DB, rdb *redis.Client) domain.MeterRepo{
	return &meterRepo{
		db: db,
		rdb: rdb,
	}
}