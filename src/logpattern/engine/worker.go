package engine

import (
	"LogPattern/store"
	"LogPattern/utils"
	"fmt"
)

var logger = utils.NewLogger("Worker", "worker")

type Worker struct {
	id         int
	shutdown   chan struct{}
	close      bool
	tokenizers []*utils.Tokenizer
}

func NewWorker(id int) *Worker {

	var tokenizers = make([]*utils.Tokenizer, 2)

	for index := range tokenizers {

		tokenizers[index] = utils.NewTokenizer(100)

	}

	return &Worker{
		id:         id,
		shutdown:   make(chan struct{}),
		tokenizers: tokenizers,
	}
}

func (worker *Worker) Start() {

	utils.WaitGroup.Add(1)

	defer utils.WaitGroup.Done()

	go func() {

		for {
			if !worker.close {
				worker.start()
			} else {
				return
			}
		}

	}()
}

func (worker *Worker) start() {

	defer func() {

		if r := recover(); r != nil {

			logger.Info(utils.MotadataString(fmt.Sprintf("Recovered from panic: %v", r)))

		}

	}()

	for {

		select {

		case request := <-utils.DetectLogPatternRequest:

			context := store.DetectPattern(request)

			utils.WaitGroup.Done()

			utils.DetectedLogPatternResponse <- context

			//utils.DetectedLogPatternResponse <- store.DetectPattern(request, worker.tokenizers)

		case <-worker.shutdown:
			return
		}

	}

}

func (worker *Worker) Stop() {

	worker.close = true

	worker.shutdown <- struct{}{}
}
