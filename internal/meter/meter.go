package meter

import (
	"fmt"
)

const (
	ElectricityMeter MeterType = iota + 0b001 // 001
	WaterMeter                            // 010
	IoT                                   // 011
	GasMeter                              // 100
	HeatMeter                             // 101
	PV                                    // 110
)



const (
	DefaultKey = "69aF7&3KY0_kk89@"
	// defaultKey = "000000J23P000078"
	// "69aF7&3KY0_kk89@"
	
)

type MeterType = int

type Meter struct {
	ID   string
	Type MeterType
}






// func GetMeterData(cl *mqtt.Client, pk packets.Packet) string {
// 	decryptPayload, err := utils.DecryptByte(pk.Payload, DefaultKey)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(pk.Payload)
// 	fmt.Println(decryptPayload)

// 	binDecry:= fmt.Sprintf("%b", decryptPayload)
// 	defer func() {
// 		err := recover()
// 		if err != nil {
// 			fmt.Println("Error:", err)
// 		}
// 	}()
// 	fmt.Println("binary get:", binDecry[0])
// 	sliceArr, _ := strconv.ParseInt(binDecry[0], 2, 8)
	

// 	meterType := (sliceArr >> 5) & 0x7
// 	fmt.Println("meter type get:", meterType)
// 	GetMeterType(meterType)

	
// 	// if (sub.Filter[-3:0])
// 	if pk.TopicName[len(pk.TopicName) -3 :] == "C2S" {
// 		fmt.Println("C2S")
// 	} else if pk.TopicName[len(pk.TopicName) -3 :] == "S2C" {
// 		fmt.Println("S2C")
// 	}
// 	fmt.Printf("get the binary: %08b\n", sliceArr)
// 	fmt.Printf("get the binary: %08x\n", 0b1111)
// 	fmt.Printf("get the binary: %04b\n", (sliceArr >> 1) & 0xf)
// 	// cmdType := sliceArr[3:]
// 	// fmt.Println("cmd type get:", cmdType)

// 	buff := make([]byte, len(decryptPayload))
// 	buffget := copy(buff, []byte(decryptPayload))
// 	fmt.Println("buff", buffget)
// 	return string(decryptPayload)
// }

func GetMeterType(meterType int) {
	switch meterType {
	case ElectricityMeter:
		fmt.Println("This is Electricity Meter")
	case WaterMeter:
		fmt.Println("This is Water Meter")
	case GasMeter:
		fmt.Println("This is Gas Meter")
	case HeatMeter:
		fmt.Println("This is Heat Meter")
	case PV:
		fmt.Println("This is PV")
	case IoT:
		fmt.Println("This is IoT")
	default:
		fmt.Println("This is not a meter")
	}
}

// func KeyExchange() []byte{
// 	// getdecryptPayload

// 	ByteArr, _ := utils.HexStrByteArray("3bd8275ffcc0609deef1286e801fc6c45ca0f705e1e85901b2f5f7582dbed900")
// 	fmt.Println(ExchangeKeyParam{ChangeNewKey: "000000J23P000078"})
// 	return ByteArr
// }

// func MeterRegistration(meterID string, pk packets.Packet) {
// 	prefix :=  + meterID
// 	param := ""
// 	data := prefix + param
// 	s.Publish(pk.TopicName, utils.PublishData(data), false, 2)
// }

func(m *Meter) HandleMeterReading() {

}


