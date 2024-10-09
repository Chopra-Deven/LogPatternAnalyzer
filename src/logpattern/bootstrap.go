package main

import (
	"LogPattern/engine"
	"LogPattern/job"
	"LogPattern/store"
	"LogPattern/utils"
	"bufio"
	"fmt"
	_ "net/http/pprof"
	"os"
	"time"
)

var (
	logger = utils.NewLogger("Bootstrap", "bootstrap")
)

/*
func main() {

	defer func() {

		if err := recover(); err != nil {

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

		cleanUpJob := job.NewPersistenceJob(1)

		cleanUpJob.Start()

		workers = make([]*engine.Worker, utils.GetMaxWorker())

		utils.DetectLogPatternRequest = make(chan utils.MotadataMap, utils.GetMaxChannelBuffer())

		for index := 0; index < 10; index++ {

			logger.Info(utils.MotadataString(fmt.Sprintf("Worker %d initiated", index)))

			workers[index] = engine.NewWorker(index)

			workers[index].Start()
		}

	}

	signal.Notify(killSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)

	<-killSignal

	server.Shutdown(workers, cleanUpJob)
}
*/

func main() {

	//file, err := os.Open("samplelogfiles/Linux_Test_Logs.txt") // Replace with your log file name
	//file, err := os.Open("/home/deven/motadata-workspace/LogPattern/samplelogfiles/Linux Bulk Logs.txt") // Replace with your log file name
	file, err := os.Open("/home/deven/motadata-workspace/LogPattern/samplelogfiles/Linix60k.txt") // Replace with your log file name
	//file, err := os.Open("/home/deven/motadata-workspace/LogPattern/samplelogfiles/Linix60k.txt") // Replace with your log file name
	//file, err := os.Open("/home/deven/motadata-workspace/LogPattern/samplelogfiles/Linix60k.txt") // Replace with your log file name

	records := 60000
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

	utils.DetectLogPatternRequest = make(chan utils.MotadataMap, 100000)

	utils.DetectedLogPatternResponse = make(chan utils.MotadataMap, 100000)

	utils.WaitGroup.Add(len(logs))

	for _, log := range logs {

		event := utils.MotadataMap{}

		event["plugin.id"] = "500009"

		event["event.category"] = "Other"

		event["message"] = log

		utils.DetectLogPatternRequest <- event

		//store.DetectPattern(event)
	}

	utils.WaitGroup.Wait()

	t := time.Now()

	fmt.Printf("\nTotal time taken : %v", t.Sub(start))

	store.Flush(utils.CurrentDir + utils.PathSeparator + utils.ConfigDirectory + utils.PathSeparator + "log-patterns")

	//fmt.Printf("\nTotal pattern detected items : %v", store.GetItemsLength())
	//fmt.Printf("\nTotal pattern detected Patterns : %v %s", patterns., "\n")

	//Convert the standard map to JSON
	//jsonData, err := json.MarshalIndent(store.GetStandardMap(), "", "  ")
	//if err != nil {
	//	fmt.Println("Error converting to JSON:", err)
	//	return
	//}

	//store.PrintPatterIDs()
	//
	//_ = jsonData

	//Print the JSON output
	//fmt.Println(string(jsonData))

}
