package game

// #include "../../include/player.h"
import "C"
import (
	"unsafe"

	"github.com/notfrancois/omp-streamer/pkg/types"
)

// Player is a wrapper around a open.mp player.
type Player struct {
	handle unsafe.Pointer
}

// Position returns the player's position.
func (p *Player) Position() types.Vector3 {
	pos := C.player_getPosition(p.handle)

	return types.Vector3{
		X: float32(pos.x),
		Y: float32(pos.y),
		Z: float32(pos.z),
	}
}
