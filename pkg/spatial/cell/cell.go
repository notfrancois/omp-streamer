package cell

import (
	"sync"

	"github.com/notfrancois/omp-streamer/pkg/types"
	"github.com/notfrancois/omp-streamer/pkg/types/common"
)

// Cell represents a cell in the spatial grid.
type Cell struct {
	mu sync.RWMutex
	ID common.CellID

	Actors          map[int32]*common.Actor
	Areas           map[int32]*common.Area
	Checkpoints     map[int32]*common.Checkpoint
	MapIcons        map[int32]*common.MapIcon
	Objects         map[int32]*common.Object
	Pickups         map[int32]*common.Pickup
	RaceCheckpoints map[int32]*common.RaceCheckpoint
	TextLabels      map[int32]*common.TextLabel

	Bounds     CellBounds
	references int
}

// CellBounds represents the bounds of a cell.
type CellBounds struct {
	MinX float32
	MinY float32
	MaxX float32
	MaxY float32
}

// Default constructor, used to create a new cell with default values.
func Create() *Cell {
	return &Cell{
		ID:         common.CellID{X: -1, Y: -1}, // as default, the cell is not valid until it is set with proper values.
		references: 0,
	}
}

// CreateWithID creates a new cell with a specific ID.
func CreateWithID(cellID common.CellID) *Cell {
	c := Create()
	c.SetID(cellID)
	return c
}

// GetID returns the ID of the cell.
func (c *Cell) GetID() common.CellID {
	var id common.CellID
	c.withRLock(func() {
		id = c.ID
	})
	return id
}

// SetID sets the ID of the cell.
func (c *Cell) SetID(id common.CellID) {
	c.withLock(func() {
		c.ID = id
	})
}

// GetActors returns the actors of the cell.
func (c *Cell) GetActors() map[int32]*common.Actor {
	var actors map[int32]*common.Actor
	c.withRLock(func() {
		actors = c.Actors
	})
	return actors
}

// GetAreas returns the areas of the cell.
func (c *Cell) GetAreas() map[int32]*common.Area {
	var areas map[int32]*common.Area
	c.withRLock(func() {
		areas = c.Areas
	})
	return areas
}

// GetCheckpoints returns the checkpoints of the cell.
func (c *Cell) GetCheckpoints() map[int32]*common.Checkpoint {
	var checkpoints map[int32]*common.Checkpoint
	c.withRLock(func() {
		checkpoints = c.Checkpoints
	})
	return checkpoints
}

// GetMapIcons returns the map icons of the cell.
func (c *Cell) GetMapIcons() map[int32]*common.MapIcon {
	var mapIcons map[int32]*common.MapIcon
	c.withRLock(func() {
		mapIcons = c.MapIcons
	})
	return mapIcons
}

// GetObjects returns the objects of the cell.
func (c *Cell) GetObjects() map[int32]*common.Object {
	var objects map[int32]*common.Object
	c.withRLock(func() {
		objects = c.Objects
	})
	return objects
}

// GetPickups returns the pickups of the cell.
func (c *Cell) GetPickups() map[int32]*common.Pickup {
	var pickups map[int32]*common.Pickup
	c.withRLock(func() {
		pickups = c.Pickups
	})
	return pickups
}

// GetRaceCheckpoints returns the race checkpoints of the cell.
func (c *Cell) GetRaceCheckpoints() map[int32]*common.RaceCheckpoint {
	var raceCheckpoints map[int32]*common.RaceCheckpoint
	c.withRLock(func() {
		raceCheckpoints = c.RaceCheckpoints
	})
	return raceCheckpoints
}

// GetTextLabels returns the text labels of the cell.
func (c *Cell) GetTextLabels() map[int32]*common.TextLabel {
	var textLabels map[int32]*common.TextLabel
	c.withRLock(func() {
		textLabels = c.TextLabels
	})
	return textLabels
}

// withLock executes the function f with the lock of the cell.
// It is used to execute functions that need to be protected by the cell's lock.
// example:
//
//	func (c *Cell) AddArea(area interfaces.Area) {
//		c.withLock(func() {
//			c.Areas[area.GetID()] = area
//		})
//	}
func (c *Cell) withLock(f func()) {
	c.mu.Lock()
	defer c.mu.Unlock()
	f()
}

// withRLock executes the function f with the read lock of the cell.
// It is used to execute functions that need to be protected by the cell's read lock.
// example:
//
//	func (c *Cell) GetID() common.CellID {
//		return c.withRLock(func() common.CellID {
//			return c.ID
//		})
//	}
func (c *Cell) withRLock(f func()) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	f()
}

// ContainsPoint checks if a point is inside the cell.
func (c *Cell) ContainsPoint(point types.Vector3) bool {
	return point.X >= c.Bounds.MinX && point.X < c.Bounds.MaxX &&
		point.Y >= c.Bounds.MinY && point.Y < c.Bounds.MaxY
}

// IsEmpty checks if the cell is empty.
func (c *Cell) IsEmpty() bool {
	var empty bool
	c.withRLock(func() {
		empty = len(c.Actors) == 0 &&
			len(c.Areas) == 0 &&
			len(c.Checkpoints) == 0 &&
			len(c.MapIcons) == 0 &&
			len(c.Objects) == 0
	})
	return empty
}
