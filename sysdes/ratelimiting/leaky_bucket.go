package ratelimiting

import "time"

const (
	leakyBucketRatePerSecond = 10
	leakyBucketQueueCapacity = 50
)

func leakyBucket(in <-chan struct{}) {
	bucket := make(chan struct{}, leakyBucketQueueCapacity)
	go func() {
		for {
			select {
			case msg := <-in:
				// put message to the bucket when the size is less than capacity
				// discard message otherwise
				select {
				case bucket <- msg:
				default:
				}
			}
		}
	}()
	for {
		select {
		case <-time.After(time.Second):
			batch := make([]struct{}, 0, leakyBucketRatePerSecond)
		BatchLoop:
			for len(batch) < leakyBucketRatePerSecond {
				select {
				case msg := <-bucket:
					batch = append(batch, msg)
				default:
					break BatchLoop
				}
			}
			// process batch
		}
	}
}
