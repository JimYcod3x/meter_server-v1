package server

import (
	"fmt"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
)

type Hook struct {
	mqtt.HookBase
}

func (h *Hook) ID() string {
	return "server-hook"
}

func (h *Hook) Provides(b byte) bool {
	return true
}

func (h *Hook) OnPacketRead(cl *mqtt.Client, pk packets.Packet) (pkx packets.Packet, err error) {
	if string(pk.Connect.Username) == "" {
		pk.Connect.UsernameFlag = false
	}
	if string(pk.Connect.Password) == "" {
		pk.Connect.PasswordFlag = false
	}
	return pk, nil
}

func (h *Hook) OnStarted() {
	fmt.Println("server started")

}

func (h *Hook) OnConnect(cl *mqtt.Client, pk packets.Packet) error {
	fmt.Println(cl.ID, "connected")
	return nil
}



func (h *Hook) OnSubscribe(cl *mqtt.Client, pk packets.Packet) packets.Packet {
	fmt.Println("subbscribed")
	return pk
}