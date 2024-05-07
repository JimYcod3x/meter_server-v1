package main

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	"github.com/JimYcod3x/meter_server/internal/meter"
	"github.com/JimYcod3x/meter_server/internal/utils"
)

var needEncrypt = "604a235000007802"

// var testmasterKey = "J23P000078000000"
var testmasterKey = "000000J230008542"
// var testmasterKey = "000000J23P000078"

// var testmasterKey = "69aF7&3KY0_kk89@"
func main() {
	payload := "4ac487b681a62ef7fdf1cab3e6fa038d"
	fmt.Println(testdecrypt(payload))
	fmt.Println(testEncrypt(needEncrypt))
	testSwitchCase("GetInstantReading")
	testParseTime("2024-05-02T00:00:02")
	fmt.Println(hex.EncodeToString(testFn2()))
	testinter("J200002335")
	c := time.Tick(5 * time.Second)
	for next := range c {
		fmt.Printf("%v, %v\n", next, statusUpdate())
	}
}

func testdecrypt(payload string) string {
	decryptByte, _ := hex.DecodeString(payload)
	testdecrypt, _ := utils.DecryptByte(decryptByte, testmasterKey)
	fmt.Println("testDecrypt byte[]", testdecrypt)
	return hex.EncodeToString(testdecrypt)
}

func testEncrypt(needEncrypt string) string {
	encreyptByte, _ := utils.EncryptPadding(needEncrypt, testmasterKey)
	return hex.EncodeToString(encreyptByte)
}

func testSwitchCase(a string) {
	switch a {
	case "GetInstantReading":
		a = utils.HexByteToHexStr(meter.DSCommandSet.GetDataFromMeter[a])
	}
	fmt.Println(a)
}

func testParseTime(timeStr any) string {
	var unixTimeStamp int64
	switch timeVal := timeStr.(type) {
	case string:
		parseTime, err := time.Parse("2006-01-02T15:04:05", timeVal)
		if err != nil {
			fmt.Println("Error parsing time: ", err)
			return ""
		}
		unixTimeStamp = parseTime.Unix()

	case int64:
		unixTimeStamp = timeVal
	default:
		fmt.Println("unsupported type")
		return ""
	}

	hexString := strconv.FormatInt(unixTimeStamp, 16)
	fmt.Println(hexString)
	return hexString
}

func testFn2() []byte {
	meterType := meter.ElectricityMeter
	meterDSCmd := meter.GetDataFromMeter
	meterID := "J230008542"
	// meterID := "J23P000078"
	dsCmdParam := "Get30minRecordLog"
	startTime := "2024-05-02T15:15:00"
	endTime := "2024-05-02T15:15:01"
	return meter.CallGetDataFromMeter(meterType, meterDSCmd, meterID, dsCmdParam, startTime, endTime)
}

func testinter(id string) {
	a := utils.MeterIDtoHex(id)
	fmt.Println(a)
}

func statusUpdate() string {return "updated"}