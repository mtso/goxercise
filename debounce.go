package exercises

import "time"

func debounce(callback func(), delay time.Duration) func() {
	return callback
}
