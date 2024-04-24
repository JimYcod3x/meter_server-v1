package packet



type DataForm struct {
	Identifier Indentifier
	MeterID string
	DataPacket []byte
}


	type Indentifier struct {
		// MeterType *MeterType
		PacketType PacketType
	}

	type PacketType struct {
	}


	// func GetCommandType(cmdType string) {
	// 	if 
	// 	switch cmdType {

	// 	}
	// }

