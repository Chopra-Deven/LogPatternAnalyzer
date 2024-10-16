package main

import (
	"LogPattern/engine"
	"LogPattern/job"
	"LogPattern/server"
	"LogPattern/store"
	"LogPattern/utils"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

var (
	logger = utils.NewLogger("Bootstrap", "bootstrap")
)

func main() {

	defer func() {

		if err := recover(); err != nil {
			stackTraceBytes := make([]byte, 1<<20)

			logger.Fatal(utils.MotadataString(fmt.Sprintf("Stack Trace of panic : \n%v", string(stackTraceBytes[:runtime.Stack(stackTraceBytes, false)]))))

			logger.Fatal(utils.MotadataString(fmt.Sprintf("Panic %v recovered", err)))
		}
	}()

	go func() {
		logger.Info(utils.MotadataString(fmt.Sprintf("Error in PPROF : %v", http.ListenAndServe("0.0.0.0:"+utils.GetPprofPort(), nil))))
	}()

	bytes, err := os.ReadFile(utils.CurrentDir + utils.PathSeparator + utils.ConfigDirectory + utils.PathSeparator + "motadata.json")

	if err != nil {

		panic(errors.New("config File not found"))
	}

	config := make(utils.MotadataMap)

	err = json.Unmarshal(bytes, &config)

	if err != nil {

		panic(utils.MotadataString(fmt.Sprintf("Config file is not in proper formate. Unmarshal Config File Error: %v", err)))
	}

	utils.InitConfig(config)

	killSignal := make(chan os.Signal, 1)

	var workers []*engine.Worker

	var cleanUpJob *job.PersistenceJob

	if server.Start() {

		store.Init()

		cleanUpJob = job.NewPersistenceJob(1)

		cleanUpJob.Start()

		workers = make([]*engine.Worker, utils.GetMaxWorker())

		for index := 0; index < utils.GetMaxWorker(); index++ {

			logger.Info(utils.MotadataString(fmt.Sprintf("Worker %d initiated", index)))

			workers[index] = engine.NewWorker(index)

			workers[index].Start()
		}

	}

	signal.Notify(killSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)

	<-killSignal

	logger.Info("Shutting down log-pattern process!")

	server.Shutdown(workers, cleanUpJob)
}
