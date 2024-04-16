package meter

const (
	MeterRead Command = iota + 1
	HalfHData
	EvnetLog
	APowerLineValue

)

type Command int

type DataForm struct {
	Identifier Indentifier
	MeterID string
	DataPacket []byte
}


	type Indentifier struct {
		MeterType MeterType
		PacketType PacketType
	}

	type PacketType struct {
	}

