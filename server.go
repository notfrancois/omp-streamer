package streamer

// #include <stdlib.h>
// #include "include/server.h"
import "C"
import (
	"fmt"
	"unsafe"
)

type LogLevel int

const (
	LogLevelDebug LogLevel = iota
	LogLevelMessage
	LogLevelWarning
	LogLevelError
)

func Log(level LogLevel, format string, a ...any) {
	msg := fmt.Sprintf(format, a...)

	cMsg := C.CString(msg)
	defer C.free(unsafe.Pointer(cMsg))

	C.server_logLnU8(C.int(level), cMsg)
}
