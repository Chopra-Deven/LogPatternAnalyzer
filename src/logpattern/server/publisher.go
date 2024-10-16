package server

import (
	"LogPattern/utils"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/golang/snappy"
	zmq "github.com/pebbe/zmq4"
)

var publisherLogger = utils.NewLogger("Publisher", "publisher")

type Publisher struct {
	socket *zmq.Socket

	Close bool
}

func GetPublisher() *Publisher {

	if publisher == nil {

		publisher = &Publisher{
			Close: false,
		}

		return publisher
	}

	return publisher
}

func (publisher *Publisher) Start() {

	publisher.socket, _ = zmq.NewSocket(zmq.PUSH)

	_ = publisher.socket.SetLinger(0)

	err := publisher.socket.Connect(string("tcp://" + utils.GetHost() + ":" + utils.GetPublisherPort()))

	if err != nil {
		panic(err) // TODO -- need to handle this properly
	}

	go func() {

		for {

			select {

			case response := <-utils.DetectedLogPatternResponse:

				_, err := publisher.socket.SendBytes(prepareEvent(response, "log.pattern.detector "), zmq.DONTWAIT)

				if err != nil {

					publisherLogger.Info(utils.MotadataString(fmt.Sprintf("error %v occured while sending publisher response ", err)))
				}
			}
		}

	}()
}

func prepareEvent(response utils.MotadataMap, topic utils.MotadataString) []byte {

	bytes, err := json.MarshalIndent(response, "", "  ")

	if err != nil {

		panic(err) // TODO-- need to handle this properly
	}

	byteArray := make([]byte, len(bytes))

	byteArray = snappy.Encode(byteArray, bytes)

	buffer := make([]byte, 2+len(topic)+len(byteArray))

	binary.LittleEndian.PutUint16(buffer, uint16(len(topic)))

	copy(buffer[2:], topic)

	copy(buffer[2+len(topic):], byteArray)

	return buffer
}
