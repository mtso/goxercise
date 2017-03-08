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
	debounced := debounce(func(arg *int) {
		*arg += 1
	}, time.Second)

	go func() {
		<-timer.C
		debounced(&called)
		debounced(&called)
	}()

	debounced(&called)
	debounced(&called)

	<-donetimer.C

	timer.Stop()
	donetimer.Stop()
	if called != 2 {
		t.Errorf("expected called to equal 2, but was %v", called)
	}
}
