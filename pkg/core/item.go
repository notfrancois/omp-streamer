package core

import (
	"fmt"
	"hash/fnv"

	"github.com/notfrancois/omp-streamer/pkg/types/common"
)

type Item struct {
	SharedActor          *common.Actor
	SharedArea           *common.Area
	SharedCheckpoint     *common.Checkpoint
	SharedMapIcon        *common.MapIcon
	SharedObject         *common.Object
	SharedPickup         *common.Pickup
	SharedRaceCheckpoint *common.RaceCheckpoint
	SharedTextLabel      *common.TextLabel

	Hash    PairHash
	Compare PairCompare
	Tuple   TupleCompare
}

// PairHash implements a hash function for pairs of values.
type PairHash struct{}

// Hash implements a hash function for pairs of values
func (p *PairHash) Hash(v1, v2 interface{}) uint64 {
	h := fnv.New64()

	// Generate hash for the first value
	switch val1 := v1.(type) {
	case int:
		fmt.Fprintf(h, "i%d", val1)
	case float64:
		fmt.Fprintf(h, "f%.6f", val1)
	case string:
		fmt.Fprintf(h, "s%s", val1)
	default:
		fmt.Fprintf(h, "o%v", val1)
	}

	// Combine with hash of the second value
	switch val2 := v2.(type) {
	case int:
		fmt.Fprintf(h, "i%d", val2)
	case float64:
		fmt.Fprintf(h, "f%.6f", val2)
	case string:
		fmt.Fprintf(h, "s%s", val2)
	default:
		fmt.Fprintf(h, "o%v", val2)
	}

	return h.Sum64()
}

// PairCompare implements comparison of pairs
type PairCompare struct{}

// Compare implements comparison of pairs
func (p *PairCompare) Compare(v1, v2 interface{}) bool {
	// Compare types first
	t1 := fmt.Sprintf("%T", v1)
	t2 := fmt.Sprintf("%T", v2)
	if t1 != t2 {
		return false
	}

	// Specific type comparison
	switch val1 := v1.(type) {
	case int:
		val2 := v2.(int)
		return val1 == val2
	case float64:
		val2 := v2.(float64)
		return val1 == val2
	case string:
		val2 := v2.(string)
		return val1 == val2
	default:
		// Generic comparison for other types
		return fmt.Sprintf("%v", v1) == fmt.Sprintf("%v", v2)
	}
}

// TupleCompare implements comparison of tuples
type TupleCompare struct{}

// Compare implements comparison of tuples
func (t *TupleCompare) Compare(t1, t2 Tuple) bool {
	// If First is different, sort by First consistently
	if t1.First != t2.First {
		return t1.First < t2.First
	}

	// Try as int first
	if v1, ok1 := t1.Second.(int); ok1 {
		if v2, ok2 := t2.Second.(int); ok2 {
			return v1 < v2
		}
	}

	// Try as float64
	if v1, ok1 := t1.Second.(float64); ok1 {
		if v2, ok2 := t2.Second.(float64); ok2 {
			return v1 < v2
		}
	}

	// If not int or float, use string comparison as fallback
	return fmt.Sprintf("%v", t1.Second) < fmt.Sprintf("%v", t2.Second)
}

// Tuple represents a tuple of two elements
type Tuple struct {
	First  int
	Second interface{}
}

// PairKey represents a pair of values that can be used as a key in a map
type PairKey struct {
	First  interface{}
	Second interface{}
}

// Hash generates a unique hash for the pair of values
func (p PairKey) Hash() uint64 {
	hasher := &PairHash{}
	return hasher.Hash(p.First, p.Second)
}

// Equals compares two PairKeys for equality
func (p PairKey) Equals(other PairKey) bool {
	comparer := &PairCompare{}
	return comparer.Compare(p.First, other.First) &&
		comparer.Compare(p.Second, other.Second)
}
