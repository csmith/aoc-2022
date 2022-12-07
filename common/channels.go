package common

// Last returns the last value received from the channel before it was closed
func Last(channel <-chan int) (res int) {
	for {
		o, more := <-channel
		if more {
			res = o
		} else {
			return
		}
	}
}
