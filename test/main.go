package main

import (
	"encoding/hex"
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"

	"github.com/JimYcod3x/meter_server/internal/meter"
	"github.com/JimYcod3x/meter_server/internal/utils"
)

// 604a235000007801204a3233503030303037383030303030303030303030304a323350303030303738
// 604a235000007801204a3233503030303037383030303030303030303030304a323350303030303738
// [96 74 35 80 0 0 120 1 32 48 48 48 48 48 48 74 50 51 80 48 48 48 48 55 56 48 48 48 48 48 48 74 50 51 80 48 48 48 48 55 56 0 0 0 0 0 0 0]
// [96 74 35 80 0 0 120 1 32 74 50 51 80 48 48 48 48 55 56 48 48 48 48 48 48 48 48 48 48 48 48 74 50 51 80 48 48 48 48 55 56 0 0 0 0 0 0 0]
// [96 74 35 80 0 0 120 1 32 74 50 51 80 48 48 48 48 55 56 48 48 48 48 48 48 48 48 48 48 48 48 74 50 51 80 48 48 48 48 55 56]
// [96 35 80 0 0 120 1 1 32 35 53 0 48 48 48 55 56 48 49 48 48 48 48 48 48 48 48 48 48 48 48 35 53 0 48 48 48 55 56 48 49 0 0 0 0 0 0 0]
// var payload = "604a235000007801203030303030304a3233503030303037383030303030304a32335030303030373800000000000000"
// var payload = "029e23ef1ab96549e116e9ab34b1dec5a02f189c2bc9dec1f99bfb0bfc2c1d20b71577945228e52362057441a39b2b56"
var payload = "a726439194f56bf03759f874ef4123ea"

