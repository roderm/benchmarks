package callback_pool

type Work struct {
	fn   func(interface{})
	data interface{}
}
type Worker struct {
	WorkerChannel chan chan Work
	Channel       chan Work
	End           chan bool
}

func NewWorker(workers chan chan Work) *Worker {
	w := &Worker{
		WorkerChannel: workers,
		Channel:       make(chan Work),
		End:           make(chan bool),
	}
	w.Start()
	return w
}
func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerChannel <- w.Channel // when the worker is available place channel in queue
			select {
			case job := <-w.Channel: // worker has received job
				job.fn(job.data)
			case <-w.End:
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	w.End <- true
}
