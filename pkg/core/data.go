package core

import (
	"math"

	"github.com/notfrancois/omp-streamer/pkg/constant"
	"github.com/notfrancois/omp-streamer/pkg/types"
	"github.com/notfrancois/omp-streamer/pkg/types/common"
)

type Data struct {
	// Slices
	DestroyedActors []int32
	TypePriority    []types.StreamerPriorityType

	// Maps
	DiscoveredActors  map[PairKey]*common.Actor
	DiscoveredObjects map[PairKey]*common.Object

	InternalActors  map[PairKey]common.Actor
	InternalObjects map[PairKey]common.Object

	Actors          map[int32]*common.Actor
	Areas           map[int32]*common.Area
	Checkpoints     map[int32]*common.Checkpoint
	MapIcons        map[int32]*common.MapIcon
	Pickups         map[int32]*common.Pickup
	TextLabels      map[int32]*common.TextLabel
	Objects         map[int32]*common.Object
	RaceCheckpoints map[int32]common.RaceCheckpoint

	Players map[int32]struct{} // @todo: replace with Player struct.

	// privates.
	globalChunkTickRate     map[types.StreamerType]int32
	globalMaxItems          map[types.StreamerType]int32
	globalMaxVisibleItems   map[types.StreamerType]int32
	globalRadiusMultipliers map[types.StreamerType]float32

	errorCallbackEnabled bool
}

// NewData creates a new Data instance.
func NewData() *Data {
	d := &Data{
		DestroyedActors: make([]int32, 0),
		TypePriority:    make([]types.StreamerPriorityType, types.StreamerMaxTypes),

		DiscoveredActors:  make(map[PairKey]*common.Actor),
		DiscoveredObjects: make(map[PairKey]*common.Object),

		InternalActors:  make(map[PairKey]common.Actor),
		InternalObjects: make(map[PairKey]common.Object),

		Actors:          make(map[int32]*common.Actor),
		Areas:           make(map[int32]*common.Area),
		Checkpoints:     make(map[int32]*common.Checkpoint),
		MapIcons:        make(map[int32]*common.MapIcon),
		Pickups:         make(map[int32]*common.Pickup),
		TextLabels:      make(map[int32]*common.TextLabel),
		Objects:         make(map[int32]*common.Object),
		RaceCheckpoints: make(map[int32]common.RaceCheckpoint),

		Players: make(map[int32]struct{}),

		globalChunkTickRate:     make(map[types.StreamerType]int32),
		globalMaxItems:          make(map[types.StreamerType]int32),
		globalMaxVisibleItems:   make(map[types.StreamerType]int32),
		globalRadiusMultipliers: make(map[types.StreamerType]float32),

		errorCallbackEnabled: false,
	}

	d.Initialize()

	return d
}

