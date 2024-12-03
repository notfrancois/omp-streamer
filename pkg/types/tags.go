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

// StreamerTag represents the different types of dynamic items
type StreamerTag int

const (
	Tag_DynamicObject StreamerTag = iota
	Tag_DynamicPickup
	Tag_DynamicCP
	Tag_DynamicRaceCP
	Tag_DynamicMapIcon
	Tag_DynamicText3D
	Tag_DynamicArea
	Tag_DynamicActor
)

// Text3DTagType determines which type of Text3D tag to use
type Text3DTagType int

const (
	StandardText3D   Text3DTagType = iota // Text3D
	DynamicText3DTag                      // DynamicText3D
	NoText3DTag                           // when STREAMER_REMOVE_TEXT3D_TAG is defined
)

// GetText3DTag returns the appropriate StreamerTag based on the configuration
func GetText3DTag(tagType Text3DTagType) StreamerTag {
	switch tagType {
	case StandardText3D:
		return Tag_DynamicText3D
	case DynamicText3DTag:
		return Tag_DynamicText3D
	case NoText3DTag:
		return Tag_DynamicObject // or any other default value
	default:
		return Tag_DynamicText3D
	}
}
