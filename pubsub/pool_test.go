package pubsub

import (
	"context"
	"fmt"
	callback "github.com/roderm/benchmarks/pubsub/callback_pool"
	channels "github.com/roderm/benchmarks/pubsub/channels_pool"
	"sync"
	"testing"
	"time"
)

func BenchmarkTopicChannelPool(b *testing.B) {
	channelTestPool(b, b.N, testSubs, testMsgs)
}

func BenchmarkTopicCallbackPool(b *testing.B) {
	callbackTestPool(b, b.N, testSubs, testMsgs)
}

func BenchmarkSubsChannelPool(b *testing.B) {
	channelTestPool(b, testTopics, b.N, testMsgs)
}

func BenchmarkSubsCallbackPool(b *testing.B) {
	callbackTestPool(b, testTopics, b.N, testMsgs)
}

func BenchmarkMsgsChannelPool(b *testing.B) {
	channelTestPool(b, testTopics, testSubs, b.N)
}

func BenchmarkMsgsCallbackPool(b *testing.B) {
	callbackTestPool(b, testTopics, testSubs, b.N)
}

func channelTestPool(b *testing.B, topics int, subs int, msgs int) {
	ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
	finished := make(chan bool)
	wg := sync.WaitGroup{}
	for i := 0; i < topics; i++ {
		func(i int) {
			t := channels.NewTopic(testWorkers)
			for s := 0; s < subs; s++ {
				wg.Add(msgs)
				ch := t.Subscribe(ctx)
				go func(t *channels.Topic, ch <-chan interface{}) {
					for m := 0; m < msgs; m++ {
						time.Sleep(time.Millisecond * time.Duration(sleepMills))
						<-ch
						wg.Done()
					}
				}(t, ch)
			}
			go func(i int, t *channels.Topic) {
				for m := 0; m < msgs; m++ {
					t.Publish(fmt.Sprintf("Topic: %d - Msg: %d", i, m))
				}
			}(i, t)
		}(i)
	}
	go func() {
		wg.Wait()
		finished <- true
	}()
	select {
	case <-finished:
		cancel()
	case <-ctx.Done():
		b.Error("failed finishing: run in timeout")
	}
}
func callbackTestPool(b *testing.B, topics int, subs int, msgs int) {
	ctx, _ := context.WithTimeout(context.Background(), testTimeout)
	finished := make(chan bool)
	wg := sync.WaitGroup{}
	for i := 0; i < topics; i++ {
		func(i int) {
			t := callback.NewTopic(testWorkers)
			for s := 0; s < subs; s++ {
				wg.Add(msgs)
				fn := func(interface{}) {
					time.Sleep(time.Millisecond * time.Duration(sleepMills))
					wg.Done()
				}
				t.Subscribe(&fn)
			}
			go func(i int, t *callback.Topic) {
				for m := 0; m < msgs; m++ {
					t.Publish(fmt.Sprintf("Topic: %d - Msg: %d", i, m))
				}
			}(i, t)
		}(i)
	}
	go func() {
		wg.Wait()
		finished <- true
	}()
	select {
	case <-finished:
	case <-ctx.Done():
		b.Error("failed finishing: run in timeout")
	}
}
