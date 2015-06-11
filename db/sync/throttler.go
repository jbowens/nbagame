package sync

func throttle(funcs []func(), maxConcurrent int) {
	c := make(chan struct{}, maxConcurrent)
	for _, f := range funcs {
		// Wait until the throttle channel has buffer space
		c <- struct{}{}
		go func(f func()) {
			// Execute the function
			f()

			// Let the next guy know that we just finished
			<-c
		}(f)
	}
}
