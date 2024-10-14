package utils

import (
	"sync"
	"time"
)

var (
	SubscriberRCVTimeout = time.Second * 5

	config MotadataMap

	WaitGroup = sync.WaitGroup{}
)

func InitConfig(configurations MotadataMap) {

	config = configurations

	SetLogLevel(config.GetIntValue(SystemLogLevel))
}

func GetMaxWorker() int {

	if config.Contains("max.worker") {

		return config.GetIntValue("max.worker")
	}

	return 10
}

func GetMaxChannelBuffer() int {

	if config.Contains("event.backlog.size") {

		return config.GetIntValue("event.backlog.size")
	}

	return 1000000
}

func GetHost() MotadataString {
	return "localhost"
}

func GetSubscriberPort() MotadataString {

	if config.Contains("event.subscriber.port") {

		return config.GetMotadataStringValue("event.subscriber.port")
	}
	return "8888"
}

func GetPprofPort() string {

	if config.Contains("pprof.port") {

		return config.GetStringValue("pprof.port")
	}
	return "6161"
}

func GetFlushTimer() int {

	if config.Contains("log.pattern.flush.timer.seconds") {

		return config.GetIntValue("log.pattern.flush.timer.seconds")
	}
	return 120
}

func GetPublisherPort() MotadataString {

	if config.Contains("event.publisher.port") {

		return config.GetMotadataStringValue("event.publisher.port")
	}
	return "8889"
}
