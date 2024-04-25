package meter

const (
	MeterRead Command = iota + 1
	HalfHData
	EvnetLog
	APowerLineValue
)

const (
	ReqRegister USCommand = iota + 0b1010
	ReqChangeKey
	ReqSucACK
	ReqFWUpgrade
	FWUpgradeOK
	FWUpgradeFail
	GetMeterIDnV
	_
	DiagInfo
	BrouteIDnPasswd
	OptsACK
	_
	BusyACK
	SIMInfoACK
	_
	NotUsed
	ParamACK
	MeterRTFailACK
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
	OTAUpCommModule  = 0b11111 - iota
	OTAUpWiFi               // 0b11110
	OTATSLHTTPSCertKeyDload // 0b11101
	OTAUpMeterFirm          // 0b11100
	OTAUpBootloader        // 0b11011
	
	
)

type Command = int
type USCommand = int
type DSCommand = int

type OptionsCtrl  map[string]byte
// Daily Meter Diagnostic 	
// Sunshine, temperature and humidity  sensor 	
// Engineering Event Log	
// Load Limit (for 2302H)	
// Load current (for 2302H)	
// Auto-connect time (for 2302H)	
// Auto-connect count (for 2302H)	
// Time to clear the Auto-connect time (for 2302H)	
// Invert control switch state (for 2302H)	
// PCS Communication Mode	
// Interval for data sending	
// IoT Platform	
// Daily Diagnostic *1	
// Route B Communication Log	
// Under LTE SoftSIM F/W, run Soft/Hard SIM *2	
// type OptionsCtrl struct {
	
// }

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
	SwitchCtrl:              map[string][]byte{
		"TurnOnSwitch":                  []byte("ON"),
		"TurnOffSwitch":                 []byte("OFF"),
	},
	OtherCtrl:               map[string]interface{}{
		"ResetLastSentoCurrentTime": 0x0,
		"ResetMeter":                0x1,
		"ResetCommModule":           0x2,
		"OptionTable":               OptionsCtrl{
			"DailyMeterDiagnostic": 0xff,
			"SunshineTempHumidity": 0x1,
			"EngineeringEventLog": 0x2,
			"LoadLimit": 0x3,
			"LoadCurrent": 0x4,
			"AutoConnectTime": 0x5,
			"AutoConnectCount": 0x6,
			"TimeToClearAutoConnectTime": 0x7,
			"InvertControlSwitchState": 0x8,
			"PCSCommMode": 0x9,
			"IntervalDataSending": 0x80,
			"IoTPlatform": 0x81,
			"DailyDiagnostic": 0x82,
			"RouteBCommLog": 0x83,
			"UnderLTESoftSIM": 0x84,
		},

	},
	BRouteMeterInfo:         map[string]byte{
		"GetBRouteID&Passwd": 0x0,
		"SetBRouteUD&Passwd": 0x1,
	},
	RS485Ctrl:               map[string]byte{
		"PCS": 0x0,
		"AlwaysOn": 0x1,
	},
	OTAUpCommModule:         map[string]byte{},
	OTAUpWiFi:               map[string]byte{},
	OTATSLHTTPSCertKeyDload: map[string]byte{},
	OTAUpMeterFirm:          map[string]byte{},
	OTAUpBootloader:         map[string]byte{},
}



// 
// Reset Meter
// Reset Comm Module
// See option table below
// SIM card info, e.g.  |APN|User name|Password|Authentication type|
// IP address, e. g. |34.84.143.129|1883|
// See parameter table below
// Get B-route user ID and password
// B-route user ID and password can be set "0" and 0" respectively. It is for disable the B-route function.
// PCS
// "RS485 device with power always on
// (e.g. Slave meter)"
// Pulse meter type = 表计类型
// Firmware information
// Firmware content
// Firmware information
// Firmware content
// TLS/ https Certiicate and Key Information
// TLS/ https Certiicate and Key Content
// Firmware information
// Firmware content
// Firmware information
// Firmware content
