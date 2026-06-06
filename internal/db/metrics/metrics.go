package metrics

import "sync/atomic"

var TasksProcessed uint64
var TasksFailed uint64
var TasksRetried uint64

func IncrementProcessed() {
	atomic.AddUint64(&TasksProcessed, 1)
}

func IncrementFailed() {
	atomic.AddUint64(&TasksFailed, 1)
}

func IncrementRetried() {
	atomic.AddUint64(&TasksRetried, 1)
}
