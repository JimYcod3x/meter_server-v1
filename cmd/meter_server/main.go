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




	<-done
	sev.Log.Warn("caught signal, stopping...")
	_ = sev.Close()
	sev.Log.Info("mqtt server stopped")
}
