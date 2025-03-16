package core

import (
	"context"
	"sync"
	"time"

	"github.com/notfrancois/omp-streamer/internal/spatial"
	"github.com/notfrancois/omp-streamer/internal/streaming"
	"github.com/notfrancois/omp-streamer/pkg/openmp"
	"github.com/notfrancois/omp-streamer/pkg/types"
)

type Manager struct {
	mu       sync.RWMutex
	streamer *streaming.Streamer
	chunkStr *streaming.ChunkStreamer
	grid     *spatial.Grid
	gamesdk  *openmp.API

	// Configuración
	config *types.StreamerConfig

	// Contadores y límites
	objectCount int32
	maxObjects  int32

	// Control de contexto
	ctx    context.Context
	cancel context.CancelFunc
}

func NewManager(ctx context.Context, cfg *types.StreamerConfig, api *openmp.API) *Manager {
	ctx, cancel := context.WithCancel(ctx)
	grid := spatial.NewGrid(cfg.GridSize)

	return &Manager{
		config:     cfg,
		grid:       grid,
		streamer:   streaming.NewStreamer(ctx, grid, api),
		chunkStr:   streaming.NewChunkStreamer(ctx, grid),
		gamesdk:    api,
		maxObjects: types.ServerObjectLimit,
		ctx:        ctx,
		cancel:     cancel,
	}
}

// CreateObject crea un nuevo objeto dinámico
func (m *Manager) CreateObject(modelID int32, pos, rot types.Vector3, drawDistance float32) (*types.DynamicObject, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.objectCount >= m.maxObjects {
		return nil, ErrMaxObjectsReached
	}

	obj := &types.DynamicObject{
		ID:           m.generateObjectID(),
		ModelID:      modelID,
		Position:     pos,
		Rotation:     rot,
		DrawDistance: drawDistance,
		Materials:    make(map[int32]*types.Material),
	}

	// Insertar en el grid
	m.grid.InsertObject(obj)

	// Intentar hacer stream si hay slots disponibles
	if err := m.streamer.TryStreamIn(obj); err != nil {
		return nil, err
	}

	m.objectCount++
	return obj, nil
}

// DestroyObject destruye un objeto dinámico
func (m *Manager) DestroyObject(objID int32) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	obj := m.getObject(objID)
	if obj == nil {
		return ErrInvalidObject
	}

	// Remover del grid
	m.grid.RemoveObject(obj)

	// Hacer stream out si está streamed
	if obj.IsStreamed {
		if err := m.streamer.StreamOut(obj); err != nil {
			return err
		}
	}

	m.objectCount--
	return nil
}

// UpdateObject actualiza la posición de un objeto
func (m *Manager) UpdateObject(objID int32, pos, rot types.Vector3) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	obj := m.getObject(objID)
	if obj == nil {
		return ErrInvalidObject
	}

	// Actualizar grid si la posición cambió
	if pos != obj.Position {
		m.grid.RemoveObject(obj)
		obj.Position = pos
		m.grid.InsertObject(obj)
	}

	obj.Rotation = rot

	// Actualizar objeto físico si está streamed
	if obj.IsStreamed {
		return m.serverAPI.UpdateObject(obj)
	}

	return nil
}

// ProcessStreaming procesa el streaming para un jugador
func (m *Manager) ProcessStreaming(playerID int32, pos types.Vector3) {
	// 1. Procesar chunks para descubrir objetos
	discoveredObjects := m.chunkStr.ProcessChunk(playerID, pos)

	// 2. Intentar hacer stream in de objetos descubiertos
	for _, obj := range discoveredObjects {
		if !obj.IsStreamed {
			if err := m.streamer.TryStreamIn(obj); err != nil {
				// Manejar error o encolar para retry
				m.streamer.VirtualQueue.PushObject(obj)
			}
		}
	}

	// 3. Verificar objetos que deben hacer stream out
	removedObjects := m.chunkStr.GetRemovedObjects(playerID)
	for _, obj := range removedObjects {
		if obj.IsStreamed {
			if err := m.streamer.StreamOut(obj); err != nil {
				// Manejar error
				continue
			}
		}
	}

	// 4. Procesar cola de prioridad
	m.streamer.ProcessQueue()
}

// SetObjectMaterial establece un material para un objeto
func (m *Manager) SetObjectMaterial(objID int32, materialIndex int32, material *types.Material) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	obj := m.getObject(objID)
	if obj == nil {
		return ErrInvalidObject
	}

	obj.Materials[materialIndex] = material

	// Actualizar material si el objeto está streamed
	if obj.IsStreamed {
		return m.serverAPI.SetObjectMaterial(obj.ID, materialIndex, material)
	}

	return nil
}

// Métodos auxiliares

func (m *Manager) generateObjectID() int32 {
	// Implementar generación de IDs únicos
	return m.objectCount + 1
}

func (m *Manager) getObject(id int32) *types.DynamicObject {
	// Implementar obtención de objeto por ID
	return nil // TODO: Implementar
}

// Start inicia el manager y sus componentes
func (m *Manager) Start() error {
	// Iniciar rutinas de actualización
	go m.updateLoop()
	return nil
}

// Stop detiene el manager y sus componentes
func (m *Manager) Stop() {
	m.cancel()
}

func (m *Manager) updateLoop() {
	ticker := time.NewTicker(m.config.UpdateInterval)
	defer ticker.Stop()

	for {
		select {
		case <-m.ctx.Done():
			return
		case <-ticker.C:
			m.update()
		}
	}
}

func (m *Manager) update() {
	// Actualizar streaming global
	// Procesar cola de prioridad
	// Limpiar objetos no usados
	// etc...
}
