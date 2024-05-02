package meter

const (
	MeterRead USCommand = iota + 0b0001
	hHourData
	EvnetLog
	APowerLineValue
	ReqACK
	RS485Res
	ThreeMonIXportE
	_
	_
	MultiTelthHourData
	MultiTelEventLog
)

const (
	ExchangeKey DSCommand = 0b00000 + iota
	GetDataFromMeter
	SwitchCtrl
	_
	_
	_
	OtherCtrl
	BRouteMeterInfo
	RS485Ctrl
)

const (
	OTAUpCommModule         = 0b11111 - iota
	OTAUpWiFi               // 0b11110
	OTATSLHTTPSCertKeyDload // 0b11101
	OTAUpMeterFirm          // 0b11100
	OTAUpBootloader         // 0b11011
)

type Command = int
type USCommand = int
type DSCommand = int

type FWUpOKFail map[string]uint16

type OptionsCtrl map[string]uint16

type ParamSet map[string]byte

type USCmd struct {
	ReqACK             map[string]interface{}
	RS485Res           map[string]byte
	ThreeMonIXportE    map[string]byte
	MultiTelthHourData map[string]byte
	MultiTelEventLog   map[string]byte
}

var USCommandSet = USCmd{
	ReqACK: map[string]interface{}{
		"ReqRegistration": 0x0,
		"ReqChangeKey":    0x1,
		"ReqSucACK":       0x2,
		"FWUpReq":         0x3,
		"FWUpOK&Fail": FWUpOKFail{
			"FWUpOK":   0x400,
			"FWUpFail": 0x401,
		},
		"MeterID&Ver":     0x5,
		"DiagInfo":        0x7,
		"BRouteID&Passwd": 0x8,
		"OptionACK":       0x9,
		"BusyACK":         0xB,
		"SIMInfo":         0xC,
		"NotUsed":         0xE,
		"ParamACK":        0xF,
		"ResetFailACK":    0xF0,
	},
	RS485Res: map[string]byte{
		"RS485Res":      0x0,
		"RS485ResErrTO": 0x1,
	},
	ThreeMonIXportE: map[string]byte{
		"ThreeMonPVImportE": 0x0,
		"ThreeMonRBExportE": 0x1,
	},
}

type DSCmd struct {
	ExchangeKey             map[string]byte
	GetDataFromMeter        map[string]byte
	SwitchCtrl              map[string][]byte
	OtherCtrl               map[string]interface{}
	BRouteMeterInfo         map[string]byte
	RS485Ctrl               map[string]byte
	OTAUpCommModule         map[string]byte
	OTAUpWiFi               map[string]byte
	OTATSLHTTPSCertKeyDload map[string]byte
	OTAUpMeterFirm          map[string]byte
	OTAUpBootloader         map[string]byte
}

var DSCommandSet = DSCmd{
	ExchangeKey: map[string]byte{
		"ChangeNewKey":           0x0,
		"ChangeRegistrationData": 0x1,
		"ConfirmNewKey":          0x2,
	},
	GetDataFromMeter: map[string]byte{
		"GetInstantReading":            0x1,
		"Get30minRecordLog":            0x2,
		"GetEventLogRecords":           0x3,
		"GetAvgInstantPowerLineValue":  0x4,
		"GetMeterID&VerInfo":           0x5,
		"Reserved":                     0x6,
		"GetDiagnosticInfo":            0x7,
		"GetCommModuleProgramFlashCRC": 0x8,
		"GetMeterProgramFashCRC":       0x9,
	},
	SwitchCtrl: map[string][]byte{
		"TurnOnSwitch":  []byte("ON"),
		"TurnOffSwitch": []byte("OFF"),
	},
	OtherCtrl: map[string]interface{}{
		"ResetLastSentoCurrentTime": 0x0,
		"ResetMeter":                0x1,
		"ResetCommModule":           0x2,
		"OptionTable": OptionsCtrl{
			"DailyMeterDiagnostic":       0x3ff,
			"SunshineTempHumidity":       0x301,
			"EngineeringEventLog":        0x302,
			"LoadLimit":                  0x303,
			"LoadCurrent":                0x304,
			"AutoConnectTime":            0x305,
			"AutoConnectCount":           0x306,
			"TimeToClearAutoConnectTime": 0x307,
			"InvertControlSwitchState":   0x308,
			"PCSCommMode":                0x309,
			"IntervalDataSending":        0x380,
			"IoTPlatform":                0x381,
			"DailyDiagnostic":            0x382,
			"RouteBCommLog":              0x383,
			"UnderLTESoftSIM":            0x384,
		},
		"SIMInfo":   0x4,
		"IPSetting": 0x5,
		"ParamTable": ParamSet{
			"RBDiagnostic": 0x0,
			"IPPortBuff":   0x1,
			"Mai&AuxSWCrt": 0x2,
		},
	},
	BRouteMeterInfo: map[string]byte{
		"GetBRouteID&Passwd": 0x0,
		"SetBRouteUD&Passwd": 0x2c,
	},
	RS485Ctrl: map[string]byte{
		"PCS":      0x0,
		"AlwaysOn": 0x1,
	},
	OTAUpCommModule: map[string]byte{
		"FWInfo":      0x0,
		"FWUpContent": 0x1,
	},
	OTAUpWiFi: map[string]byte{
		"FWInfo":      0x0,
		"FWUpContent": 0x1,
	},
	OTATSLHTTPSCertKeyDload: map[string]byte{
		"TLSCert&KeyInfo":    0x0,
		"TLSCert&KeyContent": 0x1,
	},
	OTAUpMeterFirm: map[string]byte{
		"FWInfo":      0x0,
		"FWUpContent": 0x1,
	},
	OTAUpBootloader: map[string]byte{
		"FWInfo":      0x0,
		"FWUpContent": 0x1,
	},
}



