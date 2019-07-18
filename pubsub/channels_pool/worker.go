package channels_pool

type Work struct {
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
				job.ch <- job.data
			case <-w.end:
				return
			}
		}
	}()
}

func (w *Worker) End() {
	w.end <- true
}
