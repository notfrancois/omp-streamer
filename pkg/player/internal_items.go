package player

// InternalItems wraps the internal items of a player.
type InternalItems struct {
	InternalAreas      map[int32]struct{}
	InternalMapIcons   map[int32]int32
	InternalObjects    map[int32]int32
	InternalTextLabels map[int32]int32

	RemovedAreas      map[int32]struct{}
	RemovedMapIcons   map[int32]struct{}
	RemovedObjects    map[int32]struct{}
	RemovedTextLabels map[int32]struct{}
}

// NewInternalItems creates a new InternalItems instance.
func NewInternalItems() *InternalItems {
	return &InternalItems{
		InternalAreas:      make(map[int32]struct{}),
		InternalMapIcons:   make(map[int32]int32),
		InternalObjects:    make(map[int32]int32),
		InternalTextLabels: make(map[int32]int32),

		RemovedAreas:      make(map[int32]struct{}),
		RemovedMapIcons:   make(map[int32]struct{}),
		RemovedObjects:    make(map[int32]struct{}),
		RemovedTextLabels: make(map[int32]struct{}),
	}
}

// AddInternalMapIcon adds an internal map icon to the player.
func (i *InternalItems) AddInternalMapIcon(itemID, internalID int32) {
	i.InternalMapIcons[itemID] = internalID
}

// RemoveInternalMapIcon removes an internal map icon from the player.
func (i *InternalItems) RemoveInternalMapIcon(itemID int32) {
	delete(i.InternalMapIcons, itemID)
	i.RemovedMapIcons[itemID] = struct{}{}
}
