package meter

type meterUS interface {
	HandleMeterReading()
	HandleHHdata()
	HandleEventLog()
	HandleAPowerLineValue()
	HandleReqACK()
	HandleRS485Res()
	HandleThreeMonIXportE()
	HandleMultiTelthHourData()
	HandleMultiTelEventLog()
}

