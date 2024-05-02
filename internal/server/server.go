package server

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/JimYcod3x/meter_server/internal/meter"
	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
)

const (
	tSub1 = "J23P000078C2S"
	tSub2 = "J230008542C2S"
	tSub3 = "J200002335C2S"
	TPub  = "J23P000078S2C"
)

type Server struct {
	*mqtt.Server
}

func Run(options *mqtt.Options) *mqtt.Server {

	sev := mqtt.New(options)
	sev.AddHook(new(Hook), &HookOptions{
		Server: sev,
	})

	// client := sev.NewClient(nil, mqtt.LocalListener, mqtt.InlineClientId, true)
	// sev.Clients.Add(client)
	// TODO: only subscribe the meterid in db

	sev.Subscribe(tSub1, 1, subFn)
	sev.Subscribe(tSub2, 1, subFn)
	sev.Subscribe("/#", 1, subFn)
	go func() {
		time.Sleep(time.Second * 10)
		fmt.Println("publish the message after 10 second")
		sev.Publish(TPub, pubFn2(), true, 0)
	}()

	// sev.Publish(tPub, meter.KeyExchange(), false, 2)

	return sev
}

func subFn(cl *mqtt.Client, sub packets.Subscription, pk packets.Packet) {
	fmt.Println(sub.Filter)
	hexPayload := hex.EncodeToString([]byte(pk.Payload))
	fmt.Println("hex payload", hexPayload)

	// meter.GetMeterData(cl, pk)
}

func pubFn() []byte {
	meterType := meter.IoT
	meterDSCmd := meter.GetDataFromMeter
	// meterID := "J200002335"
	meterID := "J23P000078"
	dsCmdParam := meter.DSCommandSet.GetDataFromMeter["GetMeterID&VerInfo"]
	return meter.TestMeterFn(meterType, meterDSCmd, meterID, dsCmdParam)
}

// func PublishData()

func pubFn2() []byte {
	meterType := meter.IoT
	meterDSCmd := meter.GetDataFromMeter
	// meterID := "J200002335"
	meterID := "J23P000078"
	dsCmdParam := "Get30minRecordLog"
	startTime := "2024-05-02T15:15:00"
	endTime := "2024-05-02T15:15:01"
	return meter.CallGetDataFromMeter(meterType, meterDSCmd, meterID, dsCmdParam, startTime, endTime)
}
