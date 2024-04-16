// package auth

// import (
// 	"bytes"
// 	"fmt"

// 	mqtt "github.com/mochi-mqtt/server/v2"
// 	"github.com/mochi-mqtt/server/v2/packets"
// )

// type AllowHook struct {
// 	mqtt.HookBase
// }


// func (h *AllowHook) ID() string {
// 	return "allow-all-auth"
// }

// func (h *AllowHook) Provides(b byte) bool {
// return true
// }

// func (h *AllowHook) OnPacketRead(cl *mqtt.Client, pk packets.Packet) (packets.Packet, error) {

// 	fmt.Println("packet conenct", pk)
// 	fmt.Println("packet read")
// 	return pk, nil
// }

// func (h *AllowHook) OnConnectAuthenticate(cl *mqtt.Client, pk packets.Packet) bool {
// 	fmt.Println("auth")
// 	return true
// }

// func (h *AllowHook) OnACLCheck(cl *mqtt.Client, topic string, write bool) bool {
// 	fmt.Println("acl")
// 	return true
// }
