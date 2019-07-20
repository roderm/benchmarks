package pubsub

import (
	"context"
	"fmt"
	callback "github.com/roderm/benchmarks/pubsub/callback_routines"
	channels "github.com/roderm/benchmarks/pubsub/channels_routines"
	"sync"
	"testing"
	"time"
)

func BenchmarkTopicChannelRoutine(b *testing.B) {
	channelTestRoutine(b, b.N, testSubs, testMsgs)
}

func BenchmarkTopicCallbackRoutine(b *testing.B) {
	callbackTestRoutine(b, b.N, testSubs, testMsgs)
}

func BenchmarkSubsChannelRoutine(b *testing.B) {
	channelTestRoutine(b, testTopics, b.N, testMsgs)
}

func BenchmarkSubsCallbackRoutine(b *testing.B) {
	callbackTestRoutine(b, testTopics, b.N, testMsgs)
}

func BenchmarkMsgsChannelRoutine(b *testing.B) {
	channelTestRoutine(b, testTopics, testSubs, b.N)
}

func channelTestRoutine(b *testing.B, topics int, subs int, msgs int) {
	ctx, _ := context.WithTimeout(context.Background(), testTimeout)
	finished := make(chan bool)
	wg := sync.WaitGroup{}
	for i := 0; i < topics; i++ {
		func(i int) {
			t := channels.NewTopic()
			for s := 0; s < subs; s++ {
				wg.Add(msgs)
				ch := t.Subscribe(ctx)
				go func(t *channels.Topic) {
					for m := 0; m < msgs; m++ {
						time.Sleep(time.Millisecond * time.Duration(sleepMills))
						<-ch
						wg.Done()
					}
				}(t)
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
	case <-ctx.Done():
		b.Error("failed finishing: run in timeout")
	}
}
func callbackTestRoutine(b *testing.B, topics int, subs int, msgs int) {
	ctx, _ := context.WithTimeout(context.Background(), testTimeout)
	finished := make(chan bool)
	wg := sync.WaitGroup{}
	for i := 0; i < topics; i++ {
		func(i int) {
			t := callback.NewTopic()
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
