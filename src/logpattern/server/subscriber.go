/*
 * Copyright (c) Motadata 2024.  All rights reserved.
 */

package server

import (
	"LogPattern/utils"
	"encoding/json"
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"strings"
	"time"
)

var subscriberLogger = utils.NewLogger("Subscriber", "subscriber")

type Subscriber struct {
	socket *zmq.Socket

	Close bool
}

func GetSubscriber() *Subscriber {

	if subscriber == nil {

		subscriber = &Subscriber{
			Close: false, // TODO -- need to know what does this means...
		}

		var err error

		subscriber.socket, err = zContext.NewSocket(zmq.SUB)

		_ = subscriber.socket.SetSubscribe("log.pattern.detector ")

		if err != nil {

			subscriberLogger.Info(utils.MotadataString(fmt.Sprintf("failed to start subscriber , reason :%v", err.Error())))

			return nil
		}

		subscriberLogger.Info(utils.MotadataString(fmt.Sprintf("subscribed topic for log pattern")))

	}

	return subscriber
}

func (subscriber *Subscriber) Start() {

	err := subscriber.socket.Connect(string("tcp://" + utils.GetHost() + ":" + utils.GetSubscriberPort()))

	if err != nil {

		subscriberLogger.Info(utils.MotadataString(fmt.Sprintf("failed to start subscriber , reason :%v", err.Error())))
	}

	go func() { // insertion zmq routine

		defer func() {

			if err := recover(); err != nil {

				subscriberLogger.Fatal(utils.MotadataString(fmt.Sprintf("error %v occurred in data-write subscriber...", err)))

				_ = subscriber.socket.Close()
			}

			subscriberLogger.Info("data-write subscriber exiting...")
		}()

		for {

			if !subscriber.Close {

				bytes, err := subscriber.socket.RecvBytes(0)

				if IsSocketTimeoutError(err) {

					continue
				}

				if err == nil {

					if len(bytes) > 0 {

						if utils.TraceEnabled() {

							subscriberLogger.Trace("sending event to request channel")
						}

						event := utils.MotadataMap{}

						err = json.Unmarshal(bytes, &event)

						if err != nil {

							//subscriberLogger.Fatal(utils.MotadataString(fmt.Sprintf("Unable to Unmarshal event with reason : %s", err.Error())))

							continue
						}

						utils.DetectLogPatternRequest <- event

					} else {

						subscriberLogger.Fatal("failed to send event to request channel, reason : request is empty")
					}

				} else {

					if strings.Contains(err.Error(), "Context was terminated") {

						_ = subscriber.socket.Close()

						subscriberLogger.Info("data-write socket closed, reason: zmq context was terminated")

						return
					}

					subscriberLogger.Fatal(utils.MotadataString(fmt.Sprintf("failed to subscribe data-write request,reason : error %v occured", err)))

				}

			} else {

				err = subscriber.socket.Close()

				if err != nil {

					subscriberLogger.Fatal(utils.MotadataString(fmt.Sprintf("error occured while closing data-write socket reason : %s", err.Error())))
				}

				if subscriber.Close {

					subscriberLogger.Info("data-write socket closed, reason : subscriber close event received")

				} else {

					subscriberLogger.Info("data-write socket closed, reason : shutdown event received")

				}

				return
			}

		}
	}()

}

func (subscriber *Subscriber) ShutDown() {

	if !subscriber.Close {

		subscriber.Close = true

		time.Sleep(utils.SubscriberRCVTimeout + time.Second*2)
	}

}

func IsSocketTimeoutError(err error) bool {

	if err != nil {

		return strings.Contains(err.Error(), "resource temporarily unavailable")

	}

	return false
}
