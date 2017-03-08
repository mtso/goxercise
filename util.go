package exercises

import "time"

func setTimeout(callback func(), delay time.Duration) func() {
	timer := time.NewTimer(delay)
	isStopped := false

	go func() {
		<-timer.C
		if isStopped {
			return
		}
		callback()
		timer.Stop()
	}()

	// the stopping func
	return func() {
		timer.Stop()
	}
}
