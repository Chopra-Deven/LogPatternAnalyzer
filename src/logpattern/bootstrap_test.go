package main

import (
	"LogPattern/engine"
	"LogPattern/job"
	"LogPattern/store"
	"LogPattern/utils"
	"bufio"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestBootstrap(t *testing.T) {

	file, err := os.Open("/home/deven/motadata-workspace/LogPattern/samplelogfiles/Linux_Test_Logs.txt") // Replace with your log file name
	//file, err := os.Open("/home/deven/motadata-workspace/LogPattern/samplelogfiles/Linux Bulk Logs.txt") // Replace with your log file name
	//file, err := os.Open("/home/deven/motadata-workspace/LogPattern/samplelogfiles/Linix60k.txt") // Replace with your log file name
	//file, err := os.Open("/home/deven/motadata-workspace/LogPattern/samplelogfiles/vCenter100k_Logs.txt") // Replace with your log file name
	//file, err := os.Open("/home/deven/motadata-workspace/LogPattern/samplelogfiles/Linix60k.txt") // Replace with your log file name

	records := 5000

	if err != nil {

		fmt.Println("Error opening file:", err)

		return
	}
	defer file.Close()

	var logs []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()

		logs = append(logs, line)
	}

	if err := scanner.Err(); err != nil {

		fmt.Println("Error reading file:", err)
	}

	logs = logs[0:records]

	start := time.Now()

	utils.DetectLogPatternRequest = make(chan utils.MotadataMap, 100000)

	utils.DetectedLogPatternResponse = make(chan utils.MotadataMap, 100000)

	store.Init()

	cleanUpJob := job.NewPersistenceJob(1)

	cleanUpJob.Start()

	workers := make([]*engine.Worker, utils.GetMaxWorker())

	utils.DetectLogPatternRequest = make(chan utils.MotadataMap, utils.GetMaxChannelBuffer())

	for index := 0; index < 10; index++ {

		logger.Info(utils.MotadataString(fmt.Sprintf("Worker %d initiated", index)))

		workers[index] = engine.NewWorker(index)

		workers[index].Start()
	}

	utils.WaitGroup.Add(len(logs))

	for _, log := range logs {

		event := utils.MotadataMap{}

		event["plugin.id"] = "500009"

		event["event.category"] = "Other"

		event["message"] = log

		utils.DetectLogPatternRequest <- event
	}

	utils.WaitGroup.Wait()

	endTime := time.Now()

	fmt.Printf("\nTotal time taken : %v", endTime.Sub(start))

	store.Flush(utils.CurrentDir + utils.PathSeparator + utils.ConfigDirectory + utils.PathSeparator + "log-patterns")

}
