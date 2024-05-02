package meter

import (
	"fmt"

	"github.com/JimYcod3x/meter_server/internal/utils"
	"github.com/mochi-mqtt/server/v2/packets"
)

func KeyXFn(pk packets.Packet, DecryptKey string, usCmdParam string) string {
	fmt.Println("plainPayload: ", pk.Payload)
	decryptPayload, _ := utils.DecryptByte(pk.Payload, DecryptKey)
	meterType := utils.GetMeterTypeFromPayload(decryptPayload)
	commandType := utils.GetUSCommandFromDecrypt(decryptPayload)
	getMeterID := utils.GetMeterID(pk.TopicName)
	fmt.Println("decryptPayload: ", decryptPayload)
	fmt.Println(meterType, commandType, getMeterID, decryptPayload)
	// fmt.Println("Down stream: ", getDSData(meterType, commandType, getMeterID, decryptPayload))
	fmt.Println("debugPrint: ", meterType, commandType, getMeterID, decryptPayload)
	return getKeyXDSKeyXData(meterType, getMeterID, decryptPayload, usCmdParam)
}

func getKeyXDSKeyXData(meterTypeInt int, getMeterID string, decryptPayload []byte, usCmdParam string) (DSData string) {

	switch meterTypeInt {
	case ElectricityMeter:
		return getKeyXDSCommandData(getMeterID, decryptPayload, usCmdParam)
	// case WaterMeter:
	case IoT:
		fmt.Println("This is IoT type")
		return getKeyXDSCommandData(getMeterID, decryptPayload, usCmdParam)
	// case GasMeter:
	// case HeatMeter:
	case PV:
		return getKeyXDSCommandData(getMeterID, decryptPayload, usCmdParam)
	}
	return
}

func getKeyXDSCommandData(getMeterID string, decryptPayload []byte, usCmdParam string) (DSCommandData string) {
	switch usCmdParam {
	case "ReqRegistration":
		fmt.Println("ReqRegistration")
		return RespondRegister(getMeterID, decryptPayload)
	case "ReqChangeKey":
		fmt.Println("ReqChangeKey")
		return ResponseNewDataKey(getMeterID, decryptPayload)
	case "ReqSucACK":
		fmt.Println("ReqSucACK")
		return ConfirmNewDataKey(getMeterID, decryptPayload)
	}
	return
}

func RespondRegister(getMeterID string, decryptPayload []byte) (DSCommandData string) {
	commandParam := utils.HexByteToHexStr(DSCommandSet.ExchangeKey["ChangeRegistrationData"])
	masterKey := getMeterID + "000000"
	dataKey := "000000" + getMeterID
	masterKey = utils.SecurityKeytoHex(masterKey)
	dataKey = utils.SecurityKeytoHex(dataKey)
	registrationData := masterKey + dataKey
	dataLength := fmt.Sprintf("%02x", len(registrationData)/2)
	fmt.Println("dataLength: ", len(registrationData), dataLength)
	DSCommandData = getPreData(decryptPayload, commandParam, dataLength, registrationData)
	fmt.Println("RespondRegister printout: ", DSCommandData)
	return
}

func ResponseNewDataKey(getMeterID string, decryptPayload []byte) (DSCommandData string) {
	commandParam := utils.HexByteToHexStr(DSCommandSet.ExchangeKey["ChangeNewKey"])
	dataKey := "000000" + getMeterID
	dataKey = utils.SecurityKeytoHex(dataKey)
	DSCommandData = getPreData(decryptPayload, commandParam, dataKey)
	fmt.Println("ResponseNewDataKey printout: ", DSCommandData)
	return
}

func ConfirmNewDataKey(getMeterID string, decryptPayload []byte) (DSCommandData string) {
	commandParam := utils.HexByteToHexStr(DSCommandSet.ExchangeKey["ConfirmNewKey"])
	DSCommandData = getPreData(decryptPayload, commandParam)
	fmt.Println("ConfirmNewDataKey printout: ", DSCommandData)
	return
}