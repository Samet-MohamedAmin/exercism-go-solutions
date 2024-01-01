package rotationalcipher

type limit struct {
	start, end byte
}

var limits = []limit{{'a', 'z'}, {'A', 'Z'}}

func (l limit) rotate(b, key byte) byte {
	return (b-l.start+key)%(l.end-l.start+1) + l.start
}

func rotateByte(b, key byte) byte {
	for _, l := range limits {
		if b >= l.start && b <= l.end {
			return l.rotate(b, byte(key))
		}
	}
	return b
}

// RotationalCipher returns encoded string based on given input and key
func RotationalCipher(plain string, key int) string {
	bs := []byte{}
	for _, b := range []byte(plain) {
		bs = append(bs, rotateByte(b, byte(key)))
	}
	return string(bs)
}
