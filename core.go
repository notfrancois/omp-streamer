package streamer

import "sync"

// Core representa el sistema central de streaming
type Core struct {
	Data          *StreamerData
	Grid          *Grid
	Streamer      *Streamer
	ChunkStreamer *ChunkStreamer
}

// NewCore crea una nueva instancia del core del sistema
func NewCore() *Core {
	c := &Core{
		Data: NewStreamerData(),
	}
	c.Grid = NewGrid(c)
	c.Streamer = NewStreamer(c)
	c.ChunkStreamer = NewChunkStreamer(c)
	return c
}

// StreamerData almacena todos los datos del streamer
type StreamerData struct {
	Objects         map[int]*Object
	Pickups         map[int]*Pickup
	Checkpoints     map[int]*Checkpoint
	RaceCheckpoints map[int]*RaceCheckpoint
	MapIcons        map[int]*MapIcon
	TextLabels      map[int]*TextLabel
	Areas           map[int]*Area
	Actors          map[int]*Actor
	Players         map[int]*Player
	TypePriority    []int
	Interfaces      map[int]bool
	mu              sync.RWMutex
}

// NewStreamerData crea una nueva estructura de datos para el streamer
func NewStreamerData() *StreamerData {
	sd := &StreamerData{
		Objects:         make(map[int]*Object),
		Pickups:         make(map[int]*Pickup),
		Checkpoints:     make(map[int]*Checkpoint),
		RaceCheckpoints: make(map[int]*RaceCheckpoint),
		MapIcons:        make(map[int]*MapIcon),
		TextLabels:      make(map[int]*TextLabel),
		Areas:           make(map[int]*Area),
		Actors:          make(map[int]*Actor),
		Players:         make(map[int]*Player),
		TypePriority:    make([]int, StreamerMaxTypes),
		Interfaces:      make(map[int]bool),
	}

	// Configurar prioridad por defecto
	for i := 0; i < StreamerMaxTypes; i++ {
		sd.TypePriority[i] = i
	}

	return sd
}
