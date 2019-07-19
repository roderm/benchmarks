package channels_pool

import (
	"context"
)

// Topic is a simple pubsub struct
type Topic struct {
	subs       map[chan interface{}]context.Context
	pubs       chan Work
	end        chan bool
	workers    []*Worker
	workerChan chan chan Work
}

// NewTopic creates a topic with no subscribers
func NewTopic(workers int) *Topic {
	t := &Topic{
		subs:       make(map[chan interface{}]context.Context),
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
func (t *Topic) Subscribe(ctx context.Context) <-chan interface{} {
	ch := make(chan interface{})
	t.subs[ch] = ctx
	return ch
}

// Unsubscribe ends subscription for channel
func (t *Topic) Unsubscribe(ch chan interface{}) {
	delete(t.subs, ch)
}

// Publish new value to all subscribers
func (t *Topic) Publish(msg interface{}) {
	for ch, ctx := range t.subs {
		t.pubs <- Work{
			ctx:  ctx,
			ch:   ch,
			data: msg,
		}
	}
}

// Close all channels which subscribed and deletes them
func (t *Topic) Close() {
	for _, w := range t.workers {
		w.Stop()
	}
	for ch := range t.subs {
		close(ch)
		delete(t.subs, ch)
	}
}
