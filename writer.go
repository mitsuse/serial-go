package serial

import (
	"encoding/binary"
	"io"
)

// Writer is a wrapper of "encoding/binary"'s Write.
// Even if an error is caused by (*Writer).Write,
// it doesn't return any value.
// The error can be obtained by calling (*Writer).Error.
type Writer struct {
	writer io.Writer
	err    error
}

// Create a wrapper of "encoding/binary"'s Write.
func NewWriter(writer io.Writer) *Writer {
	w := &Writer{
		writer: writer,
	}

	return w
}

// Write a value by using "encoding/binary"'s Write
// while (*Writer).Error() is nil.
// If (*Writer).Error() is not nil, do nothing.
func (w *Writer) Write(value interface{}) {
	if w.err != nil {
		return
	}

	w.err = binary.Write(w.writer, binary.LittleEndian, value)
}

// Return the first error caused by (*Writer).Write.
func (w *Writer) Error() error {
	return w.err
}
