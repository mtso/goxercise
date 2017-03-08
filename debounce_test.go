package exercises

import (
	"testing"
	"time"
)

func increment(i *int) {
	*i++
}

func Test_DebounceSimple(t *testing.T) {
	called := 0
	timer := time.NewTimer(time.Second)
	donetimer := time.NewTimer(time.Second * 3)

	// TODO: write debounce such that any function may be passed in...
	debounced := debounce(func() {
		called += 1
	}, time.Second)

	go func() {
		<-timer.C
		debounced()
		debounced()
	}()

	debounced()
	debounced()

	<-donetimer.C

	timer.Stop()
	donetimer.Stop()
	if called != 2 {
		t.Errorf("expected called to equal 2, but was %v", called)
	}
}
