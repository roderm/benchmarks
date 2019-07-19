package channels_pool

import (
	"context"
	"log"
)

type Work struct {
	ctx  context.Context
	ch   chan interface{}
	data interface{}
}
type Worker struct {
	workerChan chan chan Work
	channel    chan Work
	end        chan bool
}

func NewWorker(workers chan chan Work) *Worker {
	w := &Worker{
		workerChan: workers,
		channel:    make(chan Work),
		end:        make(chan bool),
	}
	w.Start()
	return w
}

func (w *Worker) Start() {
	go func() {
		for {
			w.workerChan <- w.channel
			select {
			case job := <-w.channel:
				select {
				case <-job.ctx.Done():
					log.Fatalf("Channel already closed: %v", job.data)
				case job.ch <- job.data: // value never read
				}
			case <-w.end:
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	w.end <- true
}
