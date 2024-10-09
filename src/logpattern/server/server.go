package server

import (
	"LogPattern/engine"
	"LogPattern/job"
	"LogPattern/utils"
	zmq "github.com/pebbe/zmq4"
	"time"
)

var (
	subscriber *Subscriber

	publisher *Publisher

	zContext *zmq.Context

	logger = utils.NewLogger("Server", "server")
)

func Start() bool {
	var err error

	zContext, err = zmq.NewContext()

	if err != nil {
		panic(err)
	}

	time.Sleep(1 * time.Second)

	err = zContext.SetIoThreads(1) // TODO - need to know the exact use case of this parameter.

	if err != nil {
		panic(err)
	}

	time.Sleep(1 * time.Second)

	subscriber = GetSubscriber()

	publisher = GetPublisher()

	if subscriber == nil && publisher == nil {

		return false
	}

	utils.DetectLogPatternRequest = make(chan utils.MotadataMap, 100000)

	utils.DetectedLogPatternResponse = make(chan utils.MotadataMap, 100000)

	subscriber.Start()

	publisher.Start()

	return true
}

func Shutdown(workers []*engine.Worker, job *job.PersistenceJob) {

	for _, worker := range workers {

		worker.Stop()
	}

	job.Stop()

	utils.WaitGroup.Wait()

}
