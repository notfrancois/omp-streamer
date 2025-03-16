package streamer

import "time"

type Player struct {
	PlayerId                  int
	Position                  Vector3
	InteriorId                int
	WorldId                   int
	ChunkTickCount            map[int]int
	ChunkTickRate             map[int]int
	CurrentVisibleObjects     int
	CurrentVisibleTextLabels  int
	MaxVisibleObjects         int
	MaxVisibleMapIcons        int
	MaxVisibleTextLabels      int
	TickCount                 int
	TickRate                  int
	RadiusMultipliers         map[int]float32
	EnabledItems              []bool
	VisibleCell               *Cell
	InternalObjects           map[int]int // StreamerID -> InternalID
	InternalMapIcons          map[int]int
	InternalTextLabels        map[int]int
	InternalAreas             map[int]bool
	DiscoveredObjects         *ItemBimap
	DiscoveredMapIcons        *ItemBimap
	DiscoveredTextLabels      *ItemBimap
	ExistingObjects           *ItemBimap
	ExistingMapIcons          *ItemBimap
	ExistingTextLabels        *ItemBimap
	RemovedObjects            map[int]bool
	RemovedMapIcons           map[int]bool
	RemovedTextLabels         map[int]bool
	ProcessingChunks          Bitset
	UpdateUsingCameraPosition bool
	UpdateWhenIdle            bool
	ActiveCheckpoint          int
	ActiveRaceCheckpoint      int
	VisibleCheckpoint         int // Faltaba este campo
	VisibleRaceCheckpoint     int
	DelayedCheckpoint         int // Campos nuevos
	DelayedRaceCheckpoint     int
	DelayedUpdate             bool
	DelayedUpdateFreeze       bool
	DelayedUpdateTime         time.Time // Equivalente a std::chrono::steady_clock::time_point
	DelayedUpdateType         int
	References                int
	RequestingClass           bool
	LastUpdatePos             Vector3 // Para IsPlayerIdle (visto en handlers.go)

	// Identifiers...
	MapIconIdentifier        *Identifier
	ObjectIdentifier         *Identifier
	TextLabelIdentifier      *Identifier
	AreaIdentifier           *Identifier
	CheckpointIdentifier     *Identifier
	RaceCheckpointIdentifier *Identifier
	PickupIdentifier         *Identifier
	ActorIdentifier          *Identifier
}

// NewPlayer crea un nuevo jugador con valores por defecto
func NewPlayer(id int) *Player {
	p := &Player{
		PlayerId:                 id,
		ChunkTickCount:           make(map[int]int),
		ChunkTickRate:            make(map[int]int),
		RadiusMultipliers:        make(map[int]float32),
		EnabledItems:             make([]bool, StreamerMaxTypes),
		VisibleCell:              NewCell(),
		InternalObjects:          make(map[int]int),
		InternalMapIcons:         make(map[int]int),
		InternalTextLabels:       make(map[int]int),
		InternalAreas:            make(map[int]bool),
		RemovedObjects:           make(map[int]bool),
		RemovedMapIcons:          make(map[int]bool),
		RemovedTextLabels:        make(map[int]bool),
		DiscoveredObjects:        NewItemBimap(),
		DiscoveredMapIcons:       NewItemBimap(),
		DiscoveredTextLabels:     NewItemBimap(),
		ExistingObjects:          NewItemBimap(),
		ExistingMapIcons:         NewItemBimap(),
		ExistingTextLabels:       NewItemBimap(),
		MapIconIdentifier:        NewIdentifier(0, 1000),
		ObjectIdentifier:         NewIdentifier(0, 1000),
		TextLabelIdentifier:      NewIdentifier(0, 1000),
		AreaIdentifier:           NewIdentifier(0, 1000),
		CheckpointIdentifier:     NewIdentifier(0, 1000),
		RaceCheckpointIdentifier: NewIdentifier(0, 1000),
		PickupIdentifier:         NewIdentifier(0, 1000),
		ActorIdentifier:          NewIdentifier(0, 1000),
		TickRate:                 50,   // Valor por defecto
		UpdateWhenIdle:           true, // Valor por defecto
		References:               1,    // Inicialmente tiene 1 referencia
		RequestingClass:          false,
		DelayedUpdate:            false,
		DelayedUpdateFreeze:      false,
	}

	// Configuración inicial
	p.ChunkTickCount[StreamerTypeObject] = 0
	p.ChunkTickCount[StreamerTypeMapIcon] = 0
	p.ChunkTickCount[StreamerType3DTextLabel] = 0

	p.ChunkTickRate[StreamerTypeObject] = 1
	p.ChunkTickRate[StreamerTypeMapIcon] = 1
	p.ChunkTickRate[StreamerType3DTextLabel] = 1

	for i := 0; i < StreamerMaxTypes; i++ {
		p.EnabledItems[i] = true
	}

	return p
}