var encrypted = "347a853d458676536f0ed66a10b72afdd10206cccbb8934c8d96109e15b293cbe8f3b905aa6ba75338544ad376d792da"
// var testmasterKey = "J23P000078000000"
// var testmasterKey = "000000J23P000078"
var testmasterKey = "69aF7&3KY0_kk89@"
func main() {
	
	// encrypted, _ := utils.Encrypt("604a2350000078003030303030304a3233503030303037380000000000000000", "000000J23P000078")
	// fmt.Println(encrypted)
	// fmt.Println(len("604a2350000078003030303030304a3233503030303037380000000000000000"))
	// 347a853d458676536f0ed66a10b72afdd10206cccbb8934c8d96109e15b293cbe8f3b905aa6ba75338544ad376d792da2e9d5eb15a6d694bdcde14d5723b40c7
	// h2b, _ := utils.HexToBinary("b9")
	// fmt.Println("h2b: ", h2b)
	// h2B, _ := hex.DecodeString("4a")
	// fmt.Println("h2B: ", string(h2B))
	// fmt.Println("h2B: ", int(0x4a))
	// fmt.Println(meter.ReqRegister)
	// fmt.Println(meter.MeterRTFailACK)
	// fmt.Printf("get the binary: %08b\n", meter.MeterRTFailACK)
	// fmt.Printf("get the binary: %012bx\n", 0b0101 | 0x01)
	// fmt.Printf("get the binary: %d\n", 0x01)
	// b2H, _ := utils.BinaryToHex("01100000")
	// fmt.Println(b2H)
	// 6a004a23500000780000000000000000
	// 604a235000007801204a3233503030303037383030303030303030303030304a323350303030303738
	// 604a235000007801204a3233503030303037383030303030303030303030304a323350303030303738
	// 604a235000007801203030303030304a3233503030303037383030303030304a32335030303030373800000000000000
	// 604a235000007801203030303030304a3233503030303037383030303030304a32335030303030373800000000000000
	// 204a200002335F01203030303030304a323030303032333335
	// 									 3030303030304a32303030303233333500000000000000

	// fmt.Println(string([]byte("017888c9ca6b6816883a128f604751d071be1cb96e9b56d77fd4b00092ff6fc899722792ab8dc2b75ae9e2d32a405ee9")))
	// byteArr, _ := utils.HexStrByteArray("3bd8275ffcc0609deef1286e801fc6c45ca0f705e1e85901b2f5f7582dbed900")
	// decrypt, _ := utils.Decrypt(byteArr, "000000J23P000078")
	// // decrypt, _ := utils.Decrypt(byteArr, "69aF7&3KY0_kk89@")
	// fmt.Println("decrypt: ", decrypt)
	// fmt.Println(utils.HexToBinary(decrypt))

	// // 204a200002335F003030303030304a3230303030323333350000000000000000
	// // 604a2350000078003030303030304a3233503030303037380000000000000000
	// id := utils.GetMeterID("J23P000078C2S")
	// bId := []byte(id)
	// hId := hex.EncodeToString(bId)
	// tHstr := "6a004a23500000780000000000000000"
	// // tHstr := "6a004a23500000780100000000000000"
	// getHByte, _ := hex.DecodeString(tHstr)
	// t, _ := strconv.ParseInt(string(getHByte[1]), 16, 64)
	// fmt.Println("t: ", t)
	// fmt.Println("t: ", string(rune(74)))
	// getHByte[3] = 0x50
	// outHex := hex.EncodeToString(getHByte)
	// fmt.Println("outHex: ", outHex)
	// fmt.Println(strings.Contains(tHstr, hId))
	// fmt.Println(hId)
	// fmt.Println([]byte(id))
	// fmt.Println("====", []byte("J23P000078C2S"))
	// // GetIdDecryptedPayload := utils.GetIdDecryptedPayload("2a004a200002335f")
	// GetIdDecryptedPayload := utils.GetIdDecryptedPayload(tHstr)
	// fmt.Println("GetIdDecryptedPayload: ", GetIdDecryptedPayload)
	test_masterkey_decrypt(encrypted)
	fmt.Println("=====================================")
	encrypted := test_masterkey_encrypt()
	encrypted1 := hex.EncodeToString(encrypted)
	fmt.Println(encrypted1)
	decrypt := test_decryptPayload()

	fmt.Println(hex.EncodeToString(decrypt))	
	fmt.Printf("%b\n", decrypt)	
	meterType := fmt.Sprintf(" MeteType %03b\n", decrypt[0] >> 5)	
	meterTypeInt := int(decrypt[0] >> 5)
	fmt.Println("print: ",meterType)	
	switch meterTypeInt{
	case meter.IoT: 
		fmt.Println("this is IoT", meterTypeInt)
	}


	command := fmt.Sprintf(" Command %b\n", decrypt[0] & 0x1f)	
	commandInt := int(decrypt[0] & 0x1f)

	switch commandInt {
	case meter.ReqRegister:
		fmt.Println("this is req register", commandInt)
	}
	fmt.Println("print: ",command)	

	meterID := fmt.Sprintf("%x", decrypt[2:8])
	fmt.Printf("print: %v\n",meterID[:])
	id, haveF := utils.FindIdHaveF(decrypt)
	fmt.Println("print idd: ",id, haveF)
	fmt.Println("print endofF: ",)
	firestLetter, _ := hex.DecodeString(meterID[:2])
	println(string(firestLetter))
	thirdLetter, _ := hex.DecodeString(meterID[4:6])
	println(string(thirdLetter))

	getID := utils.GetIdDecryptedPayload(decrypt)
	fmt.Println("print getID: ",getID)

	fmt.Printf("firestLetter %s: \n", firestLetter)
	fmt.Printf("thirdLetter %s: \n", thirdLetter)
	meterTypePayload := utils.GetMeterTypeFromPayload(decrypt)

	fmt.Println("get meter tyoe ", meterTypePayload)
	meterTypeBinaryStr := utils.IntToBinStr(meterTypePayload, meter.MeterTypeBit)
	fmt.Println("binary", meterTypeBinaryStr)
	meterCommandBinaryStr := utils.IntToBinStr(meter.ExchangeKey, meter.MeterCommandBit)
	fmt.Println("command binary", meterCommandBinaryStr)
	fmt.Println(meter.ElectricityMeter)
	masterKey := utils.SecurityKeytoHex(utils.GetSerMasterKey(getID))
	dataKey := utils.SecurityKeytoHex(utils.GetSerDataKey(getID))
	securityKey := masterKey + dataKey
	commandParam := fmt.Sprintf("%02x", 1)
	fmt.Println("commandParam: ", commandParam)
	prefix, _ := utils.BinaryToHex(meterTypeBinaryStr + meterCommandBinaryStr)
	fmt.Println("prefix: ", prefix)
	lengthOfKey := fmt.Sprintf("%02x", len(securityKey) /2)
	fmt.Println("key of len: ", lengthOfKey)
	
	predata := prefix + id + commandParam + lengthOfKey + securityKey
	// fmt.Println(len(predata) / 2)
	// fmt.Println(len("604a235000007801203030303030304a3233503030303037383030303030304a32335030303030373800000000000000") / 2)
	// fmt.Println(len(([]byte(("604a235000007801203030303030304a3233503030303037383030303030304a32335030303030373800000000000000")))))
	fmt.Println("predata: ", predata)
	testPayload, _ := utils.EncryptPadding(predata, testmasterKey)
	// hexPayload := hex.EncodeToString(testPayload)
	fmt.Println("hexPayload: ", testPayload)
	// fmt.Println(packet.OTAUpCommModule)
	test_readPayload()
	fmt.Println("=====================================")
	test_PreENcryptData(decrypt)
	fmt.Println(meter.DSCommandSet.ExchangeKey["ChangeNewKey"])
	fmt.Println(meter.DSCommandSet.OtherCtrl["ResetMeter"])
	optionsTable, ok := meter.DSCommandSet.OtherCtrl["OptionTable"].(meter.OptionsCtrl)
	if !ok {
		fmt.Println("error")
		return
	}
	fmt.Println(optionsTable["DailyMeterDiagnostic"])
	fmt.Println(hex.EncodeToString([]byte("1")))
	fmt.Println(byte(0x03))

	size := "100MB"
	re, _ := regexp.Compile("[0-9]+")
	unit := re.ReplaceAll([]byte(size), []byte(""))
	num, _ := strconv.ParseInt(strings.Replace(size, string(unit), "", 1), 10, 64)
	fmt.Println("num: ", num, unit)

	fmt.Println(utils.IntToHexStr(meter.ExchangeKey))
	fmt.Println("++++++++++++++++++++++++++")
	// meter.GetDSCommandData(meter.ExchangeKey, getID, decrypt)
	fmt.Println(utils.HexByteToHexStr(meter.DSCommandSet.ExchangeKey["ChangeRegistrationData"]))

}


