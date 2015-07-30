/*
Package "serial" provides "encoding/binary"'s reader/writer wrappers with error checking.
*/
package serial

const (
	INVALID_ID_ERROR           = "INVALID_ID_ERROR"
	INCOMPATIBLE_VERSION_ERROR = "INCOMPATIBLE_VERSION_ERROR"
	INCOMPATIBLE_ARCH_ERROR    = "INCOMPATIBLE_ARCH_ERROR"
)

const (
	ARCH_386 byte = iota
	ARCH_AMD64
	ARCH_ARM
	ARCH_UNKNOW byte = 255
)

func convertToArchType(arch string) byte {
	switch arch {
	case "386":
		return ARCH_386
	case "amd64":
		return ARCH_AMD64
	case "arm":
		return ARCH_ARM
	}

	return ARCH_UNKNOW
}
