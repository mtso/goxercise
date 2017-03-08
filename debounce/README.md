Basic Usage:
```go
called := 0
func call() {
	called++
}
debounced := debounce(call, 100) // You will define the `debounce` function
debounced()
debounced()
debounced()
debounced()

// called should equal 1
```