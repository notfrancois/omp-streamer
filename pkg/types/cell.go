package types

import (
	"fmt"

	"github.com/notfrancois/omp-streamer/pkg/types"
)

type Cell interface {
	GetID() CellID
	SetID(id CellID)
	GetActors() map[int32]*Actor
	GetAreas() map[int32]*Area
	GetCheckpoints() map[int32]*Checkpoint
	GetMapIcons() map[int32]*MapIcon
	GetObjects() map[int32]*Object
	GetPickups() map[int32]*Pickup
	GetRaceCheckpoints() map[int32]*RaceCheckpoint
	GetTextLabels() map[int32]*TextLabel

	// spatial
	ContainsPoint(point types.Vector3) bool
	IsEmpty() bool
}

type CellID struct {
	X, Y int32
}

type Position struct {
	X, Y, Z float64
}

// Hash implements a hash for CellID
func (c CellID) Hash() string {
	return fmt.Sprintf("%d:%d", c.X, c.Y)
}
