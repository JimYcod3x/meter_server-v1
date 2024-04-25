package meter

import (
	"fmt"
	"log"

	"github.com/JimYcod3x/meter_server/internal/utils"
	"github.com/mochi-mqtt/server/v2/packets"
)

// func PublishData(plainData string, sev *mqtt.Server) error {
// 	err := sev.Publish(topic string, payload []byte, retain bool, qos byte)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func EncryptPlainData(plainData string) (encryptedData []byte) {
	encryptedData, _ = utils.EncryptPadding(plainData, DefaultKey)
	return
}

func ExchangeKeyFn(pk packets.Packet) string {
	getID := utils.GetMeterID(pk.TopicName)
	fmt.Println("getID test: ", getID)
	// DecryptKey := utils.GetSerDataKey(getID)
	decryptPayload, err := utils.DecryptByte(pk.Payload, DefaultKey)
	if err != nil{
		log.Fatal("can not decrypt")
	}
	masterKey := utils.SecurityKeytoHex(utils.GetSerMasterKey(getID))
	dataKey := utils.SecurityKeytoHex(utils.GetSerDataKey(getID))
	securityKey := masterKey + dataKey

	lengthOfKey := fmt.Sprintf("%02x", len(securityKey)/2)
	prefix, id, commandParam := getPrefix(getID, decryptPayload)
	predata := prefix + id + GetParam(commandParam, lengthOfKey, securityKey)
	return predata
}

func GetParam(a ...string) string {
	return DoPredata(a)
}

func DoPredata(a []string) string {
	var v string
	fmt.Println("============================")
	for i := 0; i < len(a); i++ {
		v += a[i]
	}
	fmt.Println("value: ", v)

	return v
}

func getPrefix(getMeterID string, decryptPayload []byte) (prefix string, id string, commandParam string) {
	meterTypePayload := utils.GetMeterTypeFromPayload(decryptPayload)
	meterTypeInt := int(decryptPayload[0] >> 5)
	commandInt := int(decryptPayload[0] & 0x1f)
	meterTypeBinaryStr := utils.IntToBinStr(meterTypePayload, MeterTypeBit)
	meterCommandBinaryStr := utils.IntToBinStr(ExchangeKey, MeterCommandBit)
	prefix, _ = utils.BinaryToHex(meterTypeBinaryStr + meterCommandBinaryStr)
	id, _ = utils.FindIdHaveF(decryptPayload)
	commandParam = getDSData(meterTypeInt, commandInt, getMeterID, decryptPayload)
	fmt.Printf("prefix: %s id: %s commandParam: %s\n", prefix, id, commandParam)
	return
}

func DataTXFn(pk packets.Packet, DecryptKey string) string {
	decryptPayload, _ := utils.DecryptByte(pk.Payload, DecryptKey)
	meterType := int(decryptPayload[0] >> 5)
	commandType := int(decryptPayload[0] & 0x1f)
	getMeterID := utils.GetMeterID(pk.TopicName)
	return getDSData(meterType, commandType, getMeterID, decryptPayload)
}

func getDSData(meterTypeInt int, commandInt int, getMeterID string, decryptPayload []byte) (DSData string) {
	switch meterTypeInt {
	case ElectricityMeter:
		return getDSCommandData(commandInt, getMeterID, decryptPayload)
	case WaterMeter:
	case IoT:
		return getDSCommandData(commandInt, getMeterID, decryptPayload)
	case GasMeter:
	case HeatMeter:
	case PV:
		return getDSCommandData(commandInt, getMeterID, decryptPayload)
	}
	return
}

func getDSCommandData(command int, getMeterID string, decryptPayload []byte) (DSCommandData string) {
	switch command {
	case ReqRegister:
		return RespondRegister(getMeterID, decryptPayload) 
	
	}
	return
}

func RespondRegister(getMeterID string, decryptPayload []byte) (DSCommandData string) {
	commandParam := utils.HexByteToHexStr(DSCommandSet.ExchangeKey["ChangeRegistrationData"])
	masterKey := getMasterKey(getMeterID)
	dataKey := getDataKey(getMeterID)
	registrationData := masterKey + dataKey
	dataLength := fmt.Sprintf("%02x", len(registrationData)/2)
	DSCommandData = getPreData(decryptPayload, getMeterID, commandParam, dataLength, registrationData)
	fmt.Println("printout: ", DSCommandData)
	return 
}




func getPreData(decryptPayload []byte, getMeterID string, commandParam ...string) (getPreData string) {
	prefix, id, _ := getPrefix(getMeterID, decryptPayload)
	predata := prefix + id + GetParam(commandParam...)
	fmt.Println(GetParam(commandParam...))
	fmt.Println("predata: ", predata)
	return predata
}

func getMasterKey(meterID string) (masterKey string) {
	masterKey = utils.SecurityKeytoHex(utils.GetSerMasterKey(meterID))
	return
}

func getDataKey(meterID string) (dataKey string) {
	dataKey = utils.SecurityKeytoHex(utils.GetSerDataKey(meterID))
	return
}
