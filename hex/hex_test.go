package hex

import (
	"testing"
	"fmt"
)

func TestHex(t *testing.T) {
	b := []byte{65, 66, 67} // ABC
	expected := fmt.Sprintf("%x", b)
	result := stringHex(b)
	if result != expected {
		t.Errorf("Expected formatHex(%q) == %q, but got %q", b, expected, result)
	}
}