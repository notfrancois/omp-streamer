package streamer

type StreamerType int32
type StreamerTag int32

type ObjectID int32
type PickupID int32
type ActorID int32
type AreaID int32
type CheckpointID int32
type RaceCheckpointID int32
type MapIconID int32
type TextLabelID int32

type RaceCheckpointType int32
type MapIconType int32
type TextLabelColor int32

type MaterialSize int32
type MaterialColor int32
type MaterialTextAlignment int32

const (
	InvalidPickupID   = -1
	InvalidStreamerID = 0

	StreamerMaxTypes = 8

	StreamerTypeObject      = 0
	StreamerTypePickup      = 1
	StreamerTypeCP          = 2
	StreamerTypeRaceCP      = 3
	StreamerTypeMapIcon     = 4
	StreamerType3DTextLabel = 5
	StreamerTypeArea        = 6
	StreamerTypeActor       = 7

	StreamerMaxAreaTypes = 6

	StreamerAreaTypeCircle    = 0
	StreamerAreaTypeCylinder  = 1
	StreamerAreaTypeSphere    = 2
	StreamerAreaTypeRectangle = 3
	StreamerAreaTypeCuboid    = 4
	StreamerAreaTypePolygon   = 5

	StreamerMaxObjectTypes = 3

	StreamerObjectTypeGlobal  = 0
	StreamerObjectTypePlayer  = 1
	StreamerObjectTypeDynamic = 2

	StreamerStaticDistanceCutoff = 0.0
)

type Vector3 struct {
	X float32
	Y float32
	Z float32
}

type Vector2 struct {
	X float32
	Y float32
}
