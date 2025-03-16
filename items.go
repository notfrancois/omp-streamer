package streamer

// Attachment representa información de adjuntos para objetos y áreas
type Attachment struct {
	Object   int
	Player   int
	Vehicle  int
	Position Vector3
	Worlds   []int
}

// Object representa un objeto dinámico en el mundo
type Object struct {
	ObjectId                 int
	ModelId                  int
	Position                 Vector3
	PositionOffset           Vector3
	Rotation                 Vector3
	DrawDistance             float32
	StreamDistance           float32
	ComparableStreamDistance float32
	MoveDuration             int
	MoveSpeed                float32
	Cell                     *Cell
	Attach                   *Attachment
	Priority                 int
	Worlds                   []int
	Interiors                []int
	Players                  []int
	Areas                    []int
	Materials                map[int]Material
	NoCameraCollision        bool
	InverseAreaChecking      bool
	StreamCallbacks          bool
}

// TextLabel representa una etiqueta de texto 3D
type TextLabel struct {
	TextLabelId              int
	Text                     string
	Color                    int
	Position                 Vector3
	PositionOffset           Vector3
	DrawDistance             float32
	StreamDistance           float32
	ComparableStreamDistance float32
	TestLOS                  bool
	Cell                     *Cell
	Attach                   *Attachment
	Priority                 int
	Worlds                   []int
	Interiors                []int
	Players                  []int
	Areas                    []int
	InverseAreaChecking      bool
	StreamCallbacks          bool
}

// MapIcon representa un icono en el mapa del jugador
type MapIcon struct {
	MapIconId                int
	ModelId                  int
	Type                     int
	Color                    int
	Position                 Vector3
	PositionOffset           Vector3
	StreamDistance           float32
	ComparableStreamDistance float32
	Style                    int
	Cell                     *Cell
	Priority                 int
	Worlds                   []int
	Interiors                []int
	Players                  []int
	Areas                    []int
	InverseAreaChecking      bool
	StreamCallbacks          bool
}

// Material representa materiales para objetos
type Material struct {
	Main *MaterialMain
	Text *MaterialText
}

// MaterialMain representa materiales principales para objetos
type MaterialMain struct {
	ModelId       int
	TxdFileName   string
	TextureName   string
	MaterialColor int
}

// MaterialText representa texto como material para objetos
type MaterialText struct {
	MaterialText  string
	MaterialSize  int
	FontFace      string
	FontSize      int
	Bold          bool
	FontColor     int
	BackColor     int
	TextAlignment int
}

// Más tipos de items (Pickup, Checkpoint, etc.) se definirían similarmente...
type Pickup struct {
	PickupId       int
	ModelId        int
	Position       Vector3
	PositionOffset Vector3
}

type Checkpoint struct {
	CheckpointId   int
	Position       Vector3
	PositionOffset Vector3
}

type RaceCheckpoint struct {
	RaceCheckpointId int
	Position         Vector3
	PositionOffset   Vector3
}

type Area struct {
	AreaId         int
	Position       Vector3
	PositionOffset Vector3
}

type Actor struct {
	ActorId        int
	ModelId        int
	Position       Vector3
	PositionOffset Vector3
}
