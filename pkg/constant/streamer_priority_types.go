package constant

import t "github.com/notfrancois/omp-streamer/pkg/types"

const (
	StreamerTypeArea t.StreamerPriorityType = iota
	StreamerTypeObject
	StreamerTypeCP
	StreamerTypeRaceCP
	StreamerTypeMapIcon
	StreamerType3DTextLabel
	StreamerTypePickup
	StreamerTypeActor
)
