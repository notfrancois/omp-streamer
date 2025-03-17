package streamer

import "sync"

// Grid implementa un sistema de cuadrícula espacial para particionamiento eficiente
type Grid struct {
	core     *Core
	cellSize float32
	cells    map[CellID]*Cell
	mu       sync.RWMutex
}

// NewGrid crea una nueva cuadrícula espacial
func NewGrid(core *Core) *Grid {
	return &Grid{
		core:     core,
		cellSize: 300.0, // Valor por defecto
		cells:    make(map[CellID]*Cell),
	}
}
