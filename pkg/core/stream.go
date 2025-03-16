package core

import (
	"time"

	"github.com/notfrancois/omp-streamer/internal/core/bitsets"
	"github.com/notfrancois/omp-streamer/internal/core/events"
	"github.com/notfrancois/omp-streamer/pkg/types/common"
)

type Streamer struct {
	AttachedAreas      map[int32]*common.Area
	AttachedObjects    map[int32]*common.Object
	AttachedTextLabels map[int32]*common.TextLabel
	MovingObjects      map[int32]*common.Object

	// Misc
	TickCount, TickRate int32

	AverageElapsedTime, LastUpdateTime float64

	VelocityBoundaries struct {
		Min, Max float32
	}

	events *events.EventManager

	lastRecordedTime int64
	recordedTimes    []float64
}

// NewStreamer creates a new Streamer instance
func NewStreamer() *Streamer {
	return &Streamer{
		AverageElapsedTime: 0.0,
		LastUpdateTime:     0.0,
		TickCount:          0,
		TickRate:           50,
		VelocityBoundaries: struct {
			Min, Max float32
		}{
			Min: 0.25,
			Max: 7.5,
		},
		events: events.NewEventManager(),
	}
}

// GetLastUpdateTime returns the last update time
func (s *Streamer) GetLastUpdateTime() float64 {
	return s.LastUpdateTime
}

// GetTickCount returns the tick count
func (s *Streamer) GetTickCount() int32 {
	return s.TickCount
}

// GetTickRate returns the tick rate
func (s *Streamer) GetTickRate() int32 {
	return s.TickRate
}

// GetAverageElapsedTime returns the average elapsed time
func (s *Streamer) GetAverageElapsedTime() float64 {
	return s.AverageElapsedTime
}

// GetVelocityBoundaries returns the velocity boundaries
func (s *Streamer) GetVelocityBoundaries() (float32, float32) {
	return s.VelocityBoundaries.Min, s.VelocityBoundaries.Max
}

// GetAttachedAreas returns the attached areas
func (s *Streamer) GetAttachedAreas() map[int32]*common.Area {
	return s.AttachedAreas
}

// GetAttachedObjects returns the attached objects
func (s *Streamer) GetAttachedObjects() map[int32]*common.Object {
	return s.AttachedObjects
}

// GetAttachedTextLabels returns the attached text labels
func (s *Streamer) GetAttachedTextLabels() map[int32]*common.TextLabel {
	return s.AttachedTextLabels
}

// GetMovingObjects returns the moving objects
func (s *Streamer) GetMovingObjects() map[int32]*common.Object {
	return s.MovingObjects
}

// SetTickCount sets the tick count
func (s *Streamer) SetTickCount(tickCount int32) bool {
	if tickCount > 0 { // 0 is invalid tick count
		s.TickCount = tickCount
		return true
	}
	return false
}

// SetTickRate sets the tick rate
func (s *Streamer) SetTickRate(tickRate int32) bool {
	if tickRate > 0 { // 0 is invalid tick rate
		s.TickRate = tickRate
		return true
	}
	return false
}

// DoesPlayerSatisfyConditions checks if a player satisfies the basic conditions
func (s *Streamer) DoesPlayerSatisfyConditions(
	bitset bitsets.PlayerBitSet,
	index int,
	set1 map[int]struct{},
	val1 int,
	set2 map[int]struct{},
	val2 int,
) bool {
	// Checks the bitset and that the value is in the sets (if they are not empty)
	return bitset.Test(index) &&
		(len(set1) == 0 || containsKey(set1, val1)) &&
		(len(set2) == 0 || containsKey(set2, val2))
}

// DoesPlayerSatisfyConditionsExtended checks extended conditions for a player
func (s *Streamer) DoesPlayerSatisfyConditionsExtended(
	bitset bitsets.PlayerBitSet,
	index int,
	set1 map[int]struct{},
	val1 int,
	set2 map[int]struct{},
	val2 int,
	set3 map[int]struct{},
	set4 map[int]struct{},
	flag bool,
) bool {
	baseCondition := s.DoesPlayerSatisfyConditions(bitset, index, set1, val1, set2, val2)
	if !baseCondition {
		return false
	}

	if len(set3) == 0 {
		return true
	}

	if flag {
		return !isContainerWithinContainer(set3, set4)
	}
	return isContainerWithinContainer(set3, set4)
}

// containsKey checks if a key exists in a map
func containsKey(m map[int]struct{}, key int) bool {
	_, exists := m[key]
	return exists
}

// isContainerWithinContainer checks if one container is within another
func isContainerWithinContainer(set1, set2 map[int]struct{}) bool {
	for k := range set1 {
		if _, exists := set2[k]; !exists {
			return false
		}
	}
	return true
}

func (s *Streamer) CalculateAverageElapsedTime() {
	currentTime := time.Now().UnixNano()

	if s.lastRecordedTime == 0 {
		s.lastRecordedTime = currentTime
		return
	}

	// initialize recordedTimes if it doesn't exist
	if s.recordedTimes == nil {
		s.recordedTimes = make([]float64, 5)
	}

	// check if there is space for more records
	allPositive := true
	nextIndex := 0
	for i, t := range s.recordedTimes {
		if t <= 0 {
			allPositive = false
			nextIndex = i
			break
		}
	}

	if !allPositive {
		// calculate the elapsed time in seconds
		elapsedTime := float64(currentTime-s.lastRecordedTime) / float64(time.Second)
		s.recordedTimes[nextIndex] = elapsedTime
	} else {
		// calculate the average and multiply by 50
		sum := 0.0
		for _, t := range s.recordedTimes {
			sum += t
		}
		s.AverageElapsedTime = (sum / 5.0) * 50.0

		// reset the array
		s.recordedTimes = make([]float64, 5)
	}

	s.lastRecordedTime = currentTime
}
