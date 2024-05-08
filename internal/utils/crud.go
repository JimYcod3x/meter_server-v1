package utils

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/JimYcod3x/meter_server/models"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var meters models.Meter

func GetSerDataKey(meterID string) (dataKey string, found bool) {
	dataKey, err := FindDateKey(meterID)
	found = false
	if err != nil {
		log.Fatal("can not find the data key in db")
	}
	found = true
	return
}

func GetSerMasterKey(meterID string) (masterKey string, found bool) {
	masterKey, err := FindMasterKey(meterID)
	found = false
	if err != nil {
		log.Fatal("can not find the master key in db")
	}
	found = true
	return
}

func GetMeterIDFromDB(meterID string) bool {
	var fromDBMeterID = meterID
	return fromDBMeterID == meterID
}

func FindDateKey(meterID string) (dataKey string, err error) {
	dataKey = "000000" + meterID
	fmt.Println("db data key: ", dataKey)
	return
}

func FindMasterKey(meterID string) (masterKey string, err error) {
	masterKey = meterID + "000000"
	fmt.Println("db master Key: ", masterKey)
	return
}

func CreateMeter(db *gorm.DB, meterID string, meterType string) error {
	fmt.Println("create the meter")
	meters := models.Meter{
		MeterID:   meterID,
		MeterType: meterType,
	}
	err := db.Where("meter_id = ?", meterID).First(&meters).Error
	if err != nil {
		fmt.Println(err)
		return db.Model(models.Meter{}).Create(&meters).Error
	}
	fmt.Println(meters)
	// if err := db.First(&models.User{Email: payload.Email}).Error; errors.Is(err, gorm.ErrRecordNotFound) {

	// }
	return nil
}

func SaveToRdb(rdb *redis.Client, ctx context.Context, meterID string, key string) error {
	err := rdb.Set(ctx, "mk_"+meterID, key, 24*time.Hour).Err()
	if err != nil {
		fmt.Println("can not save to rdb", err)
	}
	return err
}
