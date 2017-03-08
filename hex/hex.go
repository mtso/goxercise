package hex

// var hexChars = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
var hexChars = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}

func stringHex(in []byte) (out string) {
	for _, v := range in {
		temp := hexChars[v & 0xF]
		v = v >> 4
		out += hexChars[v & 0xF]
		out += temp
	}
	return
}