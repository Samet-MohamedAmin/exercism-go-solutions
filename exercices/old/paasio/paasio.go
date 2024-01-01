package paasio

import "io"

// NewWriteCounter returns new WriteCounter
func NewWriteCounter(w io.Writer) WriteCounter {

	return WriteCounter{
		io.Writer:  func(p []byte) (n int, err error) { return },
		WriteCount: func(n int64, nops int) { return 0, 0 },
	}
}

// NewReadCounter returns new ReadCounter
func NewReadCounter(r io.Reader) ReadCounter {

	return nil
}

// NewReadWriteCounter returns new ReadWriteCounter
func NewReadWriteCounter(w io.Writer) ReadWriteCounter {

	return nil
}
