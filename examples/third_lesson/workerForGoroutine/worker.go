package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

const (
	UNFINISHED = 0
	FINISHED   = 1
)

type Worker interface {
	Work()
}

type WorkerImpl struct {
	ID     int
	Task   string
	Result string
	Done   int32
}

func (w *WorkerImpl) Work() {
	time.Sleep(3 * time.Second)
	w.Result = fmt.Sprintf("Worker %d completed task: %s", w.ID, w.Task)
	atomic.StoreInt32(&w.Done, FINISHED)
}

func main() {
	numWorkers := 8
	workers := make([]*WorkerImpl, numWorkers)

	startTime := time.Now()
	for i := 0; i < numWorkers; i++ {
		workers[i] = &WorkerImpl{
			ID:     i,
			Task:   "Print",
			Result: "",
			Done:   0,
		}

		go func(w *WorkerImpl) {
			w.Work()
			fmt.Println("A worker finished a job!")
		}(workers[i])
	}

	// Use Atomic number to wait goroutines
	for {
		doneCount := 0
		for _, worker := range workers {
			done := atomic.LoadInt32(&worker.Done)
			doneCount += int(done)
		}
		if doneCount == numWorkers {
			break
		}
	}
	endTime := time.Now()
	fmt.Println("Time cost: ", endTime.Second()-startTime.Second())
	fmt.Println("All workers finished jobs!")
}
