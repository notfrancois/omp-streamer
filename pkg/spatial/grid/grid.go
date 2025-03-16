package spatial

import (
	"fmt"
	"math"
	"sync"

	"github.com/notfrancois/omp-streamer/pkg/spatial/cell"
	"github.com/notfrancois/omp-streamer/pkg/types"
	"github.com/notfrancois/omp-streamer/pkg/types/common"
	"gonum.org/v1/gonum/mat"
)

type Grid struct {
	mu                sync.RWMutex
	GlobalCell        common.Cell
	cells             *cell.CellMap
	objectIndex       map[int32]common.CellID
	cellSize          float32
	cellDistance      float32
	translationMatrix *mat.Dense
}

func NewGrid() *Grid {
	g := &Grid{
		GlobalCell:        cell.Create(),
		cells:             cell.NewCellMap(),
		objectIndex:       make(map[int32]common.CellID),
		cellSize:          300.0,
		cellDistance:      360000.0,
		translationMatrix: nil, // as default, the translation matrix is not calculated until it is needed.
	}

	g.calculateTranslationMatrix() // calculate the translation matrix for the grid based on the cell size.

	return g
}

func (g *Grid) GetCellID(position types.Vector2, insert bool) common.CellID {
	// creates a box with the minimum corner of the cell
	box := &types.Box2D{
		Min: &types.Vector2{
			X: float32(math.Floor(float64(position.X/g.cellSize))) * g.cellSize,
			Y: float32(math.Floor(float64(position.Y/g.cellSize))) * g.cellSize,
		},
	}

	// calculates the maximum corner of the cell
	box.Max = &types.Vector2{
		X: box.Min.X + g.cellSize,
		Y: box.Min.Y + g.cellSize,
	}

	// calculates the centroid of the cell
	centroid := types.Vector2{
		X: box.Min.X + g.cellSize/2,
		Y: box.Min.Y + g.cellSize/2,
	}

	// creates the cell ID
	cellID := common.CellID{
		X: int32(centroid.X),
		Y: int32(centroid.Y),
	}

	// if insert is true, create the cell if it doesn't exist
	if insert {
		g.mu.Lock()
		defer g.mu.Unlock()

		if _, exists := g.cells.Get(cellID); !exists {
			g.cells.Set(cellID, cell.CreateWithID(cellID))
		}
	}

	return cellID
}

// EraseCellIfEmpty erases a cell if there are no objects left in it.
func (g *Grid) EraseCellIfEmpty(cell *common.Cell) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if cell.IsEmpty() {
		delete(g.cells, cell.ID)
	}
}

// InsertObject inserta un objeto en la celda correspondiente
func (g *Grid) InsertObject(obj *types.DynamicObject) {
	cell := g.getCell(obj.Position)
	cellID := g.getCellID(obj.Position.X, obj.Position.Y)

	g.mu.Lock()
	g.objectIndex[obj.ID] = cellID
	g.mu.Unlock()

	cell.AddObject(obj)
}

// RemoveObject elimina un objeto de su celda
func (g *Grid) RemoveObject(obj *types.DynamicObject) {
	cell := g.getCell(obj.Position)

	g.mu.Lock()
	delete(g.objectIndex, obj.ID)
	g.mu.Unlock()

	cell.RemoveObject(obj.ID)
}

// GetNearbyObjects obtiene objetos cercanos a una posición
func (g *Grid) GetNearbyObjects(pos types.Vector3, radius float32) []*types.DynamicObject {
	cellRadius := int32(radius/g.cellSize) + 1
	baseX := int32(pos.X / g.cellSize)
	baseY := int32(pos.Y / g.cellSize)

	nearby := make([]*types.DynamicObject, 0)

	g.mu.RLock()
	defer g.mu.RUnlock()

	// Buscar en celdas cercanas
	for x := -cellRadius; x <= cellRadius; x++ {
		for y := -cellRadius; y <= cellRadius; y++ {
			cellID := fmt.Sprintf("%d:%d", baseX+x, baseY+y)
			if cell, exists := g.cells[cellID]; exists {
				for _, obj := range cell.GetObjects() {
					if obj.Position.DistanceTo(pos) <= radius {
						nearby = append(nearby, obj)
					}
				}
			}
		}
	}

	return nearby
}

// GetObject usando el índice
func (g *Grid) GetObject(id int32) *types.DynamicObject {
	g.mu.RLock()
	defer g.mu.RUnlock()

	if cellID, exists := g.objectIndex[id]; exists {
		if cell, exists := g.cells[cellID]; exists {
			cell.mu.RLock()
			obj := cell.Objects[id]
			cell.mu.RUnlock()
			return obj
		}
	}
	return nil
}

func (g *Grid) GetNearbyCells(pos types.Vector3, radius float32) []*common.Cell {
	g.mu.RLock()
	defer g.mu.RUnlock()

	cells := make([]*common.Cell, 0)
	cellRadius := int32(radius/g.cellSize) + 1
	baseX := int32(pos.X / g.cellSize)
	baseY := int32(pos.Y / g.cellSize)

	for x := baseX - cellRadius; x <= baseX+cellRadius; x++ {
		for y := baseY - cellRadius; y <= baseY+cellRadius; y++ {
			cellID := fmt.Sprintf("%d:%d", x, y)
			if cell, exists := g.cells[cellID]; exists {
				cells = append(cells, cell)
			}
		}
	}

	return cells
}

// GetObjectCount retorna el número total de objetos en el grid
func (g *Grid) GetObjectCount() int32 {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return int32(len(g.objectIndex))
}

// UpdateObjectPosition actualiza la posición de un objeto y lo mueve a la celda correcta si es necesario
func (g *Grid) UpdateObjectPosition(obj *types.DynamicObject, newPos types.Vector3) {
	oldCellID := g.objectIndex[obj.ID]
	newCellID := g.getCellID(newPos.X, newPos.Y)

	// Si cambió de celda
	if oldCellID != newCellID {
		g.mu.Lock()
		// Remover de celda anterior
		if oldCell, exists := g.cells[oldCellID]; exists {
			oldCell.RemoveObject(obj.ID)
		}

		// Actualizar posición
		obj.Position = newPos

		// Insertar en nueva celda
		newCell := g.getCell(newPos)
		newCell.AddObject(obj)

		// Actualizar índice
		g.objectIndex[obj.ID] = newCellID
		g.mu.Unlock()
	} else {
		// Solo actualizar posición
		obj.Position = newPos
	}
}

// calculateTranslationMatrix calculates the translation matrix for the grid.
func (g *Grid) calculateTranslationMatrix() {
	data := []float64{
		0.0, 0.0, float64(g.cellSize),
		float64(g.cellSize), float64(g.cellSize * -1.0), 0.0,
		float64(g.cellSize * -1.0), float64(g.cellSize), float64(g.cellSize * -1.0),
		0.0, float64(g.cellSize), 0.0,
		float64(g.cellSize), 0.0, float64(g.cellSize * -1.0),
		float64(g.cellSize), float64(g.cellSize * -1.0), float64(g.cellSize * -1.0),
	}
	g.translationMatrix = mat.NewDense(2, 9, data)
}
