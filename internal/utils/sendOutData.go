package utils

import (
	"fmt"
)

const (
	MeterTypeBit string = "3"
	MeterCommandBit string = "5"
)


func SentOutPrefix(meterType, meterDSCmd int, meterID string) (prefix string){
	meterTypeBinaryStr := IntToBinStr(meterType, MeterTypeBit)
	meterCommandBinaryStr := IntToBinStr(meterDSCmd, MeterCommandBit)
	prefix, _ = BinaryToHex(meterTypeBinaryStr + meterCommandBinaryStr)
	fmt.Println("prefix in testmeterfn", prefix)
	id := MeterIDtoHex(meterID)
	return prefix + id 
}
