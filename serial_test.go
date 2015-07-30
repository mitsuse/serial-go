package serial

import (
	"bytes"
	"errors"
	"testing"
)

func TestWriterWriteSucceedsWithAcceptableValueForBinaryWrite(t *testing.T) {
	var id string = "test"
	var version byte = 0
	var test int64 = 1

	buffer := bytes.NewBuffer([]byte{})
	writer := NewWriter(id, version, buffer)

	writer.Write(&test)

	if err := writer.Error(); err != nil {
		t.Fatalf("An expected error is caused with an acceptable value: %s", err)
	}

	expectation := []byte{1, 0, 0, 0, 0, 0, 0, 0}
	result := buffer.Bytes()

	for i, b := range expectation {
		if r := result[i]; r != b {
			t.Fatalf("An element at the index %d should be %d, but is %d.", i, b, r)
		}
	}
}

func TestWriterWriteFailsWithUnacceptableValueForBinaryWrite(t *testing.T) {
	var id string = "test"
	var version byte = 0
	var test int = 1

	buffer := bytes.NewBuffer([]byte{})
	writer := NewWriter(id, version, buffer)

	writer.Write(&test)

	if err := writer.Error(); err == nil {
		t.Fatal("An error should be caused with an unacceptable value.")
	}
}

func TestWriterWriteIgnoresValueWithError(t *testing.T) {
	var id string = "test"
	var version byte = 0
	var test int64 = 1

	buffer := bytes.NewBuffer([]byte{})
	writer := &Writer{
		id:      id,
		version: version,
		writer:  buffer,
		err:     errors.New("dummy"),
	}

	writer.Write(&test)

	if len(buffer.Bytes()) == 0 {
		return
	}

	t.Fatal("A value should be ignored if the writer has an error.")
}

func TestReaderReadSucceedsWithAcceptableValueForBinaryRead(t *testing.T) {
	var id string = "test"
	var version byte = 0
	var test int64 = 0
	var expectation int64 = 1

	buffer := bytes.NewReader([]byte{1, 0, 0, 0, 0, 0, 0, 0})
	reader := NewReader(id, version, buffer)

	reader.Read(&test)

	if err := reader.Error(); err != nil {
		t.Fatalf("An expected error is caused with an acceptable value: %s", err)
	}

	if test != expectation {
		t.Fatalf("%d should be read, but %d is read.", expectation, test)
	}
}

func TestReaderReadFailsWithUnacceptableValueForBinaryWrite(t *testing.T) {
	var id string = "test"
	var version byte = 0
	var test int64 = 0

	buffer := bytes.NewReader([]byte{1, 0, 0, 0})
	reader := NewReader(id, version, buffer)

	reader.Read(&test)

	if err := reader.Error(); err == nil {
		t.Fatal("An error should be caused with an unacceptable value.")
	}
}

func TestReaderReadIgnoresValueWithError(t *testing.T) {
	var id string = "test"
	var version byte = 0
	var test int64 = 0

	buffer := bytes.NewReader([]byte{1, 0, 0, 0, 0, 0, 0, 0})
	reader := &Reader{
		id:      id,
		version: version,
		reader:  buffer,
		err:     errors.New("dummy"),
	}

	reader.Read(&test)

	if test == 0 {
		return
	}

	t.Fatal("A value should be ignored if the reader has an error.")
}
