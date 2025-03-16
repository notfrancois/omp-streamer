package streamer

import "sync"

// ItemPair almacena un par de prioridad y distancia para un item
type ItemPair struct {
	Priority int
	Distance float32
	ItemID   int
	ItemPtr  interface{} // Puntero al item real (Object, TextLabel, etc.)
}

// ItemBimap implementa la funcionalidad del boost::bimap usado en el código original
type ItemBimap struct {
	items    []*ItemPair
	itemsMap map[int]*ItemPair
	mu       sync.RWMutex
}

// NewItemBimap crea un nuevo mapa bidireccional vacío
func NewItemBimap() *ItemBimap {
	return &ItemBimap{
		items:    make([]*ItemPair, 0),
		itemsMap: make(map[int]*ItemPair),
	}
}

// Insert añade un nuevo item ordenado por prioridad y distancia
func (b *ItemBimap) Insert(priority int, distance float32, id int, ptr interface{}) {
	b.mu.Lock()
	defer b.mu.Unlock()

	pair := &ItemPair{
		Priority: priority,
		Distance: distance,
		ItemID:   id,
		ItemPtr:  ptr,
	}

	b.itemsMap[id] = pair

	// Insertar ordenando por prioridad (mayor primero) y luego por distancia (menor primero)
	insertIdx := 0
	for i, item := range b.items {
		if item.Priority < priority || (item.Priority == priority && item.Distance > distance) {
			insertIdx = i
			break
		}
		insertIdx = i + 1
	}

	// Insertar en la posición correcta
	if insertIdx == len(b.items) {
		b.items = append(b.items, pair)
	} else {
		b.items = append(b.items, nil)
		copy(b.items[insertIdx+1:], b.items[insertIdx:])
		b.items[insertIdx] = pair
	}
}

// Find busca un item por ID
func (b *ItemBimap) Find(id int) (bool, *ItemPair) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if pair, exists := b.itemsMap[id]; exists {
		return true, pair
	}
	return false, nil
}

// Remove elimina un item por ID
func (b *ItemBimap) Remove(id int) bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	pair, exists := b.itemsMap[id]
	if !exists {
		return false
	}

	// Encontrar y eliminar de la lista ordenada
	for i, item := range b.items {
		if item == pair {
			b.items = append(b.items[:i], b.items[i+1:]...)
			break
		}
	}

	// Eliminar del mapa
	delete(b.itemsMap, id)
	return true
}

// Clear vacía el bimap
func (b *ItemBimap) Clear() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.items = b.items[:0]
	b.itemsMap = make(map[int]*ItemPair)
}

// GetFirst devuelve el primer elemento (mayor prioridad, menor distancia)
func (b *ItemBimap) GetFirst() (*ItemPair, bool) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if len(b.items) > 0 {
		return b.items[0], true
	}
	return nil, false
}

// GetLast returns the last item in the bimap (least priority/furthest)
func (b *ItemBimap) GetLast() (int, float32, int, interface{}) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if len(b.items) == 0 {
		return 0, 0.0, 0, nil
	}

	lastItem := b.items[len(b.items)-1]
	return lastItem.Priority, lastItem.Distance, lastItem.ItemID, lastItem.ItemPtr
}

// Size devuelve el número de elementos en el bimap
func (b *ItemBimap) Size() int {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return len(b.items)
}

// IsEmpty verifica si el bimap está vacío
func (b *ItemBimap) IsEmpty() bool {
	return b.Size() == 0
}
