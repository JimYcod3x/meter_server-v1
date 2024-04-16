package server

import (
	"encoding/hex"
	"fmt"

	"github.com/JimYcod3x/meter_server/internal/meter"
	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
)

const (
	tSub1 = "J23P000078C2S"
	tSub2 = "J230008542C2S"
	tSub3 = "J200002335C2S"
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
	fmt.Println(sub.Filter)
	hexPayload := hex.EncodeToString([]byte(pk.Payload))
	fmt.Println("hex payload", hexPayload)
	meter.GetMeterData(cl, sub, pk)
}