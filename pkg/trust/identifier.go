package trust

import "sort"

// Identifier is a simple identifier that can be used to track a set of items.
// It is used to track the internal ids of items that are being processed by a player or a chunk.
type Identifier struct {
	ids    []int32
	unused []int32
}

// NewIdentifier creates a new identifier.
func NewIdentifier() *Identifier {
	return &Identifier{
		ids:    make([]int32, 0),
		unused: make([]int32, 0),
	}
}

// Get returns the next available id.
func (id *Identifier) Get() int32 {
	// first, try to use an unused id
	if len(id.unused) > 0 {
		lastIndex := len(id.unused) - 1
		value := id.unused[lastIndex]
		id.unused = id.unused[:lastIndex]
		return value
	}

	// if there are no unused ids, create a new one
	newID := int32(len(id.ids))
	id.ids = append(id.ids, newID)
	return newID
}

// Remove marks an id as unused.
func (id *Identifier) Remove(value int32, maxValue int32) {
	// if the value is within the valid range
	if value >= 0 && int(value) < len(id.ids) {
		// mark as unused
		id.unused = append(id.unused, value)

		// if we have too many unused ids, reduce the slice
		if len(id.unused) > int(maxValue/2) {
			id.Optimize()
		}
	}
}

// Optimize optimizes the identifier by removing duplicate ids and sorting the unused ids.
func (id *Identifier) Optimize() {
	// sort unused ids
	sort.Slice(id.unused, func(i, j int) bool {
		return id.unused[i] < id.unused[j]
	})

	// remove duplicate ids
	if len(id.unused) > 0 {
		newUnused := make([]int32, 0, len(id.unused))
		lastID := id.unused[0]
		newUnused = append(newUnused, lastID)

		for i := 1; i < len(id.unused); i++ {
			if id.unused[i] != lastID {
				lastID = id.unused[i]
				newUnused = append(newUnused, lastID)
			}
		}
		id.unused = newUnused
	}
}
