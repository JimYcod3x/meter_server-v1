package server

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"github.com/JimYcod3x/meter_server/config"
	"github.com/JimYcod3x/meter_server/database"
	"github.com/JimYcod3x/meter_server/internal/meter"
	"github.com/JimYcod3x/meter_server/internal/utils"
	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
)

const (
	tSub1 = "J23P000078C2S"
	tSub2 = "J230008542C2S"
	tSub3 = "J200002335C2S"
	TPub  = "J23P000078S2C"
)

type Server struct {
	*mqtt.Server
}

func Run(options *mqtt.Options) (*mqtt.Server, *sql.DB){
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load environment variables", err)
	}

	db := database.ConnectionDB(&loadConfig)
	rdb := database.ConnectionRedisDb(&loadConfig)

	utils.InitDB(db)

	utils.CreateMeter(db, "J23P000078", "IoT")
	utils.CreateMeter(db, "J230008542", "PV")
	utils.CreateMeter(db, "J200002335", "PV")

	// err = utils.UpdateKeyToDb(db, "master_key", "J23P000078000000", "J23P000078")
	// if err != nil {
	// 	fmt.Println("what is the error", err)
	// }


	// var ctx = context.Background()
	// utils.SaveToRdb(rdb, ctx, "J23P000078","J23P000078"+ "000000")


	sev := mqtt.New(options)
	sev.AddHook(new(Hook), &HookOptions{
		Server: sev,
		db: db,
		rdb: rdb,
	})
	// client := sev.NewClient(nil, mqtt.LocalListener, mqtt.InlineClientId, true)
	// sev.Clients.Add(client)
	// TODO: only subscribe the meterid in db

	sev.Subscribe(tSub1, 1, subFn)
	sev.Subscribe(tSub2, 1, subFn)
	sev.Subscribe("/#", 1, subFn)
	go func() {
		time.Sleep(time.Second * 10)
		fmt.Println("publish the message after 10 second")
		sev.Publish(TPub, pubFn2(), true, 0)
	}()

	// sev.Publish(tPub, meter.KeyExchange(), false, 2)

	return sev, db
}

func subFn(cl *mqtt.Client, sub packets.Subscription, pk packets.Packet) {
	fmt.Println(sub.Filter)
	hexPayload := hex.EncodeToString([]byte(pk.Payload))
	fmt.Println("hex payload", hexPayload)

	// meter.GetMeterData(cl, pk)
}

func pubFn() []byte {
	meterType := meter.IoT
	meterDSCmd := meter.GetDataFromMeter
	// meterID := "J200002335"
	meterID := "J23P000078"
	dsCmdParam := meter.DSCommandSet.GetDataFromMeter["GetMeterID&VerInfo"]
	return meter.TestMeterFn(meterType, meterDSCmd, meterID, dsCmdParam)
}

// func PublishData()

func pubFn2() []byte {
	meterType := meter.IoT
	meterDSCmd := meter.GetDataFromMeter
	meterID := "J200002335"
	// meterID := "J23P000078"
	dsCmdParam := "Get30minRecordLog"
	startTime := "2024-05-02T15:15:00"
	endTime := "2024-05-02T15:15:01"
	return meter.CallGetDataFromMeter(meterType, meterDSCmd, meterID, dsCmdParam, startTime, endTime)
}
