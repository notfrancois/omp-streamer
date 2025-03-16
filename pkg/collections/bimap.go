package collections

import "sync"

// Bimap implements a thread-safe bidirectional map
type Bimap[K comparable, V comparable] struct {
	forward map[K]V
	reverse map[V]K
	mu      sync.RWMutex
}

// NewBimap creates a new Bimap
func NewBimap[K comparable, V comparable]() *Bimap[K, V] {
	return &Bimap[K, V]{
		forward: make(map[K]V),
		reverse: make(map[V]K),
	}
}

// Insert adds a key-value pair to the bimap
func (b *Bimap[K, V]) Insert(key K, value V) {
	b.mu.Lock()
	defer b.mu.Unlock()

	// If the key already exists, we first remove the old value
	if oldValue, exists := b.forward[key]; exists {
		delete(b.reverse, oldValue)
	}

	// If the value already exists, we remove the old key
	if oldKey, exists := b.reverse[value]; exists {
		delete(b.forward, oldKey)
	}

	b.forward[key] = value
	b.reverse[value] = key
}

// GetByKey gets a value using the key
func (b *Bimap[K, V]) GetByKey(key K) (V, bool) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	val, exists := b.forward[key]
	return val, exists
}

// GetByValue gets a key using the value
func (b *Bimap[K, V]) GetByValue(value V) (K, bool) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	key, exists := b.reverse[value]
	return key, exists
}

// RemoveByKey removes an entry using the key
func (b *Bimap[K, V]) RemoveByKey(key K) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if value, exists := b.forward[key]; exists {
		delete(b.forward, key)
		delete(b.reverse, value)
	}
}

// RemoveByValue removes an entry using the value
func (b *Bimap[K, V]) RemoveByValue(value V) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if key, exists := b.reverse[value]; exists {
		delete(b.forward, key)
		delete(b.reverse, value)
	}
}

// Len returns the number of elements in the bimap
func (b *Bimap[K, V]) Len() int {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return len(b.forward)
}

// Clear clears all elements from the bimap
func (b *Bimap[K, V]) Clear() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.forward = make(map[K]V)
	b.reverse = make(map[V]K)
}
