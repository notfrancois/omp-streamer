package bitsets

const MAX_PLAYERS = 1000

type PlayerBitSet []uint64

// NewPlayerBitSet creates a new player bitset
func NewPlayerBitSet() PlayerBitSet {
	return make([]uint64, (MAX_PLAYERS+63)/64)
}

// Set sets a bit in the bitset
func (b PlayerBitSet) Set(index int) {
	b[index/64] |= 1 << uint(index%64)
}

// Clear clears a bit in the bitset
func (b PlayerBitSet) Clear(index int) {
	b[index/64] &^= 1 << uint(index%64)
}

// Test checks if a bit is set in the bitset
func (b PlayerBitSet) Test(index int) bool {
	return b[index/64]&(1<<uint(index%64)) != 0
}
