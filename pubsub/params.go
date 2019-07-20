package pubsub

import (
	"time"
)

var sleepMills = 2 // Add blocking for subscriptions
var testWorkers = 10
var testTopics = 10
var testMsgs = 100
var testSubs = 1000
var testTimeout = time.Second * 60