func test_masterkey_encrypt() []byte{
	encryptPayload, _ := utils.EncryptPadding(payload, testmasterKey)
	fmt.Println("EncryptPayload: ", encryptPayload)
	fmt.Println("EncryptPayload: ", hex.EncodeToString(encryptPayload))
	return encryptPayload
}

func test_masterkey_decrypt(encryptedPayload string) []byte{
	byteArrPayload, _ := hex.DecodeString(encryptedPayload)
	decryptPayload, _ := utils.DecryptByte(byteArrPayload, testmasterKey)
	fmt.Println("decryptPayload: ", hex.EncodeToString(decryptPayload))
	return decryptPayload
}

func test_decryptPayload() []byte{
	decryptByte, _ := hex.DecodeString(payload)
	testdecrypt, _:= utils.DecryptByte(decryptByte, testmasterKey)
	fmt.Println("test_decryptPayload", hex.EncodeToString(testdecrypt))
	return testdecrypt
}

func test_readPayload() {
	fmt.Println("begin test_readPayload")
	r, w := net.Pipe()
	fmt.Println("then test_readPayload")
	buffer := []byte{148, 87, 220, 184, 232, 54, 84, 223, 228, 102, 196, 98 ,48, 240, 85, 161}
	
	go func() {
	_, err := w.Write(buffer)
	fmt.Println("then test_readPayload")
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}
}()
	fmt.Println("then test_readPayload")
	bytesRead := make([]byte, len(buffer))
	n, err := r.Read(bytesRead)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}


	fmt.Printf("received %d bytes %x\n", n, bytesRead) 

	
}

func test_PreENcryptData(decrypt []byte) string {
	getID := utils.GetIdDecryptedPayload(decrypt)
	fmt.Println("getID test in preencrypt: ", getID)
	id, _ := utils.FindIdHaveF(decrypt)
	masterKey := utils.SecurityKeytoHex(utils.GetSerMasterKey(getID))
	dataKey := utils.SecurityKeytoHex(utils.GetSerDataKey(getID))
	securityKey := masterKey + dataKey
	
	commandParam := fmt.Sprintf("%02x", 1)
	lengthOfKey := fmt.Sprintf("%02x", len(securityKey) /2)
	prefix := getPrefix(decrypt)
	predata := prefix + id + commandParam + lengthOfKey + securityKey
	fmt.Println("predata test: ", predata)
	return predata
}

func getPrefix(decryptPayload []byte) string {
	meterTypePayload := utils.GetMeterTypeFromPayload(decryptPayload)
	meterTypeBinaryStr := utils.IntToBinStr(meterTypePayload, meter.MeterTypeBit)
	meterCommandBinaryStr := utils.IntToBinStr(meter.ExchangeKey, meter.MeterCommandBit)
	 prefix, _ := utils.BinaryToHex(meterTypeBinaryStr + meterCommandBinaryStr)
	 fmt.Println("prefix test: ", prefix)
	return prefix
}

// keyexchange()
// datatransfer()
	// getPayload()
		// publish()
	