package serial

import (
	"bytes"
	"errors"
	"runtime"
	"testing"
)

func TestConvertToArchTypeSucceedsFor386(t *testing.T) {
	arch := "386"

	if convertToArchType(arch) != ARCH_386 {
		t.Fatal("386 should be converted to ARCH_386")
	}
}

func TestConvertToArchTypeSucceedsForAmd64(t *testing.T) {
	arch := "amd64"

	if convertToArchType(arch) != ARCH_AMD64 {
		t.Fatal("amd64 should be converted to ARCH_AMD64")
	}
}

func TestConvertToArchTypeSucceedsForArm(t *testing.T) {
	arch := "arm"

	if convertToArchType(arch) != ARCH_ARM {
		t.Fatal("arm should be converted to ARCH_ARM")
	}
}
func TestConvertToArchTypeSucceedsForUnknown(t *testing.T) {
	arch := "---"

	if convertToArchType(arch) != ARCH_UNKNOW {
		t.Fatal("--- should be converted to ARCH_UNKNOW")
	}
}

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

	if len(expectation) != len(result) {
		t.Fatalf(
			"The size of byte-sequence representation should be %d, but is %d.",
			len(expectation),
			len(result),
		)
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

func TestWriteIdUsesWrite(t *testing.T) {
	var id string = "test"
	var version byte = 0

	buffer := bytes.NewBuffer([]byte{})
	writer := NewWriter(id, version, buffer)

	writer.WriteId()

	if err := writer.Error(); err != nil {
		t.Fatalf("An expected error is caused with an acceptable value: %s", err)
	}
}

func TestWriteVersionUsesWrite(t *testing.T) {
	var id string = "test"
	var version byte = 0

	buffer := bytes.NewBuffer([]byte{})
	writer := NewWriter(id, version, buffer)

	writer.WriteVersion()

	if err := writer.Error(); err != nil {
		t.Fatalf("An expected error is caused with an acceptable value: %s", err)
	}
}

func TestWriteArchUsesWrite(t *testing.T) {
	var id string = "test"
	var version byte = 0

	buffer := bytes.NewBuffer([]byte{})
	writer := NewWriter(id, version, buffer)

	writer.WriteArch()

	if err := writer.Error(); err != nil {
		t.Fatalf("An expected error is caused with an acceptable value: %s", err)
	}
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

func TestReadIdSucceedsForValidId(t *testing.T) {
	var id string = "test"
	var version byte = 1

	buffer := bytes.NewBuffer([]byte(id))
	reader := NewReader(id, version, buffer)

	reader.ReadId()

	if err := reader.Error(); err != nil {
		t.Fatalf("An expected error occured on reading a valid id: %s", err)
	}
}

func TestReadIdFailsForIncompatibleVersion(t *testing.T) {
	var id string = "test"
	var version byte = 1

	buffer := bytes.NewBuffer([]byte("tess"))
	reader := NewReader(id, version, buffer)

	reader.ReadId()

	if err := reader.Error(); err == nil {
		t.Fatal("An error should occur for an invalid id.")
	}
}

func TestReadIdFailsWithFailureOfBinaryRead(t *testing.T) {
	var id string = "test"
	var version byte = 0

	buffer := bytes.NewBuffer([]byte{})
	reader := NewReader(id, version, buffer)

	reader.ReadId()

	if err := reader.Error(); err == nil {
		t.Fatal("An error should be caused by \"binary\".Read.")
	}
}

func TestReadVersionFailsWithFailureOfBinaryRead(t *testing.T) {
	var id string = "test"
	var version byte = 0

	buffer := bytes.NewBuffer([]byte{})
	reader := NewReader(id, version, buffer)

	reader.ReadVersion()

	if err := reader.Error(); err == nil {
		t.Fatal("An error should be caused by \"binary\".Read.")
	}
}

func TestReadVersionSucceedsForCompatibleVersion(t *testing.T) {
	var id string = "test"
	var version byte = 1

	buffer := bytes.NewBuffer([]byte{version})
	reader := NewReader(id, version, buffer)

	reader.ReadVersion()

	if err := reader.Error(); err != nil {
		t.Fatalf("An expected error occured on reading compatible version: %s", err)
	}
}

func TestReadVersionFailsForIncompatibleVersion(t *testing.T) {
	var id string = "test"
	var version byte = 1

	buffer := bytes.NewBuffer([]byte{0})
	reader := NewReader(id, version, buffer)

	reader.ReadVersion()

	if err := reader.Error(); err == nil {
		t.Fatal("An error should occur for incompatible version.")
	}
}

func TestReadArchFailsWithFailureOfBinaryRead(t *testing.T) {
	var id string = "test"
	var version byte = 0

	buffer := bytes.NewBuffer([]byte{})
	reader := NewReader(id, version, buffer)

	reader.ReadArch()

	if err := reader.Error(); err == nil {
		t.Fatal("An error should be caused by \"binary\".Read.")
	}
}

func TestReadArchSucceedsForCompatibleArch(t *testing.T) {
	var id string = "test"
	var version byte = 1

	buffer := bytes.NewBuffer([]byte{convertToArchType(runtime.GOARCH)})
	reader := NewReader(id, version, buffer)

	reader.ReadArch()

	if err := reader.Error(); err != nil {
		t.Fatalf("An expected error occured on reading compatible version: %s", err)
	}
}

func TestReadArchFailsForIncompatibleArch(t *testing.T) {
	var id string = "test"
	var version byte = 1

	buffer := bytes.NewBuffer([]byte{254})
	reader := NewReader(id, version, buffer)

	reader.ReadArch()

	if err := reader.Error(); err == nil {
		t.Fatal("An error should occur for incompatible version.")
	}
}
