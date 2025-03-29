package streamer

// #cgo CFLAGS: -Iinclude
// #cgo LDFLAGS: -L. -lstreamer
// #include "streamer.h"
import "C"

func CreateDynamicObject(modelid int, x float32, y float32, z float32, rx float32, ry float32, rz float32, worldid int, interiorid int, playerid int) int {
	return int(C.streamer_createDynamicObject(C.int(modelid), C.float(x), C.float(y), C.float(z), C.float(rx), C.float(ry), C.float(rz), C.int(worldid), C.int(interiorid), C.int(playerid)))
}

func DestroyDynamicObject(objectid int) bool {
	// Validate object ID before making C call
	if objectid <= 0 {
		return false
	}
	C.streamer_destroyDynamicObject(C.int(objectid))
	return true
}
