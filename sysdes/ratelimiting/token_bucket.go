package ratelimiting

import "time"

const (
	tokenBucketCapacity              = 50
	tokenBucketRefillNumber          = 10
	tokenBucketRefillIntervalSeconds = 10
)

func tokenBucket(in <-chan struct{}) {
	bucket := make(chan struct{}, tokenBucketCapacity)
	go func() {
		for {
			// refill the bucket with X messages every Y seconds
			// can be changed with atomic counter
			select {
			case <-time.After(time.Second * tokenBucketRefillIntervalSeconds):
			BucketRefillLoop:
				for i := 0; i < tokenBucketRefillNumber; i++ {
					select {
					case bucket <- struct{}{}:
					default:
						break BucketRefillLoop
					}
				}
			}
		}
	}()
	for {
		// get a token and process one message corresponding to this token
		select {
		case <-bucket:
			select {
			case msg := <-in:
				// process message
				_ = msg
			}
		}
	}
}
