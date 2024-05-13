package server

import (
	"context"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/JimYcod3x/meter_server/internal/meter"
	"github.com/JimYcod3x/meter_server/internal/utils"
	"github.com/JimYcod3x/meter_server/models"
	"github.com/go-redis/redis/v8"
	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
	"github.com/qustavo/dotsql"
)

const DefaultKey = "69aF7&3KY0_kk89@"

var (
	ctx = context.Background()
 Meters models.Meter
)

type Hook struct {
	mqtt.HookBase
	config *HookOptions
}

type HookOptions struct {
	Server *mqtt.Server
	db     *sql.DB
	rdb    *redis.Client
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
	// if h.config.db == nil {
	// 	return mqtt.ErrInvalidConfigType
	// }
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
	// h.config.rdb.Set(ctx,"test",241, 0)
	// result, _ := h.config.rdb.Get(ctx, "test").Result()
	// fmt.Println("result", result)
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
	// if cl.ID != "inline" {
	h.Log.Info(fmt.Sprintf("published payload: %v", hex.EncodeToString(pk.Payload)))
	h.Log.Info(fmt.Sprintf("published payload: %v %v", cl.ID, len(pk.Payload)))

	//
	// _  = h.config.Server.Publish("J23P000078S2C", []byte("test"), false, 0)

	// }

	// if err != nil {
	// 	go func() {
	// 		defer recover()
	// 	}()
	// }
	return pk, nil
}

