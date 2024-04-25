package server

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/JimYcod3x/meter_server/internal/meter"
	"github.com/JimYcod3x/meter_server/internal/utils"
	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
)

const DefaultKey = "69aF7&3KY0_kk89@"

type Hook struct {
	mqtt.HookBase
	config *HookOptions
}

type HookOptions struct {
	Server *mqtt.Server
}

func (h *Hook) Init(config any) error {
	h.Log.Info("initialised")
	if _, ok := config.(*HookOptions); !ok && config != nil {
		return mqtt.ErrInvalidConfigType
	}

	h.config = config.(*HookOptions)
	if h.config.Server == nil {
		return mqtt.ErrInvalidConfigType
	}
	return nil
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
	log.Println("server started")

}

func (h *Hook) OnConnect(cl *mqtt.Client, pk packets.Packet) error {
	fmt.Println(cl.ID, "connected")

	return nil
}

func (h *Hook) OnSubscribe(cl *mqtt.Client, pk packets.Packet) packets.Packet {

	fmt.Println(cl.ID, "subbscribed", pk.Filters[0].Filter)

	return pk
}

func (h *Hook) OnPublish(cl *mqtt.Client, pk packets.Packet) (packets.Packet, error) {
	if cl.ID != "inline" {
	h.Log.Info(fmt.Sprintf("published payload: %v", hex.EncodeToString(pk.Payload)))

		
	// 
	// _  = h.config.Server.Publish("J23P000078S2C", []byte("test"), false, 0)

	}

	// if err != nil {
	// 	go func() {
	// 		defer recover()
	// 	}()
	// }
	return pk, nil
}

func (h *Hook) OnPublished(cl *mqtt.Client, pk packets.Packet) {

	// datatransfer()
	// getPayload()
	// publish()

	// Get meterID from topic

	// fmt.Println("Connect Params:", pk.Connect)
	// fmt.Println("Properties:", pk.Properties)
	// fmt.Println("Payload:", string(pk.Payload))
	// fmt.Println("Reason Codes:", pk.ReasonCodes)
	// fmt.Println("Filters:", pk.Filters)
	// fmt.Println("Topic Name:", pk.TopicName)
	// fmt.Println("Origin:", pk.Origin)
	// fmt.Println("Fixed Header:", pk.FixedHeader)
	// fmt.Println("Created:", pk.Created)
	// fmt.Println("Expiry:", pk.Expiry)
	// fmt.Println("Mods:", pk.Mods)
	// fmt.Println("pk ID:", pk.PacketID)
	// fmt.Println("Protocol Version:", pk.ProtocolVersion)
	// fmt.Println("Session Present:", pk.SessionPresent)
	// fmt.Println("Reason Code:", pk.ReasonCode)
	// fmt.Println("Reserved Bit:", pk.ReservedBit)
	// fmt.Println("Ignore:", pk.Ignore)

	// if !string.Contains(str, "meterID")
	// use master key to decrypt payload

	// if !string.Contains(str, "master")

	// meter.GetMeterData(cl, pk)
	if cl.ID != "inline" {

		meterID := utils.GetMeterIDFromTopic(pk)
//  Get dataKey from server
		dataKey := utils.GetSerDataKey(meterID)
		payload := pk.Payload
		// Validate if payload contains meter id
		isValidBySerDataKey := utils.ValidateMeter(meterID, payload, dataKey)
		fmt.Println(isValidBySerDataKey)
		if !isValidBySerDataKey {
			fmt.Println("Meter ID not found in payload")
			fmt.Println("Try the master key decrypting the payload")
			fmt.Println("=====================================")
			// Use master key to decrypt payload
			masterKey := utils.GetSerMasterKey(meterID)
			// data, err := utils.Decrypt(pk.Payload, masterKey)
			isValidBySerMasterKey := utils.ValidateMeter(meterID, payload, masterKey)
	
			if !isValidBySerMasterKey {
				fmt.Println("Meter ID not found")
				fmt.Println("User default key to decrypt payload")
				fmt.Println("=====================================")
				// Use default key to decrypt payload
				defaultKey := DefaultKey
				isValidBySerDefalutKey := utils.ValidateMeter(meterID, payload, defaultKey)
				if !isValidBySerDefalutKey {
					fmt.Println("Meter ID not found")
					fmt.Println("Meter validation failed")
					fmt.Println("=====================================")
					return
				}
				fmt.Println("Meter validation completed")
				fmt.Println("Begin to Register Meter....")
				// publishPayload, err := utils.EncryptPadding(
					publishPayload := meter.DataTXFn(pk, DefaultKey)
				// if err != nil {
				// 	fmt.Println("Error encrypting payload")
				// 	return 
				// }
				fmt.Println(publishPayload)
				// keyexchange()
				// _  = h.config.Server.Publish(utils.DSTopic(pk), publishPayload, false, 0)
				// fmt.Println("keyExchange")
				// sendoutPayload := fmt.Sprintf("%x", publishPayload)
				// fmt.Println((sendoutPayload))
				// main.Sev.Publish(tPub, meter.KeyExchange(), false, 2)
	
			}
	
		}
		fmt.Println("DownStream Topic: ", utils.DSTopic(pk))
		_  = h.config.Server.Publish(utils.DSTopic(pk), []byte("datatransfer"), false, 0)
		// fmt.Println("prefix: ", meter.ExchangeKeyFn(pk))
	
}

}

// OnConnectAuthenticate method

// OnSysInfoTick method
// func (h *Hook) OnSysInfoTick(info *system.Info) {
// 	// Implement your logic here
// 	fmt.Println("OnSysInfoTick")
// }

