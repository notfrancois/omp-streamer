package streamer

// StreamCallback representa una callback de streaming
type StreamCallback struct {
	ItemType int // Tipo de elemento (objeto, pickup, etc.)
	ItemID   int // ID del elemento
	PlayerID int // ID del jugador
}

// Bitset es una implementación simplificada de std::bitset de C++
type Bitset struct {
	bits uint64
}

// IsEmpty verifica si todos los bits están en 0
func (b *Bitset) IsEmpty() bool {
	return b.bits == 0
}

// Set establece un bit en 1
func (b *Bitset) Set(index int) {
	if index >= 0 && index < 64 {
		b.bits |= (1 << uint(index))
	}
}

// ResetBit establece un bit específico en 0
func (b *Bitset) ResetBit(index int) {
	if index >= 0 && index < 64 {
		b.bits &= ^(1 << uint(index))
	}
}

// Reset establece todos los bits en 0
func (b *Bitset) Reset() {
	b.bits = 0
}

// IsSet verifica si un bit está en 1
func (b *Bitset) IsSet(index int) bool {
	if index >= 0 && index < 64 {
		return (b.bits & (1 << uint(index))) != 0
	}
	return false
}
