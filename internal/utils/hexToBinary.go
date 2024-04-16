package utils

import (
	"encoding/hex"
	"fmt"
)

//

func HexToBinary(hexS string) ([]string, error) {
	ui, err := hex.DecodeString(hexS)
	fmt.Println(ui)
	if err != nil {
		return nil, err
	}
	var binarySlice []string
	for _, b := range ui {
		binarySlice = append(binarySlice, fmt.Sprintf("%08b", b))
	}
	fmt.Println(binarySlice)
	return binarySlice, nil
}