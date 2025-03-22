package streamer

// #include <stdlib.h>
// #include "modules/streamer/src/streamer_wrapper.h"
import "C"
import "unsafe"

// goToC_Vector3 convierte un Vector3 de Go a C
func goToC_Vector3(v Vector3) C.Vector3 {
	return C.Vector3{
		x: C.float(v.X),
		y: C.float(v.Y),
		z: C.float(v.Z),
	}
}

// cToGo_Vector3 convierte un Vector3 de C a Go
func cToGo_Vector3(v C.Vector3) Vector3 {
	return Vector3{
		X: float32(v.x),
		Y: float32(v.y),
		Z: float32(v.z),
	}
}

// goToC_Vector2 convierte un Vector2 de Go a C
func goToC_Vector2(v Vector2) C.Vector2 {
	return C.Vector2{
		x: C.float(v.X),
		y: C.float(v.Y),
	}
}

// cToGo_Vector2 convierte un Vector2 de C a Go
func cToGo_Vector2(v C.Vector2) Vector2 {
	return Vector2{
		X: float32(v.x),
		Y: float32(v.y),
	}
}

// goToC_String convierte una cadena de Go a C
func goToC_String(s string) C.String {
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))

	return C.String{
		buf:    cstr,
		length: C.size_t(len(s)),
	}
}

// Funciones adicionales de conversión para otros tipos...
