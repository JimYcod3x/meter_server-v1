package meter

import (
	"fmt"

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

func PreEncryptData(pk packets.Packet) string {
	getID := utils.GetMeterID(pk.TopicName)
	fmt.Println("getID test: ", getID)
	DecryptKey := "69aF7&3KY0_kk89@"
	decryptPayload, _ := utils.DecryptByte(pk.Payload, DecryptKey)
	id, _ := utils.FindIdHaveF(decryptPayload)
	masterKey := utils.SecurityKeytoHex(DecryptKey)
	// masterKey := utils.SecurityKeytoHex(utils.GetSerMasterKey(getID))
	dataKey := utils.SecurityKeytoHex(DecryptKey)
	// dataKey := utils.SecurityKeytoHex(utils.GetSerDataKey(getID))
	securityKey := masterKey + dataKey
	
	commandParam := getCommandParam()
	lengthOfKey := fmt.Sprintf("%02x", len(securityKey) /2)
	prefix := getPrefix(decryptPayload)
	predata := prefix + id + commandParam + lengthOfKey + securityKey
	
	return predata
}

func getPrefix(decryptPayload []byte) string {
	meterTypePayload := utils.GetMeterTypeFromPayload(decryptPayload)
	meterTypeBinaryStr := utils.IntToString(meterTypePayload, MeterTypeBit)
	meterCommandBinaryStr := utils.IntToString(ExchangeKey, MeterCommandBit)
	 prefix, _ := utils.BinaryToHex(meterTypeBinaryStr + meterCommandBinaryStr)
	 fmt.Println("prefix test: ", prefix)
	return prefix
}

func getCommandParam() string {
	return fmt.Sprintf("%02x", 1)
}