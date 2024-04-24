package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/JimYcod3x/meter_server/internal/server"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/config"
	"github.com/mochi-mqtt/server/v2/packets"
)
func main() {

	flag.Parse()

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

	sev := server.Run(options)
	fmt.Println("server started: ", sev)
	go func() {
	
	err := sev.Serve()
	if err != nil {
		log.Fatal(err)
	}
	}()

	go func() {
		// Inline subscriptions can also receive retained messages on subscription.
		_ = sev.Publish("direct/retained", []byte("retained message"), true, 0)
		_ = sev.Publish("direct/alternate/retained", []byte("some other retained message"), true, 0)

		// Subscribe to a filter and handle any received messages via a callback function.
		callbackFn := func(cl *mqtt.Client, sub packets.Subscription, pk packets.Packet) {
			sev.Log.Info("inline client received message from subscription", "client", cl.ID, "subscriptionId", sub.Identifier, "topic", pk.TopicName, "payload", string(pk.Payload))
		}
		sev.Log.Info("inline client subscribing")
		_ = sev.Subscribe("direct/#", 1, callbackFn)
		_ = sev.Subscribe("direct/#", 2, callbackFn)
	}()


	<-done
	sev.Log.Warn("caught signal, stopping...")
	_ = sev.Close()
	sev.Log.Info("mqtt server stopped")
}
