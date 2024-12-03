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

type ObjectID int32

// ObjectDistance represents the distance of an object
type ObjectDistance float32

type ObjectDistanceType int32

const (
	STREAMER_OBJECT_SD ObjectDistanceType = iota
	STREAMER_OBJECT_DD
)

const (
	STREAMER_OBJECT_SD_DEFAULT = ObjectDistance(300.0)
	STREAMER_OBJECT_DD_DEFAULT = ObjectDistance(300.0)
)

// MaterialSize represents the sizes of materials available
type MaterialSize int32

const (
	OBJECT_MATERIAL_SIZE_256x128 MaterialSize = iota
	// Add other sizes as needed
)
