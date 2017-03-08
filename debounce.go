package exercises

import "time"

func debounce(callback func(arg *int), delay time.Duration) func(arg *int) {
	return callback
}