// // OnSessionEstablish method
// func (h *Hook) OnSessionEstablish(cl *mqtt.Client, pk packets.Packet) {
// 	// Implement your logic here
// 	fmt.Println("OnSessionEstablish")
// }

// // OnSessionEstablished method
// func (h *Hook) OnSessionEstablished(cl *mqtt.Client, pk packets.Packet) {
// 	// Implement your logic here
// 	fmt.Println("OnSessionEstablished")
// }

// // OnDisconnect method
// func (h *Hook) OnDisconnect(cl *mqtt.Client, err error, expire bool) {
// 	// Implement your logic here
// 	fmt.Println("OnDisconnect")
// }

// // OnPacketEncode method
// func (h *Hook) OnPacketEncode(cl *mqtt.Client, pk packets.Packet) packets.Packet {
// 	fmt.Println("OnPacketEncode")
// 	// Implement your logic here
// 	return pk // Placeholder return value, change as per your implementation
// }

// // OnPacketSent method
// func (h *Hook) OnPacketSent(cl *mqtt.Client, pk packets.Packet, b []byte) {
// 	// Implement your logic here
// 	fmt.Println("OnPacketSent")
// }

// // OnPacketProcessed method
// func (h *Hook) OnPacketProcessed(cl *mqtt.Client, pk packets.Packet, err error) {
// 	// Implement your logic here
// 	fmt.Println("OnPacketProcessed")
// }

// // OnSubscribed method
// func (h *Hook) OnSubscribed(cl *mqtt.Client, pk packets.Packet, reasonCodes []byte) {
// 	// Implement your logic here
// 	fmt.Println("OnSubscribed")
// }

// // OnSelectSubscribers method
// func (h *Hook) OnSelectSubscribers(subs *mqtt.Subscribers, pk packets.Packet) *mqtt.Subscribers {
// 	// Implement your logic here
// 	fmt.Println("OnSelectSubscribers")
// 	return subs // Placeholder return value, change as per your implementation
// }

// // OnUnsubscribed method
// func (h *Hook) OnUnsubscribed(cl *mqtt.Client, pk packets.Packet) {
// 	// Implement your logic here
// 	fmt.Println("OnUnsubscribed")
// }

// // OnPublishDropped method
// func (h *Hook) OnPublishDropped(cl *mqtt.Client, pk packets.Packet) {
// 	// Implement your logic here
// 	fmt.Println("OnPublishDropped")
// }

// // OnRetainMessage method
// func (h *Hook) OnRetainMessage(cl *mqtt.Client, pk packets.Packet, r int64) {
// 	// Implement your logic here
// 	fmt.Println("OnRetainMessage")
// }

// // OnRetainPublished method
// func (h *Hook) OnRetainPublished(cl *mqtt.Client, pk packets.Packet) {
// 	// Implement your logic here
// 	fmt.Println("OnRetainPublished")
// }

// // OnQosPublish method
// func (h *Hook) OnQosPublish(cl *mqtt.Client, pk packets.Packet, sent int64, resends int) {
// 	// Implement your logic here
// 	fmt.Println("OnQosPublish")
// }

// // OnQosComplete method
// func (h *Hook) OnQosComplete(cl *mqtt.Client, pk packets.Packet) {
// 	// Implement your logic here
// 	fmt.Println("OnQosComplete")
// }

// // OnQosDropped method
// func (h *Hook) OnQosDropped(cl *mqtt.Client, pk packets.Packet) {
// 	// Implement your logic here
// 	fmt.Println("OnQosDropped")
// }

// // OnPacketIDExhausted method
// func (h *Hook) OnPacketIDExhausted(cl *mqtt.Client, pk packets.Packet) {
// 	// Implement your logic here
// 	fmt.Println("OnPacketIDExhausted")
// }

// // OnWill method
// func (h *Hook) OnWill(cl *mqtt.Client, will mqtt.Will) (mqtt.Will, error) {
// 	// Implement your logic here
// 	fmt.Println("OnWill")
// 	return mqtt.Will{}, nil // Placeholder return values, change as per your implementation
// }

// // OnWillSent method
// func (h *Hook) OnWillSent(cl *mqtt.Client, pk packets.Packet) {
// 	// Implement your logic here
// 	fmt.Println("OnWillSent")
// }

// // OnClientExpired method
// func (h *Hook) OnClientExpired(cl *mqtt.Client) {
// 	// Implement your logic here
// 	fmt.Println("OnClientExpired")
// }

// // OnRetainedExpired method
// func (h *Hook) OnRetainedExpired(filter string) {
// 	// Implement your logic here
// 	fmt.Println("OnRetainedExpired")
// }

// // StoredClients method
// func (h *Hook) StoredClients() ([]storage.Client, error) {
// 	// Implement your logic here
// 	fmt.Println("StoredClients")
// 	return nil, nil // Placeholder return values, change as per your implementation
// }

// // StoredSubscriptions method
// func (h *Hook) StoredSubscriptions() ([]storage.Subscription, error) {
// 	// Implement your logic here\
// 	fmt.Println("StoredSubscriptions")
// 	return nil, nil // Placeholder return values, change as per your implementation
// }

// // StoredInflightMessages method
// func (h *Hook) StoredInflightMessages() ([]storage.Message, error) {
// 	// Implement your logic here
// 	fmt.Println("StoredInflightMessages")
// 	return nil, nil // Placeholder return values, change as per your implementation
// }

// // StoredRetainedMessages method
// func (h *Hook) StoredRetainedMessages() ([]storage.Message, error) {
// 	// Implement your logic here
// 	fmt.Println("StoredRetainedMessages")
// 	return nil, nil // Placeholder return values, change as per your implementation
// }
