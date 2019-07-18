package callback

// Topic is a simple pubsub struct
type Topic struct {
	subs []*func(interface{})
}

// NewTopic creates a topic with no subscribers
func NewTopic() *Topic {
	return &Topic{}
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
		cb := *f
		cb(msg)
	}
}

// Close all channels which subscribed and deletes them
func (t *Topic) Close() {
	// Do nothing...?
}
