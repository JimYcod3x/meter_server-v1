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

func (h *Hook) OnStarted() {
	fmt.Println("server started")

}

func (h *Hook) OnConnect(c *mqtt.Client, pk packets.Packet) error {
	fmt.Println("client connected")
	return nil
}

func (h *Hook) OnPacketRead(cl *mqtt.Client, pk packets.Packet) (pkx packets.Packet, err error) {
	fmt.Println("packet read")
	return pk, nil
}

func (h *Hook) OnSubscribe(cl *mqtt.Client, pk packets.Packet) packets.Packet {
	fmt.Println("subbscribed")
	return pk
}