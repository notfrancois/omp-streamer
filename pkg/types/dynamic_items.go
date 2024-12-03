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

// ItemType represents the type of dynamic item.
type ItemType int32

const (
	ItemTypeActor ItemType = iota
	ItemTypeArea
	ItemTypeCheckpoint
	ItemTypeObject
	ItemTypeMapIcon
	ItemTypePickup
	ItemTypeRaceCheckpoint
	ItemTypeTextLabel
)

// DynamicActor represents a dynamic actor.
type DynamicActor struct {
	ID           int32
	ModelID      int32
	Position     Vector3
	Rotation     float32
	Health       float32
	Invulnerable bool
	Virtual      bool
}

// GetID implements the DynamicItem interface.
func (a *DynamicActor) GetID() int32 {
	return a.ID
}

// GetPosition implements the DynamicItem interface.
func (a *DynamicActor) GetPosition() Vector3 {
	return a.Position
}

// GetType implements the DynamicItem interface.
func (a *DynamicActor) GetType() ItemType {
	return ItemTypeActor
}

// DynamicArea represents a dynamic area.
type DynamicArea struct {
	ID       int32
	Position Vector3
	Type     AreaType
	Size     Vector3
	Height   float32
	Virtual  bool
}

// DynamicCheckpoint represents a dynamic checkpoint.
type DynamicCheckpoint struct {
	ID       int32
	Position Vector3
	Size     float32
	Virtual  bool
}

// DynamicMapIcon represents a dynamic map icon.
type DynamicMapIcon struct {
	ID       int32
	Position Vector3
	Type     int32
	Color    uint32
	Style    int32
	Virtual  bool
}

// DynamicPickup represents a dynamic pickup.
type DynamicPickup struct {
	ID       int32
	ModelID  int32
	Type     int32
	Position Vector3
	Virtual  bool
}

// DynamicRaceCheckpoint represents a dynamic race checkpoint.
type DynamicRaceCheckpoint struct {
	ID       int32
	Type     int32
	Position Vector3
	NextPos  Vector3
	Size     float32
	Virtual  bool
}

// DynamicTextLabel represents a dynamic 3D text label.
type DynamicTextLabel struct {
	ID           int32
	Text         string
	Color        uint32
	Position     Vector3
	DrawDistance float32
	TestLOS      bool
	Virtual      bool
	Attached     bool
}

// AreaType represents the type of area.
type AreaType int32

const (
	AreaTypeCircle AreaType = iota
	AreaTypeCylinder
	AreaTypeSphere
	AreaTypeRectangle
	AreaTypeCuboid
	AreaTypePolygon
)

func (a *DynamicArea) GetID() int32         { return a.ID }
func (a *DynamicArea) GetPosition() Vector3 { return a.Position }
func (a *DynamicArea) GetType() ItemType    { return ItemTypeArea }

func (c *DynamicCheckpoint) GetID() int32         { return c.ID }
func (c *DynamicCheckpoint) GetPosition() Vector3 { return c.Position }
func (c *DynamicCheckpoint) GetType() ItemType    { return ItemTypeCheckpoint }

func (m *DynamicMapIcon) GetID() int32         { return m.ID }
func (m *DynamicMapIcon) GetPosition() Vector3 { return m.Position }
func (m *DynamicMapIcon) GetType() ItemType    { return ItemTypeMapIcon }

func (p *DynamicPickup) GetID() int32         { return p.ID }
func (p *DynamicPickup) GetPosition() Vector3 { return p.Position }
func (p *DynamicPickup) GetType() ItemType    { return ItemTypePickup }

func (r *DynamicRaceCheckpoint) GetID() int32         { return r.ID }
func (r *DynamicRaceCheckpoint) GetPosition() Vector3 { return r.Position }
func (r *DynamicRaceCheckpoint) GetType() ItemType    { return ItemTypeRaceCheckpoint }

func (t *DynamicTextLabel) GetID() int32         { return t.ID }
func (t *DynamicTextLabel) GetPosition() Vector3 { return t.Position }
func (t *DynamicTextLabel) GetType() ItemType    { return ItemTypeTextLabel }
