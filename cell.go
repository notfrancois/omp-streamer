package streamer

type CellID struct {
	X, Y int
}

// Cell representa una celda en la cuadrícula espacial que contiene referencias a entidades
type Cell struct {
	Objects         map[int]*Object
	Pickups         map[int]*Pickup
	Checkpoints     map[int]*Checkpoint
	RaceCheckpoints map[int]*RaceCheckpoint
	MapIcons        map[int]*MapIcon
	TextLabels      map[int]*TextLabel
	Areas           map[int]*Area
	Actors          map[int]*Actor
}

// NewCell crea una nueva celda vacía
func NewCell() *Cell {
	return &Cell{
		Objects:         make(map[int]*Object),
		Pickups:         make(map[int]*Pickup),
		Checkpoints:     make(map[int]*Checkpoint),
		RaceCheckpoints: make(map[int]*RaceCheckpoint),
		MapIcons:        make(map[int]*MapIcon),
		TextLabels:      make(map[int]*TextLabel),
		Areas:           make(map[int]*Area),
		Actors:          make(map[int]*Actor),
	}
}
