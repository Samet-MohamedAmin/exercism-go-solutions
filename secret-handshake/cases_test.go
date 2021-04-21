package secret

// Source: exercism/problem-specifications
// Commit: a9e4df8 secret-handshake: Apply new "input" policy
// Problem Specifications Version: 1.2.0

type secretHandshakeTest struct {
	code uint
	h    []string
}

var tests = []secretHandshakeTest{
	{1, []string{"wink"}},                  // 1
	{2, []string{"double blink"}},          // 10
	{4, []string{"close your eyes"}},       // 100
	{8, []string{"jump"}},                  // 1000
	{3, []string{"wink", "double blink"}},  // 11
	{19, []string{"double blink", "wink"}}, // 10011
	{24, []string{"jump"}},                 // 11000
	{16, []string{}},                       // 10000
	{15, []string{"wink", "double blink", "close your eyes", "jump"}}, // 1111
	{31, []string{"jump", "close your eyes", "double blink", "wink"}}, // 11111
	{0, []string{}},
}
