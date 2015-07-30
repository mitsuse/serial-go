package serial

import (
	"encoding/binary"
	"io"
	"runtime"
)

// Writer is a wrapper of "encoding/binary"'s Write.
// Even if an error is caused by (*Writer).Write,
// it doesn't return any value.
// The error can be obtained by calling (*Writer).Error.
type Writer struct {
	id      string
	version byte
	writer  io.Writer
	err     error
}

// Create a wrapper of "encoding/binary"'s Write.
func NewWriter(id string, version byte, writer io.Writer) *Writer {
	w := &Writer{
		id:      id,
		version: version,
		writer:  writer,
	}

	return w
}

// Write the identifier of type by using (*Writer).Write.
func (w *Writer) WriteId() {
	for _, b := range []byte(w.id) {
		w.Write(b)
	}
}

// Write the version of byte-sequence representation by using (*Writer).Write.
func (w *Writer) WriteVersion() {
	w.Write(w.version)
}

// Write the current archtechture by using (*Writer).Write.
func (w *Writer) WriteArch() {
	w.Write(convertToArchType(runtime.GOARCH))
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