func (h *Hook) OnPublished(cl *mqtt.Client, pk packets.Packet) {
	if cl.ID != "inline" && len(pk.Payload) >= 16 {
		db := h.config.db
		// if getMeterID from db else discard

		// 	if getdatakey(getMeterID) from db else (if decrypt from default key get the request keyexchange command=> keyexchange)
		// 		decrypt(datakey) => data transfer else decrypt(master key) get request change key command => response new datakey to meter =>
		pwd, _ := os.Getwd()
		dot, err := dotsql.LoadFromFile(pwd + "/sql/meter.sql")
		if err != nil {
			fmt.Println("can not load the sql script")
		}
		dot.Exec(db, "switch-to-database")
		meterID := utils.GetMeterIDFromTopic(pk)
		fmt.Println("get id from topic", meterID)
		payload := pk.Payload
		meterExis, _ := h.config.rdb.Get(ctx, "*"+meterID).Result()
		if len(meterExis) == 0 {
			fmt.Println("meter not found in rdb")
			// err := h.config.db.First(&Meters, "meter_id = ?", meterID).Error
			res, _ := dot.QueryRow(db, "find-one-meter-by-meter_id", meterID)
			err = res.Scan(&Meters.MeterID)
			if err != nil {
				fmt.Println("can not get the record from db")
				return
			}
			fmt.Println("what is the meter.meterid", Meters.MeterID)
			if Meters.MeterID != meterID || err != nil {
				fmt.Println("Meter can not found in db, discard")
				return
			}
		}
		// get datakey from db
		dataKey, _ := h.config.rdb.Get(ctx, "dk_"+meterID).Result()
		if len(dataKey) == 0 {
			fmt.Println("datakey can not found in rdb")
			_, err := dot.QueryRow(db, "find-one-meter-dk-by-meter_id", meterID)
			if err != nil {
				fmt.Println("data key can not found in db")
				// get master key from db
				masterKey, _ := h.config.rdb.Get(ctx, "mk_"+meterID).Result()
				if len(masterKey) == 0 {
					fmt.Println("master key can not found in rdb")
					_, err := dot.QueryRow(db, "find-one-meter-mk-by-meter_id", meterID)
					if err != nil {
						fmt.Println("master key can not found in db")
						// use default key to decrypt

						valid := utils.ValidateMeter(meterID, payload, DefaultKey)
						fmt.Println("testvvalid", valid)
						if valid {
							decryptByteDF, _ := utils.DecryptByte(payload, DefaultKey)
							command := utils.GetUSCommandFromDecrypt(decryptByteDF)
							commandParam := decryptByteDF[8]
							fmt.Printf("command df: %04b\n", command)
							fmt.Println(command == meter.ReqACK)
							fmt.Println(meter.ReqACK)
							fmt.Println(int(commandParam) == (meter.USCommandSet.ReqACK["ReqRegistration"]).(int))
							fmt.Println(meter.USCommandSet.ReqACK["ReqRegistration"])
							fmt.Println(command == meter.ReqACK && int(commandParam) == (meter.USCommandSet.ReqACK["ReqRegistration"]).(int))
							if command == meter.ReqACK && int(commandParam) == (meter.USCommandSet.ReqACK["ReqRegistration"]).(int) {
								fmt.Println("ReqRegistration")
								publishPayload := meter.KeyXFn(pk, DefaultKey, "ReqRegistration")
								publishPayloadByte, err := utils.EncryptPadding(publishPayload, DefaultKey)
								if err != nil {
									h.Log.Info("can not encrypt the payload")
									return
								}
								fmt.Println("DownStream Topic DF: ", utils.DSTopic(pk))
								h.config.Server.Publish(utils.DSTopic(pk), publishPayloadByte, false, 0)
								return
							}
							materKey := meterID + "000000"
							if command == meter.ReqACK && int(commandParam) == (meter.USCommandSet.ReqACK["ReqSucACK"]).(int) {
								fmt.Println("req ack")
								err := h.config.rdb.Set(ctx, "mk_"+meterID, materKey, 24*time.Hour).Err()
								if err != nil {
									fmt.Println("can not save to rdb")
									return
								}
								// err = h.config.db.Model(&models.Meter{
								// stmt, err := db.Prepare("UPDATE meter SET master_key = ? WHERE mete_id = ?")
								// if err != nil {
								// 	fmt.Println("can not save mk to db")
								// 	return
								// }
								// defer stmt.Close()

								// _, err = stmt.Exec(materKey, meterID)
								if err = utils.UpdateKeyToDb(db, "master_key", materKey, meterID); err != nil {
									return
								}
								err = h.config.rdb.Set(ctx, "mk_"+meterID, materKey, 24*time.Hour).Err()
								if err != nil {
									fmt.Println("can not save to rdb")
									return
								}
								fmt.Println("save to db")
							}
							return
						}
						fmt.Println("can not decrypt from the dk discard...")
						return
					}
				}
				// decrypt the payload (masterKey)
				fmt.Println("masterkey", masterKey)
				// masterKey := keys
				fmt.Println("jsadkflja", masterKey)
				valid := utils.ValidateMeter(meterID, payload, masterKey)
				if valid {
					fmt.Println("master key decrypt & get the request change key command")
					decryptByteMK, _ := utils.DecryptByte(payload, masterKey)
					command := utils.GetUSCommandFromDecrypt(decryptByteMK)
					fmt.Printf("command mk: %04b\n", command)
					commandParam := decryptByteMK[8]
					fmt.Println("commandParam: ", commandParam)
					if command == meter.ReqACK && int(commandParam) == (meter.USCommandSet.ReqACK["ReqChangeKey"]).(int) {
						publishPayload := meter.KeyXFn(pk, masterKey, "ReqChangeKey")
						fmt.Println("preencrypt: ", publishPayload)
						publishPayloadByte, err := utils.EncryptPadding(publishPayload, masterKey)
						if err != nil {
							h.Log.Info("can not encrypt the payload")
						}
						fmt.Println("DownStream Topic MK: ", utils.DSTopic(pk))
						h.config.Server.Publish(utils.DSTopic(pk), publishPayloadByte, false, 0)
						return
					}
				}
			}
		}
		// dataKey := keys
		// decrypt the payload(dataKey)
		dataKey, _ = utils.FindDateKey(meterID)
		fmt.Println("jsadkflja", dataKey)
		valid := utils.ValidateMeter(meterID, payload, dataKey)
		fmt.Println("meterID: ", meterID)
		if valid {
			fmt.Println("valid")
			decryptByteDK, _ := utils.DecryptByte(payload, dataKey)
			command := utils.GetUSCommandFromDecrypt(decryptByteDK)
			fmt.Printf("command mk: %04b\n", command)
			commandParam := decryptByteDK[8]
			if command == meter.ReqACK && int(commandParam) == (meter.USCommandSet.ReqACK["ReqSucACK"]).(int) {
				if utils.UpdateKeyToDb(db, "data_key", dataKey, meterID); err != nil {
					return
				}
				err = h.config.rdb.Set(ctx, "dk_"+meterID, dataKey, 24*time.Hour).Err()
				if err != nil {
					fmt.Println("can not save to rdb")
					return
				}
				fmt.Println("save to db")
				publishPayload := meter.KeyXFn(pk, dataKey, "ReqSucACK")
				publishPayloadByte, err := utils.EncryptPadding(publishPayload, dataKey)
				if err != nil {
					h.Log.Info("can not encrypt the payload")
				}
				h.config.Server.Publish(utils.DSTopic(pk), publishPayloadByte, false, 0)
				return
			}
		}
		masterKey, _ := h.config.rdb.Get(ctx, "mk_"+meterID).Result()
		if len(masterKey) == 0 {
			fmt.Println("can not get the master key in rdb")
			return
		}
		// masterKey := keys
		fmt.Println("jsadkflja11", masterKey)
		valid = utils.ValidateMeter(meterID, payload, masterKey)
		if valid {
			fmt.Println("master key decrypt & get the request change key command")
			decryptByteMK, _ := utils.DecryptByte(payload, masterKey)
			command := utils.GetUSCommandFromDecrypt(decryptByteMK)
			fmt.Printf("command mk: %04b\n", command)
			commandParam := decryptByteMK[8]
			fmt.Println("commandParam: ", commandParam)
			if command == meter.ReqACK && int(commandParam) == (meter.USCommandSet.ReqACK["ReqChangeKey"]).(int) {
				publishPayload := meter.KeyXFn(pk, masterKey, "ReqChangeKey")
				fmt.Println("preencrypt: ", publishPayload)
				publishPayloadByte, err := utils.EncryptPadding(publishPayload, masterKey)
				if err != nil {
					h.Log.Info("can not encrypt the payload")
				}
				fmt.Println("DownStream Topic MK: ", utils.DSTopic(pk))
				h.config.Server.Publish(utils.DSTopic(pk), publishPayloadByte, false, 0)
				return
			}
		}

		// _, err := h.config.rdb.Get(ctx, meterID).Result()
		// 111111
		// if err == nil {
		// 	// if utils.GetMeterIDFromDB(meterID) {
		// 	// dataKey, found := utils.GetSerDataKey(meterID)
		// 	dataKey, err := h.config.rdb.Get(ctx, "dk_"+meterID).Result()
		// 	if err == nil && err == redis.Nil {
		// 		err = h.config.db.Model(models.Meter{}).Select("data_key").Where("meter_id=?", meterID).Find(&meters).Error
		// 		if err != nil {

		// 		}
		// 		// fmt.Println("can not find the data key in db")
		// 		// masterKey, found := utils.GetSerMasterKey(meterID)
		// 		// if !found {

		// 	}

		// 	masterKey, _ := utils.GetSerMasterKey(meterID)

		// }
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
