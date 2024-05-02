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
	if err != nil {
		log.Fatal("can not decrypt")
	}

	masterKey := getMasterKey(getID)
	dataKey := getDataKey(getID)
	securityKey := masterKey + dataKey
	commandParam := utils.HexByteToHexStr(DSCommandSet.ExchangeKey["ChangeRegistrationData"])
	lengthOfKey := fmt.Sprintf("%02x", len(securityKey)/2)
	prefix, id := getPrefix(decryptPayload)
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

func getPrefix(decryptPayload []byte) (prefix string, id string) {
	meterTypePayload := utils.GetMeterTypeFromPayload(decryptPayload)
	meterTypeBinaryStr := utils.IntToBinStr(meterTypePayload, utils.MeterTypeBit)
	meterCommandBinaryStr := utils.IntToBinStr(ExchangeKey, utils.MeterCommandBit)
	prefix, _ = utils.BinaryToHex(meterTypeBinaryStr + meterCommandBinaryStr)
	id, _ = utils.FindIdHaveF(decryptPayload)
	fmt.Println("prefix binary", meterCommandBinaryStr)
	fmt.Printf("prefix: %s id: %s\n", prefix, id)
	return
}


// func ChangeKeyFn(pk packets.Packet, DecryptKey string) string {
// 	fmt.Println("plainPayload: ", pk.Payload)
// 	decryptPayload, _ := utils.DecryptByte(pk.Payload, DecryptKey)
// 	meterType := utils.GetUSCommandFromDecrypt(decryptPayload)
// 	commandType := int(decryptPayload[0] & 0x1f)
// 	getMeterID := utils.GetMeterID(pk.TopicName)
// 	fmt.Println("decryptPayload: ", decryptPayload)
// 	fmt.Println(meterType, commandType, getMeterID, decryptPayload)
// 	fmt.Println("Down stream: ", getDSData(meterType, commandType, getMeterID, decryptPayload))
// 	fmt.Println("debugPrint: ", meterType, commandType, getMeterID, decryptPayload)
// 	return getDSData(meterType, commandType, getMeterID, decryptPayload)
// }

// func getDSData(meterTypeInt int, commandInt int, getMeterID string, decryptPayload []byte) (DSData string) {

// 	switch meterTypeInt {
// 	case ElectricityMeter:
// 		return getDSCommandData(commandInt, getMeterID, decryptPayload)
// 	// case WaterMeter:
// 	case IoT:
// 		fmt.Println("This is IoT type")
// 		return getDSCommandData(commandInt, getMeterID, decryptPayload)
// 	// case GasMeter:
// 	// case HeatMeter:
// 	case PV:
// 		return getDSCommandData(commandInt, getMeterID, decryptPayload)
// 	}
// 	return
// }

// func getDSCommandData(command int, getMeterID string, decryptPayload []byte) (DSCommandData string) {

// 	switch command {
// 	case USC:
// 		fmt.Println("ReqRegister")
// 		return RespondRegister(getMeterID, decryptPayload)
// 	case ReqChangeKey:
// 		fmt.Println("ReqChangeKey")
// 		return ResponseNewDataKey(getMeterID, decryptPayload)
// 	case ReqSucACK:
// 		fmt.Println("ReqSucACK")
// 		return ConfirmNewDataKey(getMeterID, decryptPayload)
// 	}
// 	return
// }

func getPreData(decryptPayload []byte, commandParam ...string) (getPreData string) {
	prefix, id := getPrefix(decryptPayload)
	predata := prefix + id + GetParam(commandParam...)
	// fmt.Println(GetParam(commandParam...))
	fmt.Println("predata: ", predata)
	return predata
}

func getMasterKey(meterID string) (masterKey string) {
	masterKey, found := utils.GetSerMasterKey(meterID)
	if !found {
		log.Fatal("can not get the master key from database")
	}
	masterKey = utils.SecurityKeytoHex(masterKey)
	return
}

func getDataKey(meterID string) (dataKey string) {
	dataKey, found := utils.GetSerDataKey(meterID)
	if !found {
		log.Fatal("can not get the datakey from database")
	}
	dataKey = utils.SecurityKeytoHex(dataKey)
	return
}
