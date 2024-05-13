package utils

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/JimYcod3x/meter_server/models"
	"github.com/go-redis/redis/v8"
	"github.com/qustavo/dotsql"
)

var(
	meters models.Meter
	pwd, _ = os.Getwd()
	dot, err = dotsql.LoadFromFile(pwd + "/sql/meter.sql")
) 
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

func InitDB(db *sql.DB) {
	fmt.Println("Init the meter db")
	res, err := dot.Exec(db, "create-database")
	checkErr(err)
	fmt.Println(res.LastInsertId())
	res, err = dot.Exec(db, "switch-to-database")
	checkErr(err)
	fmt.Println(res.LastInsertId())
	res, err = dot.Exec(db, "create-meter-table")
	checkErr(err)
	fmt.Println(res.LastInsertId())
}

func CreateMeter(db *sql.DB, meterID string, meterType string) error {
	fmt.Println("create the meter")

	if err != nil {
		fmt.Println("can not load the sql file", err)
	}
	res, _ := dot.QueryRow(db, "find-one-meter-by-meter_id", meterID)
	err = res.Scan(&meters.MeterID)
	if err == sql.ErrNoRows {
		_, err := dot.Exec(db, "create-meter", meterID, meterType)
		if err != nil {
			fmt.Println("can not create the meter", err)
			return err
		}
		return nil
	}
	


	return err
}

func SaveToRdb(rdb *redis.Client, ctx context.Context, meterID string, key string) error {
	err := rdb.Set(ctx, "mk_"+meterID, key, 24*time.Hour).Err()
	if err != nil {
		fmt.Println("can not save to rdb", err)
	}
	return err
}

func UpdateKeyToDb(db *sql.DB, key string, args ...any) error {
	stmt, err := db.Prepare("UPDATE meter SET " + key + " = ? WHERE mete_id = ?")
	if err != nil {
		fmt.Println("prepare the " + key + " to db", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(key)
	if err != nil {
		fmt.Println("can not save " + key + " to db")
		return err
	}
	return nil
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}