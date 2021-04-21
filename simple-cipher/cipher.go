package cipher

// Cipher contains Encode and Decode methods
type Cipher interface {
	Encode(string) string
	Decode(string) string
}
