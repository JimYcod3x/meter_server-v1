package utils

import (
	"fmt"
	"log"
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
	return  fromDBMeterID == meterID 
}

func FindDateKey(meterID string) (dataKey string, err error) {
	dataKey = "000000" + meterID
	fmt.Println("db data key: ", dataKey)
	return 
}

func FindMasterKey(meterID string) (masterKey string, err error) {
	masterKey =  meterID + "000000"
	fmt.Println("db master Key: ", masterKey)
	return 
}



