package bitsets

import "github.com/notfrancois/omp-streamer/pkg/types"

type ItemBitSet []uint64

// NewItemBitSet creates a new item bitset
func NewItemBitSet() ItemBitSet {
	return make([]uint64, (types.StreamerMaxTypes+63)/64)
}

// Set sets a bit in the bitset
func (b ItemBitSet) Set(index int) {
	b[index/64] |= 1 << uint(index%64)
}

func (b ItemBitSet) SetAll() {
	for i := range b {
		b[i] = ^uint64(0)
	}
}

// Clear clears a bit in the bitset
func (b ItemBitSet) Clear(index int) {
	b[index/64] &^= 1 << uint(index%64)
}

// Test checks if a bit is set in the bitset
func (b ItemBitSet) Test(index int) bool {
	return b[index/64]&(1<<uint(index%64)) != 0
}
