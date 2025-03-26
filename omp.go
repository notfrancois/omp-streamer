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

//export onGameModeInit
func onGameModeInit() C.bool {
	defer handlePanic()

	C.loadComponent()

	return true
}

//export onGameModeExit
func onGameModeExit() C.bool {
	defer handlePanic()

	C.unloadComponent()

	return true
}
