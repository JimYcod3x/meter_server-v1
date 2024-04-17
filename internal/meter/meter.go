package meter

import (
	"fmt"
	"log"

	"github.com/JimYcod3x/meter_server/internal/utils"
	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
)

const (
	ElectricityMeter MeterType = iota + 1 // 001
	WaterMeter                            // 010
	IoT                                   //011
	GasMeter                              // 100
	HeatMeter                             // 101
	PV                                    // 110
)

const (
	defaultKey = "69aF7&3KY0_kk89@"
)

type MeterType int

type Meter struct {
	ID   string
	Type MeterType
}

func GetMeterData(cl *mqtt.Client, sub packets.Subscription, pk packets.Packet) string {
	decryptPayload, err := utils.Decrypt(pk.Payload, defaultKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pk.Payload)
	fmt.Println(decryptPayload)

	binDecry, _ := utils.HexToBinary(decryptPayload)
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Error:", err)
		}
	}()
	fmt.Println("binary get:", binDecry[0])
	sliceArr := binDecry[0]
	meterType := sliceArr[0:3]
	fmt.Println("meter type get:", meterType)
	GetMeterType(meterType)

	cmdType := sliceArr[3:]
	fmt.Println("cmd type get:", cmdType)

	buff := make([]byte, len(decryptPayload))
	buffget := copy(buff, []byte(decryptPayload))
	fmt.Println("buff", buffget)
	return decryptPayload
}

func GetMeterType(meterType string) {
	switch meterType {
	case fmt.Sprintf("%03b", ElectricityMeter):
		fmt.Println("This is Electricity Meter")
	case fmt.Sprintf("%03b", WaterMeter):
		fmt.Println("This is Water Meter")
	case fmt.Sprintf("%03b", GasMeter):
		fmt.Println("This is Gas Meter")
	case fmt.Sprintf("%03b", HeatMeter):
		fmt.Println("This is Heat Meter")
	case fmt.Sprintf("%03b", PV):
		fmt.Println("This is PV")
	case fmt.Sprintf("%03b", IoT):
		fmt.Println("This is IoT")
	default:
		fmt.Println("This is not a meter")
	}
}

