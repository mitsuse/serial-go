package serial

import (
	"encoding/binary"
	"io"
)

// Reader is a wrapper of "encoding/binary"'s Read.
// Even if an error is caused by (*Reader).Read,
// it doesn't return any value.
// The error can be obtained by calling (*Reader).Error.
type Reader struct {
	reader io.Reader
	err    error
}

// Create a wrapper of "encoding/binary"'s Read.
func NewReader(reader io.Reader) *Reader {
	r := &Reader{
		reader: reader,
	}

	return r
}

// Read a value by using "encoding/binary"'s Read
// while (*Reader).Error() is nil.
// If (*Reader).Error() is not nil, do nothing.
func (r *Reader) Read(value interface{}) {
	if r.err != nil {
		return
	}

	r.err = binary.Read(r.reader, binary.LittleEndian, value)
}

// Return the first error caused by (*Reader).Read.
func (r *Reader) Error() error {
	return r.err
}
