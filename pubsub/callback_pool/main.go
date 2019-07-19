package callback_pool

// Topic is a simple pubsub struct
type Topic struct {
	subs       []*func(interface{})
	pubs       chan Work
	end        chan bool
	workers    []*Worker
	workerChan chan chan Work
}

// NewTopic creates a topic with no subscribers
func NewTopic(workers int) *Topic {
	t := &Topic{
		pubs:       make(chan Work),
		end:        make(chan bool),
		workerChan: make(chan chan Work),
	}
	for i := 0; i < workers; i++ {
		t.workers = append(t.workers, NewWorker(t.workerChan))
	}
	t.Start()
	return t
}

func (t *Topic) Start() {
	go func() {
		for {
			select {
			case <-t.end:
				for _, w := range t.workers {
					w.Stop() // stop worker
				}
				return
			case work := <-t.pubs:
				worker := <-t.workerChan // wait for available channel
				worker <- work           // dispatch work to worker
			}
		}
	}()
}

// Subscribe a topic with a context, if context closed, subscribton ends
func (t *Topic) Subscribe(cb *func(interface{})) {
	t.subs = append(t.subs, cb)
}

// Unsubscribe ends subscription for channel
func (t *Topic) Unsubscribe(cb *func(interface{})) {
	for i, f := range t.subs {
		if f == cb {
			t.subs = append(t.subs[:i], t.subs[i+1:]...)
		}
	}
}

// Publish new value to all subscribers
func (t *Topic) Publish(msg interface{}) {
	for i, f := range t.subs {
		if *f == nil {
			t.subs = append(t.subs[:i], t.subs[i+1:]...)
			continue
		}
		t.pubs <- Work{
			data: msg,
			fn:   *f,
		}
	}
}

// Close all channels which subscribed and deletes them
func (t *Topic) Close() {
	t.end <- true
	// Do nothing...?
}
