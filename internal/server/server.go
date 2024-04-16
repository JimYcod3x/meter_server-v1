package server

import (
	"github.com/JimYcod3x/meter_server/internal/meter"
	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
)

const (
	tSub1 = "J23P000078C2S"
	tSub2 = "J23P0000542C2S"
	tPub = "J200002335S2C"
)

type Server struct {
	*mqtt.Server	
}

func Run(options *mqtt.Options) *mqtt.Server {

	sev := mqtt.New(options)

	client := sev.NewClient(nil, mqtt.LocalListener, mqtt.InlineClientId, true)
	sev.Clients.Add(client)
	_ = sev.AddHook(new(Hook), nil)



	sev.Subscribe(tSub1, 1, subFn)
	sev.Subscribe(tSub2, 1, subFn)

	return sev
}

func subFn(cl *mqtt.Client, sub packets.Subscription, pk packets.Packet) {
	meter.GetMeterData(cl, sub, pk)
}