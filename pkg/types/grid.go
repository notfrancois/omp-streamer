package types

import (
	"github.com/notfrancois/omp-streamer/pkg/types"
)

type Grid interface {
	AddActor(actor *Actor)
	AddArea(area *Area)
	AddCheckpoint(checkpoint *Checkpoint)
	AddMapIcon(mapIcon *MapIcon)
	AddObject(object *Object)
	AddPickup(pickup *Pickup)
	AddRaceCheckpoint(raceCheckpoint *RaceCheckpoint)
	AddTextLabel(textLabel *TextLabel)

	GetCellSize() float32
	GetCellDistance() float32

	SetCellSize(size float32)
	SetCellDistance(distance float32)

	RebuildGrid()

	RemoveActor(actor *Actor, reassign bool)
	RemoveArea(area *Area, reassign bool)
	RemoveCheckpoint(checkpoint *Checkpoint, reassign bool)
	RemoveMapIcon(mapIcon *MapIcon, reassign bool)
	RemoveObject(object *Object, reassign bool)
	RemovePickup(pickup *Pickup, reassign bool)
	RemoveRaceCheckpoint(raceCheckpoint *RaceCheckpoint, reassign bool)
	RemoveTextLabel(textLabel *TextLabel, reassign bool)

	CalculateTranslationMatrix()
	EraseCellIfEmpty(cell *Cell)

	GetCellID(position types.Vector2) CellID
	ProcessDiscoveredCellsForPlayer(playerid int32, playerCells []CellID, discoveredCells map[CellID]struct{})
}
