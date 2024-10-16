package job

import (
	"LogPattern/store"
	"LogPattern/utils"
	"fmt"
	"time"
)

var logger = utils.NewLogger("Job", "cleanUpJob")

type PersistenceJob struct {
	id       int
	shutdown chan struct{}
	close    bool
}

func NewPersistenceJob(id int) *PersistenceJob {

	return &PersistenceJob{
		id:       id,
		shutdown: make(chan struct{}),
	}
}

func (job *PersistenceJob) Start() {

	utils.WaitGroup.Add(1)

	defer utils.WaitGroup.Done()

	go func() {

		for {
			if !job.close {
				job.start()
			} else {
				return
			}
		}

	}()
}

func (job *PersistenceJob) start() {

	defer func() {

		if r := recover(); r != nil {

			logger.Info(utils.MotadataString(fmt.Sprintf("Recovered from panic: %v", r)))

		}

	}()

	logger.Info(utils.MotadataString(fmt.Sprintf("Starting job #%d", job.id)))

	ticker := time.NewTicker(time.Duration(utils.GetFlushTimer()) * time.Second)

	for {

		select {

		case <-ticker.C:

			store.Flush(utils.CurrentDir + utils.PathSeparator + utils.ConfigDirectory + utils.PathSeparator + "log-patterns")

		case <-job.shutdown:

			store.Flush(utils.CurrentDir + utils.PathSeparator + utils.ConfigDirectory + utils.PathSeparator + "log-patterns")

			return
		}

	}

}

func (job *PersistenceJob) Stop() {

	job.close = true

	job.shutdown <- struct{}{}
}
