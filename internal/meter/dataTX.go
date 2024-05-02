package meter

import (
	"encoding/hex"
	"fmt"

	"github.com/JimYcod3x/meter_server/internal/utils"
)

// func DataTXFn(pk packets.Packet, DecryptKey string, usCmdParam string) string {
// 	fmt.Println("plainPayload: ", pk.Payload)
// 	decryptPayload, _ := utils.DecryptByte(pk.Payload, DecryptKey)
// 	meterType := utils.GetMeterTypeFromPayload(decryptPayload)
// 	commandType := utils.GetUSCommandFromDecrypt(decryptPayload)
// 	getMeterID := utils.GetMeterID(pk.TopicName)
// 	fmt.Println("decryptPayload: ", decryptPayload)
// 	fmt.Println(meterType, commandType, getMeterID, decryptPayload)
// 	fmt.Println("Down stream: ", getDSData(meterType, commandType, getMeterID, decryptPayload))
// 	fmt.Println("debugPrint: ", meterType, commandType, getMeterID, decryptPayload)
// 	return getKeyXDSKeyXData(meterType, getMeterID, decryptPayload, usCmdParam)
// }

// func getDSPrefix() (prefix string, id string) {

// }

func TestMeterFn(meterType , meterDSCmd int, meterID string, dsCmdParam byte) []byte {
	meterTypeBinaryStr := utils.IntToBinStr(meterType, utils.MeterTypeBit)
	meterCommandBinaryStr := utils.IntToBinStr(meterDSCmd, utils.MeterCommandBit)
	prefix, _ := utils.BinaryToHex(meterTypeBinaryStr + meterCommandBinaryStr)
	fmt.Println("prefix in testmeterfn", prefix)
	id := utils.MeterIDtoHex(meterID)
	commandParam := hex.EncodeToString([]byte{dsCmdParam})
	preEncry := prefix + id + GetParam(commandParam)
	dataKey, _ := utils.GetSerDataKey(meterID)
	preEncyhexStr, _ := hex.DecodeString(preEncry)
	fmt.Println("preEncrypt payload in testmeterfn", preEncyhexStr)
	sentouPayload, _ := utils.EncryptPadding(preEncry, dataKey)
	return sentouPayload
}

