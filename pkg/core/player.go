package core

import (
	"time"

	"github.com/notfrancois/omp-streamer/internal/spatial"
	"github.com/notfrancois/omp-streamer/pkg/collections"
	"github.com/notfrancois/omp-streamer/pkg/core/bitsets"
	"github.com/notfrancois/omp-streamer/pkg/player"
	"github.com/notfrancois/omp-streamer/pkg/trust"
	"github.com/notfrancois/omp-streamer/pkg/types"
	"github.com/notfrancois/omp-streamer/pkg/types/common"
)

// Player is a struct that represents a player in the game.
// It contains all the information about the player and the items that are being processed by the player.
type Player struct {
	ID                     int32
	ActiveRaceCheckpointID int32
	ActiveCheckpointID     int32

	ChunkTickCount [types.StreamerMaxTypes]int32
	ChunkTickRate  [types.StreamerMaxTypes]int32

	CurrentVisibleObjects    int32
	CurrentVisibleTextLabels int32

	DelayedCheckpointID     int32
	DelayedRaceCheckpointID int32

	DelayedUpdate       bool
	DelayedUpdateFreeze int32
	DelayedUpdateTime   time.Time
	DelayedUpdateType   int32

	InteriorID int32
	WorldID    int32

	MaxVisibleMapIcons   int32
	MaxVisibleObjects    int32
	MaxVisibleTextLabels int32

	Position          types.Vector3
	RadiusMultipliers [types.StreamerMaxTypes]float32
	References        int32
	RequestingClass   bool

	TickCount int32
	TickRate  int32

	UpdateUsingCameraPosition bool
	UpdateWhenIdle            bool

	VisibleCell             *spatial.Cell
	VisibleCheckpointID     int32
	VisibleRaceCheckpointID int32

	EnabledItems     bitsets.ItemBitSet
	ProcessingChunks bitsets.ItemBitSet

	DiscoveredMapIcons   collections.Bimap[int32, *common.MapIcon]
	DiscoveredObjects    collections.Bimap[int32, *common.Object]
	DiscoveredTextLabels collections.Bimap[int32, *common.TextLabel]

	ExistingMapIcons   collections.Bimap[int32, *common.MapIcon]
	ExistingObjects    collections.Bimap[int32, *common.Object]
	ExistingTextLabels collections.Bimap[int32, *common.TextLabel]

	InternalItems     *player.InternalItems
	MapIconIdentifier *trust.Identifier
}

// NewPlayer creates a new player using the given id in the Streamer system.
func NewPlayer() *Player {
	return &Player{
		ID:                int32(GlobalCore.GameSdk().Player.ID()),
		InternalItems:     player.NewInternalItems(),
		MapIconIdentifier: trust.NewIdentifier(),
		VisibleCell:       nil, // TODO: Initialize the cell
	}
}

func (p *Player) Initialize() {
	p.ActiveRaceCheckpointID = 0
	p.ActiveCheckpointID = 0

	p.ChunkTickCount[types.ItemTypeObject] = 0
	p.ChunkTickCount[types.ItemTypeTextLabel] = 0
	p.ChunkTickCount[types.ItemTypeMapIcon] = 0

	p.ChunkTickRate[types.ItemTypeObject] = 1
	p.ChunkTickRate[types.ItemTypeTextLabel] = 1
	p.ChunkTickRate[types.ItemTypeMapIcon] = 1

	p.CurrentVisibleObjects = GlobalCore.GetData().GetGlobalMaxVisibleItems(types.StreamerTypeObject)
	p.CurrentVisibleTextLabels = GlobalCore.GetData().GetGlobalMaxVisibleItems(types.StreamerType3DTextLabel)

	p.DelayedCheckpointID = 0
	p.DelayedRaceCheckpointID = 0

	p.DelayedUpdate = false
	p.DelayedUpdateType = 0

	if !GlobalCore.GameSdk().Player.IsNPC() {
		p.EnabledItems.SetAll()
	}

	p.InteriorID = 0
	p.WorldID = 0

	p.MaxVisibleMapIcons = GlobalCore.GetData().GetGlobalMaxVisibleItems(types.StreamerTypeMapIcon)
	p.MaxVisibleObjects = GlobalCore.GetData().GetGlobalMaxVisibleItems(types.StreamerTypeObject)
	p.MaxVisibleTextLabels = GlobalCore.GetData().GetGlobalMaxVisibleItems(types.StreamerType3DTextLabel)

	p.Position = types.Vector3{
		X: 0,
		Y: 0,
		Z: 0,
	}

	for t := types.StreamerType(0); t < types.StreamerMaxTypes; t++ {
		p.RadiusMultipliers[t] = GlobalCore.GetData().GetGlobalRadiusMultiplier(t)
	}

	p.RequestingClass = false
	p.TickCount = 0
	p.TickRate = 50
	p.UpdateUsingCameraPosition = false
	p.UpdateWhenIdle = false

	p.VisibleCheckpointID = 0
	p.VisibleRaceCheckpointID = 0
}
