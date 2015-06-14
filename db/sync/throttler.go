package sync

import (
	"strings"
	"sync"
)

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

type multierror []error

// Error implements the error interface.
func (e multierror) Error() string {
	var errMessages []string
	for _, err := range e {
		errMessages = append(errMessages, err.Error())
	}
	return strings.Join(errMessages, ", ")
}

type throttler struct {
	wg        sync.WaitGroup
	errors    []error
	pendingCh chan struct{}
	errorsCh  chan error
	doneCh    chan struct{}
}

func newThrottler(maxConcurrent int) *throttler {
	t := &throttler{
		pendingCh: make(chan struct{}, maxConcurrent),
		errorsCh:  make(chan error),
		doneCh:    make(chan struct{}),
	}

	go t.loop()

	return t
}

func (t *throttler) wait() error {
	// Wait for evertyhing to exit.
	t.wg.Wait()
	close(t.pendingCh)

	// Tell the throttler's errors goroutine to exit.
	close(t.errorsCh)
	// Wait for the errors goroutine to exit.
	<-t.doneCh

	// Return a single aggregate error.
	if len(t.errors) > 0 {
		return multierror(t.errors)
	}
	return nil
}

func (t *throttler) loop() {
	for err := range t.errorsCh {
		t.errors = append(t.errors, err)
	}
	close(t.doneCh)
}

func (t *throttler) run(errorableFunc func() error) {
	t.wg.Add(1)
	t.pendingCh <- struct{}{}
	go func() {
		err := errorableFunc()
		if err != nil {
			t.errorsCh <- err
		}

		<-t.pendingCh
		t.wg.Done()
	}()
}
