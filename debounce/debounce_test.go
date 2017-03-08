package debounce

import (
	"testing"
	"time"
)

func Test_DebounceSimple(t *testing.T) {
	called := 0
	timer := time.NewTimer(time.Second)
	donetimer := time.NewTimer(time.Second * 3)

	// TODO: write debounce such that any function may be passed in...
	// including ones with arguments
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

func Test_Debounce(t *testing.T) {
	called := 0
	debounced := debounce(func() {
		called += 1
	}, time.Millisecond*50)

	go func() {
		debounced()
		debounced()
	}()

	done := make(chan int)

	setTimeout(func() {
		if called > 0 {
			t.Errorf("expected called to == 0, but got %v", called)
		}
	}, time.Millisecond*10)

	setTimeout(func() {
		if called > 1 {
			t.Errorf("expected called to == 1, but got %v", called)
		}
		done <- 0
	}, time.Millisecond*100)

	<-done
}
