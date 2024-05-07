package meter

import (
	"encoding/hex"
	"fmt"

	"github.com/JimYcod3x/meter_server/internal/utils"
)

type meterDS interface {
	GetDataFromMeter(meterType, meterDSCmd int, meterID string, dsCmdParam byte) []byte
	SwitchCtrl()
	OtherCtrl()
	BRouteMeterInfo()
	RS485Ctrl()
	OTAUpCommModule()
	OTAUpWiFi()
	OTATSLHTTPSCertKeyDload()
	OTAUpMeterFirm()
	OTAUpBootloader()
}

func  CallGetDataFromMeter(meterType, meterDSCmd int, meterID string, dsCmdParam ...string) []byte {
	prefix := utils.SentOutPrefix(meterType, meterDSCmd, meterID)
	var commandParam = ""
	switch dsCmdParam[0] {
	case "GetInstantReading":
		commandParam = WithOutAdditionParam(dsCmdParam[0])
	case "Get30minRecordLog":
		commandParam = Get30minRecordLog(dsCmdParam[0], dsCmdParam[1], dsCmdParam[2])
	case "GetEventLogRecords":
	case "GetAvgInstantPowerLineValue":
	case "GetMeterID&VerInfo":
	case "Reserved":
	case "GetDiagnosticInfo":
	case "GetCommModuleProgramFlashCRC":
	case "GetMeterProgramFashCRC":
	}
	preEncry := prefix + GetParam(commandParam)
	dataKey, _ := utils.GetSerDataKey(meterID)
	preEncyhexStr, _ := hex.DecodeString(preEncry)
	fmt.Println("preEncrypt payload in testmeterfn", preEncry)
	fmt.Println("preEncrypt payload in testmeterfn", preEncyhexStr)
	sentouPayload, _ := utils.EncryptPadding(preEncry, dataKey)
	return sentouPayload                                       
}


func WithOutAdditionParam(dsCmdParam string) string{
	return utils.HexByteToHexStr(DSCommandSet.GetDataFromMeter[dsCmdParam])
}

func Get30minRecordLog(dsCmdParam, start, end string) string{
	commandParam := utils.HexByteToHexStr(DSCommandSet.GetDataFromMeter[dsCmdParam])
	startHexStr := utils.ParseTimeToHexStr(start)
	endHexStr := utils.ParseTimeToHexStr(end)
	return commandParam + startHexStr + endHexStr
	
}