package serial

import (
	"encoding/binary"
	"errors"
	"io"
	"runtime"
)

// Reader is a wrapper of "encoding/binary"'s Read.
// Even if an error is caused by (*Reader).Read,
// it doesn't return any value.
// The error can be obtained by calling (*Reader).Error.
type Reader struct {
	id      string
	version byte
	reader  io.Reader
	err     error
}

// Create a wrapper of "encoding/binary"'s Read.
// "id" is the identifier of byte-sequence representation to be read.
// "version" is used to checking the compatibility of byte-sequence representation.
// If "id" or "version is not same as read ones,
// the reader stores an error value into itself.
func NewReader(id string, version byte, reader io.Reader) *Reader {
	r := &Reader{
		id:      id,
		version: version,
		reader:  reader,
	}

	return r
}

// Read the identifier of type by using (*Reader).Read
// and check whether it equals to r.id.
// If it is not same as r.id,
// the reader stores an error value into itself.
func (r *Reader) ReadId() {
	id := []byte(r.id)
	size := len(id)

	var b byte

	for i := 0; i < size; i++ {
		r.Read(&b)

		if r.Error() != nil {
			return
		}

		if id[i] != b {
			r.err = errors.New(INVALID_ID_ERROR)
			return
		}
	}
}

// Read the version of byte-sequence representation by using (*Reader).Read
// and check whether it equals to r.version with element-wise manner.
// If it is not same as r.version,
// the reader stores an error value into itself.
func (r *Reader) ReadVersion() {
	var version byte

	r.Read(&version)

	if r.Error() != nil {
		return
	}

	if r.version != version {
		r.err = errors.New(INCOMPATIBLE_VERSION_ERROR)
		return
	}
}

// Read the architecture  by using (*Reader).Read
// and check whether it equals to the current architecture.
// If it is not same as r.version,
// the reader stores an error value into itself.
func (r *Reader) ReadArch() {
	var arch byte

	r.Read(&arch)

	if r.Error() != nil {
		return
	}

	if convertToArchType(runtime.GOARCH) != arch {
		r.err = errors.New(INCOMPATIBLE_ARCH_ERROR)
		return
	}
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
