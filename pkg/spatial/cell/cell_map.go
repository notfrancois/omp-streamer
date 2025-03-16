package cell

import (
	"sync"

	"github.com/notfrancois/omp-streamer/pkg/types/common"
)

// CellMap is a map of cells.
type CellMap struct {
	mu    sync.RWMutex
	items map[string]*Cell
}

// NewCellMap creates a new CellMap.
func NewCellMap() *CellMap {
	return &CellMap{
		items: make(map[string]*Cell),
	}
}

// Get returns a cell from the map.
func (m *CellMap) Get(id common.CellID) (*Cell, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	cell, exists := m.items[id.Hash()]
	return cell, exists
}

// Set sets a cell in the map.
func (m *CellMap) Set(id common.CellID, cell *Cell) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.items[id.Hash()] = cell
}

// Delete deletes a cell from the map.
func (m *CellMap) Delete(id common.CellID) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.items, id.Hash())
}
