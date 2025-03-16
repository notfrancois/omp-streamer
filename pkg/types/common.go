/*
 * Copyright (C) 2024 wasabi aka Copacabana.
 * Streamer plugin for open.mp for Go.
 *
 * Special thanks to Incognito for the original Streamer plugin for SA-MP.
 * And to the Open Multiplayer Community.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package types

type StreamerType int32

const (
	// Invalid IDs
	InvalidPickupID   int32 = -1
	InvalidStreamerID int32 = 0

	// Limits
	StreamerMaxTypes       = 8
	StreamerMaxAreaTypes   = 6
	StreamerMaxObjectTypes = 3

	// Streamer types
	StreamerTypeObject      StreamerType = iota
	StreamerTypePickup      StreamerType = 1
	StreamerTypeCP          StreamerType = 2
	StreamerTypeRaceCP      StreamerType = 3
	StreamerTypeMapIcon     StreamerType = 4
	StreamerType3DTextLabel StreamerType = 5
	StreamerTypeArea        StreamerType = 6
	StreamerTypeActor       StreamerType = 7

	// Area types
	StreamerAreaTypeCircle    = iota
	StreamerAreaTypeCylinder  = 1
	StreamerAreaTypeSphere    = 2
	StreamerAreaTypeRectangle = 3
	StreamerAreaTypeCuboid    = 4
	StreamerAreaTypePolygon   = 5

	// Object types
	StreamerObjectTypeGlobal  = 0
	StreamerObjectTypePlayer  = 1
	StreamerObjectTypeDynamic = 2

	// Static distance cutoff
	StreamerStaticDistanceCutoff float32 = 0.0
)

type AreaPosition interface {
	Contains(pos Vector3) bool
	GetBounds() Box2D
	GetCenter() Vector2
}

// Box2D represents a 2D box
type Box2D struct {
	Min, Max *Vector2
}

// Box3D represents a 3D box
type Box3D struct {
	Min, Max Vector3
}

// Vector2 represents a 2D vector
type Vector2 struct {
	X, Y float32
}

// Polygon2D represents a 2D polygon
type Polygon2D struct {
	Points []Vector2
}

type Point2DPosition Vector2
type Point3DPosition Vector3