// Initialize initializes the data.
func (d *Data) Initialize() {
	d.globalChunkTickRate[types.StreamerTypeObject] = 1
	d.globalChunkTickRate[types.StreamerTypeMapIcon] = 1
	d.globalChunkTickRate[types.StreamerType3DTextLabel] = 1

	d.globalMaxItems[types.StreamerTypeObject] = math.MaxInt32
	d.globalMaxItems[types.StreamerTypePickup] = math.MaxInt32
	d.globalMaxItems[types.StreamerTypeCP] = math.MaxInt32
	d.globalMaxItems[types.StreamerTypeRaceCP] = math.MaxInt32
	d.globalMaxItems[types.StreamerTypeMapIcon] = math.MaxInt32
	d.globalMaxItems[types.StreamerType3DTextLabel] = math.MaxInt32
	d.globalMaxItems[types.StreamerTypeArea] = math.MaxInt32
	d.globalMaxItems[types.StreamerTypeActor] = math.MaxInt32

	d.globalMaxVisibleItems[types.StreamerTypeObject] = 500
	d.globalMaxVisibleItems[types.StreamerTypeMapIcon] = 100
	d.globalMaxVisibleItems[types.StreamerType3DTextLabel] = 1024
	d.globalMaxVisibleItems[types.StreamerTypePickup] = 4096
	d.globalMaxVisibleItems[types.StreamerTypeActor] = 1000

	d.globalRadiusMultipliers[types.StreamerTypeObject] = 1.0
	d.globalRadiusMultipliers[types.StreamerTypePickup] = 1.0
	d.globalRadiusMultipliers[types.StreamerTypeCP] = 1.0
	d.globalRadiusMultipliers[types.StreamerTypeRaceCP] = 1.0
	d.globalRadiusMultipliers[types.StreamerTypeMapIcon] = 1.0
	d.globalRadiusMultipliers[types.StreamerType3DTextLabel] = 1.0
	d.globalRadiusMultipliers[types.StreamerTypeArea] = 1.0
	d.globalRadiusMultipliers[types.StreamerTypeActor] = 1.0

	d.TypePriority[constant.StreamerTypeArea] = 0
	d.TypePriority[constant.StreamerTypeObject] = 1
	d.TypePriority[constant.StreamerTypeCP] = 2
	d.TypePriority[constant.StreamerTypeRaceCP] = 3
	d.TypePriority[constant.StreamerTypeMapIcon] = 4
	d.TypePriority[constant.StreamerType3DTextLabel] = 5
	d.TypePriority[constant.StreamerTypePickup] = 6
	d.TypePriority[constant.StreamerTypeActor] = 7
}

// GetGlobalChunkTickRate returns the global chunk tick rate for a given type.
func (d *Data) GetGlobalChunkTickRate(t types.StreamerType) int32 {
	if t < 0 || t >= types.StreamerMaxTypes {
		return 0
	}
	return d.globalChunkTickRate[t]
}

// SetGlobalChunkTickRate sets the global chunk tick rate for a given type.
func (d *Data) SetGlobalChunkTickRate(t types.StreamerType, rate int32) {
	if t < 0 || t >= types.StreamerMaxTypes {
		return
	}
	d.globalChunkTickRate[t] = rate
}

// GetGlobalMaxItems returns the global max items for a given type.
func (d *Data) GetGlobalMaxItems(t types.StreamerType) int32 {
	if t < 0 || t >= types.StreamerMaxTypes {
		return 0
	}
	return d.globalMaxItems[t]
}

// SetGlobalMaxItems sets the global max items for a given type.
func (d *Data) SetGlobalMaxItems(t types.StreamerType, maxItems int32) {
	if t < 0 || t >= types.StreamerMaxTypes {
		return
	}
	d.globalMaxItems[t] = maxItems
}

// GetGlobalMaxVisibleItems returns the global max visible items for a given type.
func (d *Data) GetGlobalMaxVisibleItems(t types.StreamerType) int32 {
	if t < 0 || t >= types.StreamerMaxTypes {
		return 0
	}
	return d.globalMaxVisibleItems[t]
}

// SetGlobalMaxVisibleItems sets the global max visible items for a given type.
func (d *Data) SetGlobalMaxVisibleItems(t types.StreamerType, maxItems int32) {
	if t < 0 || t >= types.StreamerMaxTypes {
		return
	}
	d.globalMaxVisibleItems[t] = maxItems
}

// GetGlobalRadiusMultiplier returns the global radius multiplier for a given type.
func (d *Data) GetGlobalRadiusMultiplier(t types.StreamerType) float32 {
	if t < 0 || t >= types.StreamerMaxTypes {
		return 0
	}
	return d.globalRadiusMultipliers[t]
}

// SetGlobalRadiusMultiplier sets the global radius multiplier for a given type.
func (d *Data) SetGlobalRadiusMultiplier(t types.StreamerType, multiplier float32) {
	if t < 0 || t >= types.StreamerMaxTypes {
		return
	}
	d.globalRadiusMultipliers[t] = multiplier
}
