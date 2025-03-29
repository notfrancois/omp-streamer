package streamer

// #include <stdlib.h>
// #include "include/omp.h"
import "C"
import (
	"fmt"
	"runtime/debug"
	"strings"
)

//export handlePanic
func handlePanic() {
	if r := recover(); r != nil {
		stackTrace := strings.TrimSuffix(string(debug.Stack()), "\n")

		Log(LogLevelError, "%s", fmt.Sprint(r))
		Log(LogLevelError, "%s", stackTrace)
	}
}

//export OnGmInit
func OnGmInit() C.bool {
	defer handlePanic()

	C.loadSdk()

	return true
}

//export OnGmExit
func OnGmExit() C.bool {
	defer handlePanic()

	C.unloadSdk()

	return true
}
