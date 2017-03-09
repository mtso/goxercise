package util

func AssertEqualBytes(actual, expected []byte) (isEqual bool) {
	if len(actual) != len(expected) {
		isEqual = false
		return
	}
	isEqual = true
	for i, v := range expected {
		if v != actual[i] {
			isEqual = false
		}
	}
	return
}