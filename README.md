# Serial

[![License](https://img.shields.io/badge/license-MIT-yellowgreen.svg?style=flat-square)][license]
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)][godoc]
[![Version](https://img.shields.io/github/tag/mitsuse/serial-go.svg?style=flat-square)][release]
[![Wercker](http://img.shields.io/wercker/ci/55b8de7c58d352a50100488b.svg?style=flat-square)][wercker]
[![Coverage](https://img.shields.io/codecov/c/github/mitsuse/serial-go.svg?style=flat-square)][coverage]

[license]: LICENSE.txt
[godoc]: http://godoc.org/github.com/mitsuse/serial-go
[release]: https://github.com/mitsuse/serial-go/releases
[wercker]: https://app.wercker.com/project/bykey/cddb9d0e91001805e9ed62c37157b234
[coverage]: https://codecov.io/github/mitsuse/serial-go

A tiny library to make it easy to serialize/deserialize with error checking.

[golang]: http://golang.org/


## Installation

For installation, execute the following command:

```
$ go get github.com/mitsuse/serial-go
```


## Example


### Writer

[`*Writer`](http://godoc.org/github.com/mitsuse/serial-go#Writer) is used to implement serialization.

```go
import (
    "io"

    "github.com/mitsuse/serial-go"
)

const (
    id = "github.com/mitsuse/cat" // Arbitrary string is available for "id".
    version = 0
)

type Cat {
    name string
}

func (c *Cat) Serialize(writer io.Writer) error {
    w := serial.NewWriter(id, version, writer)

    // Write the identifier (name) to check the type of byte-sequence on deserialization.
    w.WriteId()

    // Write the version of the byte-sequence representation for compatibility.
    // Accept the version written in byte-sequence
    // which equals to the version of an implementation used on deserialization.
    w.WriteVersion()

    // Write an arbitary value such as field of struct.
    w.Write(c.name)

    // Return the error caused during serialization.
    // If an error occrurs,
    // any operations for serialization are ignored.
    return w.Error()
}
```


### Reader

[`*Reader`](http://godoc.org/github.com/mitsuse/serial-go#Reader) is used to implement deserialization.

```go
import (
    "io"

    "github.com/mitsuse/serial-go"
)

const (
    id = "github.com/mitsuse/cat" // Arbitrary string is available for "id".
    version = 0
)

type Cat {
    name string
}

func Deserialize(reader io.Reader) (*Cat, error) {
    c := &Cat{}

    r := serial.NewReader(id, version, reader)

    // Read the identifier (name).
    // If the identifier doesn't equal to the current implementation's one,
    // an error value is stored into the reader.
    r.ReadId()

    // Read the version of the byte-sequence representation.
    // If the version doesn't equal to the current implementation's one,
    // an error value is stored into the reader.
    r.ReadVersion()

    // Write an arbitary value such as field of struct.
    r.Read(&c.name)

    // Return the error caused during deserialization.
    // If an error occrurs,
    // any operations for deserialization are ignored.
    return c, r.Error()
}
```


## License

Please read [LICENSE.txt](LICENSE.txt).
