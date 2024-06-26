package utils

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/mochi-mqtt/server/v2/packets"
)

func HexToBinary(hexS string) (binStr []string, err error) {
	bHex, err := hex.DecodeString(hexS)
	if err != nil {
		return nil, err
	}
	for _, b := range bHex {
		binStr = append(binStr, fmt.Sprintf("%08b", b))
	}
	return
}

func BinaryToHex(binStr string) (hexS string, err error) {
	binaryToInteger, err := strconv.ParseInt(binStr, 2, 64)
	if err != nil {
		return "", err
	}
	hexS = fmt.Sprintf("%02x", binaryToInteger)
	return
}

func SecurityKeytoHex(key string) (hexKey string) {
	byteKey := []byte(key)
	hexKey = hex.EncodeToString(byteKey)
	return
}

func IntToBinStr(i int, numberOfBit string) string {
	return fmt.Sprintf("%0"+numberOfBit+"b", i)
}

func IntToHexStr(i int) string {
	return fmt.Sprintf("%02x", i)
}

func HexByteToHexStr(b byte) string {
	return fmt.Sprintf("%02x", b)
}

func HexStrByteArray(HexS string) (byteArr []byte, err error) {
	byteArr, err = hex.DecodeString(HexS)
	if err != nil {
		return nil, err
	}
	return
}

func GetMeterID(topicStr string) (meterID string) {
	return topicStr[:len(topicStr)-3]
}

func GetIdDecryptedPayload(dcptPayload []byte) (meterID string) {
	idHexA, endWithF := FindIdHaveF(dcptPayload)
	lenOfId := len(idHexA)
	finalId := ""

	if endWithF {
		finalId := idHexA[:lenOfId-1]
		// fmt.Println("before", finalId)

		idFirstLetter, _ := hex.DecodeString(idHexA[:2])

		finalId = strings.Replace(finalId, finalId[:2], string(idFirstLetter), 1)

		// fmt.Println("FinalId" ,finalId)

		return finalId
	}

	firstLetter, _ := hex.DecodeString(idHexA[:2])
	fourletter, _ := hex.DecodeString(idHexA[4:6])

	RuneArr := strings.Replace(idHexA, idHexA[:2], string(firstLetter), 1)
	finalId = strings.Replace(RuneArr, idHexA[4:6], string(fourletter), 1)

	// fmt.Println("finalId: ", finalId)

	return finalId
}

func FindIdHaveF(payload []byte) (id string, haveF bool) {
	char := "f"

	meterId := fmt.Sprintf("%x", payload[2:8])

	endOfId := meterId[len(meterId) - 1:]

	if endOfId == char {
		return meterId, true
	}
	return meterId, false
}


func GetMeterIDFromTopic(pk packets.Packet) string{
	
	topic := pk.TopicName
	fmt.Println(topic)
	meterID := topic[:len(topic)-3]
	return meterID
}

func GetMeterTypeFromPayload(payload []byte) (meterType int) {
	meterBStr  := fmt.Sprintf("%0d", payload[0] >> 5)	
	meterType, _ = strconv.Atoi(meterBStr)
	return meterType
}


func ValidateMeter(meterID string, payload []byte, decryptKey string) (valid bool) {
	// Use server dataKey Decrypt payload
	// data, err := utils.Decrypt(pk.Payload, dataKey)
	fmt.Println("get payload: ", payload)
	data, err := DecryptByte(payload, decryptKey)
	if err != nil {
		fmt.Println("Data DecrptError:", err)
	}
	fmt.Println("decrypt data: ", data)

	idGet := GetIdDecryptedPayload(data)
	fmt.Println("id from payload: ", idGet)
	fmt.Println("test: ", meterID)
	return meterID == idGet
}

func DSTopic(pk packets.Packet) (topic string) {
	topic = strings.Replace(pk.TopicName, "C2S", "S2C",1)
	return 
}

func GetUSCommandFromDecrypt(decrypt []byte) (command int) {
	commandByte := decrypt[0] & 0x1f
	decryptBin := fmt.Sprintf("%08b", decrypt[0])
	commandBStr := fmt.Sprintf("%04b", commandByte >> 1)
	commandInt, _ := strconv.ParseInt(commandBStr, 2, 64)
	command = int(commandInt)
	fmt.Println("new command", command, decryptBin, commandBStr)
	fmt.Printf("command get %08b\n", (decrypt[0] & 0x1f))
	// fmt.Printf("command get %0b\n", (decrypt[0] & 0x1f) >> 1)
	return
}


func MeterIDtoHex(id string) string {
	hexStr := ""

	for _, char := range id {
		switch char{
		case 'J':
			hexStr += "4a"
		case 'P':
			hexStr += "50"
		default:
			hexStr += string(char)
		}
	}
	if len(hexStr) < 12 {
		hexStr += "f"
	}
	fmt.Println(hexStr)
	return hexStr
}

func ParseTimeToHexStr(timeStr any) string{
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

// J23P000078


// 4a2300008542
