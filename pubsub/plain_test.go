package pubsub

import (
	"context"
	"fmt"
	"github.com/roderm/benchmarks/pubsub/callback"
	"github.com/roderm/benchmarks/pubsub/channels"
	"sync"
	"testing"
)

func BenchmarkTopicChannel(b *testing.B) {
	channelTest(b, b.N, testSubs, testMsgs)
}

func BenchmarkTopicCallback(b *testing.B) {
	callbackTest(b, b.N, testSubs, testMsgs)
}

func BenchmarkSubsChannel(b *testing.B) {
	channelTest(b, testTopics, b.N, testMsgs)
}

func BenchmarkSubsCallback(b *testing.B) {
	callbackTest(b, testTopics, b.N, testMsgs)
}

/*
func BenchmarkMsgsChannel(b *testing.B) {
	channelTest(b, testTopics, testSubs, b.N)
}

func BenchmarkMsgsCallback(b *testing.B) {
	callbackTest(b, testTopics, testSubs, b.N)
}
*/
func channelTest(b *testing.B, topics int, subs int, msgs int) {
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
func callbackTest(b *testing.B, topics int, subs int, msgs int) {
	ctx, _ := context.WithTimeout(context.Background(), testTimeout)
	finished := make(chan bool)
	wg := sync.WaitGroup{}
	for i := 0; i < topics; i++ {
		func(i int) {
			t := callback.NewTopic()
			for s := 0; s < subs; s++ {
				wg.Add(msgs)
				fn := func(interface{}) {
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
