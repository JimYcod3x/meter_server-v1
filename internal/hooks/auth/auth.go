package auth

import (
	"bytes"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
)

type AllowHook struct {
	mqtt.HookBase
}


func (h *AllowHook) ID() string {
	return "allow-all-auth"
}

func (h *AllowHook) Provides(b byte) bool {
	return bytes.Contains([]byte{
		mqtt.OnConnectAuthenticate,
		mqtt.OnACLCheck,
	}, []byte{b})
}

func (h *AllowHook) OnConnectAuthenticate(cl *mqtt.Client, pk packets.Packet) bool {
	return true
}

func (h *AllowHook) OnACLCheck(cl *mqtt.Client, topic string, write bool) bool {
	return true
}
