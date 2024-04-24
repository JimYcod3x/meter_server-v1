package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/JimYcod3x/meter_server/internal/server"
	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/config"
	"github.com/mochi-mqtt/server/v2/hooks/auth"
	"github.com/mochi-mqtt/server/v2/packets"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		done <- true
	}()

	configBytes, err := os.ReadFile("config/configs.yml")
	if err != nil {
		log.Fatal(err)
	}

	options, err := config.FromBytes(configBytes)
	if err != nil {
		log.Fatal(err)
	}
	server := mqtt.New(options)
	

	_ = server.AddHook(new(auth.AllowHook), nil)

	// Add custom hook (ExampleHook) to the server
	err = server.AddHook(new(ExampleHook), &ExampleHookOptions{
		Server: server,
	})

	if err != nil {
		log.Fatal(err)
	}

	// Start the server
	go func() {
		err := server.Serve()
		if err != nil {
			log.Fatal(err)
		}
	}()

	<-done
	server.Log.Warn("caught signal, stopping...")
	_ = server.Close()
	server.Log.Info("main.go finished")
}

// Options contains configuration settings for the hook.
type ExampleHookOptions struct {
	Server *mqtt.Server
}

type ExampleHook struct {
	mqtt.HookBase
	config *ExampleHookOptions
}

func (h *ExampleHook) ID() string {
	return "events-example"
}

func (h *ExampleHook) Provides(b byte) bool {
	return bytes.Contains([]byte{
		mqtt.OnConnect,
		mqtt.OnDisconnect,
		mqtt.OnSubscribed,
		mqtt.OnUnsubscribed,
		mqtt.OnPublished,
		mqtt.OnPublish,
	}, []byte{b})
}

func (h *ExampleHook) Init(config any) error {
	h.Log.Info("initialised")
	if _, ok := config.(*ExampleHookOptions); !ok && config != nil {
		return mqtt.ErrInvalidConfigType
	}

	h.config = config.(*ExampleHookOptions)
	if h.config.Server == nil {
		return mqtt.ErrInvalidConfigType
	}
	return nil
}

func (h *ExampleHook) OnConnect(cl *mqtt.Client, pk packets.Packet) error {
	h.Log.Info("client connected", "client", cl.ID)
	return nil
}

func (h *ExampleHook) OnDisconnect(cl *mqtt.Client, err error, expire bool) {
	if err != nil {
		h.Log.Info("client disconnected", "client", cl.ID, "expire", expire, "error", err)
	} else {
		h.Log.Info("client disconnected", "client", cl.ID, "expire", expire)
	}

}

func (h *ExampleHook) OnSubscribed(cl *mqtt.Client, pk packets.Packet, reasonCodes []byte) {
	h.Log.Info(fmt.Sprintf("subscribed qos=%v", reasonCodes), "client", cl.ID, "filters", pk.Filters)
}

func (h *ExampleHook) OnUnsubscribed(cl *mqtt.Client, pk packets.Packet) {
	h.Log.Info("unsubscribed", "client", cl.ID, "filters", pk.Filters)
}

func (h *ExampleHook) OnPublish(cl *mqtt.Client, pk packets.Packet) (packets.Packet, error) {
	if cl.ID != "inline" {
		h.Log.Info("--> OnPublish from normal client", "client", cl.ID, "payload", string(pk.Payload))
		h.config.Server.Publish(server.TPub, []byte("xxxxx"), false, 0)
	} else {
		h.Log.Info("--> OnPublish from inline client", "client", cl.ID, "payload", string(pk.Payload))
	}
	return pk, nil
}

func (h *ExampleHook) OnPublished(cl *mqtt.Client, pk packets.Packet) {
	if cl.ID != "inline" {
		h.Log.Info("--> OnPublished from normal client", "client", cl.ID, "payload", string(pk.Payload))
	} else {
		h.Log.Info("--> OnPublished from inline client", "client", cl.ID, "payload", string(pk.Payload))
	}
}