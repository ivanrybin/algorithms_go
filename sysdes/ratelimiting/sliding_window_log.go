package ratelimiting

import (
	"container/list"
	"time"
)

const slidingWindowLogRatePerSecond = 50

func slidingWindowLog(in <-chan struct{}) {
	log := list.New()
	for {
		startTime := time.Now()
		// free space in the log for outdated messages
		for head := log.Front(); log.Len() > 0; {
			if head.Value.(time.Time).Add(time.Second).Before(startTime) {
				next := head.Next()
				log.Remove(head)
				head = next
			} else {
				break
			}
		}
		// process messages
		for log.Len() < slidingWindowLogRatePerSecond {
			select {
			case msg := <-in:
				log.PushBack(time.Now())
				_ = msg
			default:
			}
		}
	}
}
