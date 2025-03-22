package ratelimiting

import "time"

const fixedWindowCounterRatePerSecond = 50

func fixedWindowCounter(in <-chan struct{}) {
	for {
		select {
		case startTime := <-time.After(time.Second):
			endTime := startTime.Add(time.Second)
			for i := 0; i < fixedWindowCounterRatePerSecond && time.Now().Before(endTime); i++ {
				select {
				case msg := <-in:
					// process message
					_ = msg
				default:
				}
			}
		}
	}
}
