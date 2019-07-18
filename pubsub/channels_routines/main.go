package channels

import (
	"context"
)

// Topic is a simple pubsub struct
type Topic struct {
	subs map[chan interface{}]context.Context
}

// NewTopic creates a topic with no subscribers
func NewTopic() *Topic {
	return &Topic{
		subs: make(map[chan interface{}]context.Context),
	}
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
		go func(ctx context.Context, ch chan interface{}) {
			select {
			case ch <- msg:
				return
			case <-ctx.Done():
				go t.Unsubscribe(ch)
			}
		}(ctx, ch)
	}
}

// Close all channels which subscribed and deletes them
func (t *Topic) Close() {
	for ch := range t.subs {
		close(ch)
		delete(t.subs, ch)
	}
}
