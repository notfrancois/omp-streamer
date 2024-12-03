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

package streamer

import "github.com/notfrancois/omp-streamer/pkg/types"

type Native interface {
	Streamer_GetTickRate() int32
	Streamer_SetTickRate(tickrate int32)
	Streamer_GetPlayerTickRate(playerid int32) int32
	Streamer_SetPlayerTickRate(playerid int32, tickrate int32)
	Streamer_ToggleChunkStreaming(toggle bool)
	Streamer_IsToggleChunkStreaming() bool
	Streamer_GetChunkTickRate(streamerType types.StreamerType, playerid int32) int32
	Streamer_SetChunkTickRate(streamerType types.StreamerType, tickrate int32, playerid int32)
	Streamer_GetChunkSize(streamerType types.StreamerType) float32
	Streamer_SetChunkSize(streamerType types.StreamerType, size float32)
	Streamer_GetMaxItems(streamerType types.StreamerType) int32
	Streamer_SetMaxItems(streamerType types.StreamerType, maxItems int32)
	Streamer_GetVisibleItems(streamerType types.StreamerType, playerid int32) int32
	Streamer_SetVisibleItems(streamerType types.StreamerType, visibleItems int32, playerid int32)
	Streamer_GetRadiusMultiplier(streamerType types.StreamerType, playerid int32) float32
	Streamer_SetRadiusMultiplier(streamerType types.StreamerType, multiplier *float32, playerid int32)
	Streamer_GetTypePriority(types []int32, maxtypes int32) int32
	Streamer_SetTypePriority(types []int32, maxtypes int32)
	Streamer_GetCellDistance() float32
	Streamer_SetCellDistance(distance float32)
	Streamer_GetCellSize() float32
	Streamer_SetCellSize(size float32)
	Streamer_ToggleItemStatic(streamerType types.StreamerType, id types.StreamerTag, toggle bool)
	Streamer_IsToggleItemStatic(streamerType types.StreamerType, id types.StreamerTag) bool
	Streamer_ToggleItemInvAreas(streamerType types.StreamerType, id types.StreamerTag, toggle bool)
	Streamer_IsToggleItemInvAreas(streamerType types.StreamerType, id types.StreamerTag) bool
	Streamer_ToggleItemCallbacks(streamerType types.StreamerType, id types.StreamerTag, toggle bool)
	Streamer_IsToggleItemCallbacks(streamerType types.StreamerType, id types.StreamerTag) bool
	Streamer_ToggleErrorCallback(toggle bool)
	Streamer_IsToggleErrorCallback() bool

	// Updates
	Streamer_ProcessActiveItems()
	Streamer_ToggleIdleUpdate(playerid int32, toggle bool)
	Streamer_IsToggleIdleUpdate(playerid int32) bool
	Streamer_ToggleCameraUpdate(playerid int32, toggle bool)
	Streamer_IsToggleCameraUpdate(playerid int32) bool
	Streamer_ToggleItemUpdate(playerid int32, streamerType types.StreamerType, toggle bool)
	Streamer_IsToggleItemUpdate(playerid int32, streamerType types.StreamerType) bool
	Streamer_GetLastUpdateTime() float32
	Streamer_Update(playerid int32, streamerType types.StreamerType)
	Streamer_UpdateEx(playerid int32, x float32, y float32, z float32, worldid int32, interiorid int32, streamerType types.StreamerType, compensatedtime int32, freezeplayer bool)

	// Data manipulation
	Streamer_GetFloatData(streamerType types.StreamerType, id types.StreamerTag, data int32) float32
	Streamer_SetFloatData(streamerType types.StreamerType, id types.StreamerTag, data int32, value float32)
	Streamer_GetIntData(streamerType types.StreamerType, id types.StreamerTag, data int32) int32
	Streamer_SetIntData(streamerType types.StreamerType, id types.StreamerTag, data int32, value int32)
	Streamer_RemoveIntData(streamerType types.StreamerType, id types.StreamerTag, data int32)
	Streamer_HasIntData(streamerType types.StreamerType, id types.StreamerTag, data int32) bool

	// Array data manipulation
	Streamer_GetArrayData(streamerType types.StreamerType, id types.StreamerTag, data int32) []int32
	Streamer_SetArrayData(streamerType types.StreamerType, id types.StreamerTag, data int32, src []int32)
	Streamer_IsInArrayData(streamerType types.StreamerType, id types.StreamerTag, data int32, value int32) bool
	Streamer_AppendArrayData(streamerType types.StreamerType, id types.StreamerTag, data int32, value int32)
	Streamer_RemoveArrayData(streamerType types.StreamerType, id types.StreamerTag, data int32, value int32)
	Streamer_HasArrayData(streamerType types.StreamerType, id types.StreamerTag, data int32) bool
	Streamer_GetArrayDataLength(streamerType types.StreamerType, id types.StreamerTag, data int32) int32
	Streamer_GetUpperBound(streamerType types.StreamerType) int32

	// Miscellaneous

	// Item manipulation
	Streamer_GetDistanceToItem(pos types.Vector3, streamerType types.StreamerType, id types.StreamerTag, dimensions int32) float32
	Streamer_ToggleItem(playerid int32, streamerType types.StreamerType, id types.StreamerTag, toggle bool)
	Streamer_IsToggleItem(playerid int32, streamerType types.StreamerType, id types.StreamerTag) bool
	Streamer_ToggleAllItems(playerid int32, streamerType types.StreamerType, toggle bool, exceptions []int32)
	Streamer_GetItemInternalID(playerid int32, streamerType types.StreamerType, streamerid types.StreamerTag) int32
	Streamer_GetItemStreamerID(playerid int32, streamerType types.StreamerType, internalid int32) types.StreamerTag
	Streamer_IsItemVisible(playerid int32, streamerType types.StreamerType, id types.StreamerTag) bool

	// Bulk operations
	Streamer_DestroyAllVisibleItems(playerid int32, streamerType types.StreamerType, serverwide bool)
	Streamer_CountVisibleItems(playerid int32, streamerType types.StreamerType, serverwide bool) int32
	Streamer_DestroyAllItems(streamerType types.StreamerType, serverwide bool)
	Streamer_CountItems(streamerType types.StreamerType, serverwide bool) int32

	// Position and offset
	Streamer_GetNearbyItems(pos types.Vector3, streamerType types.StreamerType, radius float32, worldid int32) []types.StreamerTag
	Streamer_GetAllVisibleItems(playerid int32, streamerType types.StreamerType) []types.StreamerTag
	Streamer_GetItemPos(streamerType types.StreamerType, id types.StreamerTag) types.Vector3
	Streamer_SetItemPos(streamerType types.StreamerType, id types.StreamerTag, pos types.Vector3)
	Streamer_GetItemOffset(streamerType types.StreamerType, id types.StreamerTag) types.Vector3
	Streamer_SetItemOffset(streamerType types.StreamerType, id types.StreamerTag, offset types.Vector3)

	// Objects
	CreateDynamicObject(modelid int32, pos types.Vector3, rot types.Vector3, worldid int32, interiorid int32,
		playerid int32, streamdistance float32, drawdistance float32, areaid types.AreaID, priority int32) types.ObjectID
	DestroyDynamicObject(objectid types.ObjectID)
	IsValidDynamicObject(objectid types.ObjectID) bool
	GetDynamicObjectPos(objectid types.ObjectID) types.Vector3
	SetDynamicObjectPos(objectid types.ObjectID, pos types.Vector3)
	GetDynamicObjectRot(objectid types.ObjectID) types.Vector3
	SetDynamicObjectRot(objectid types.ObjectID, rot types.Vector3)
	GetDynamicObjectNoCameraCol(objectid types.ObjectID) bool
	SetDynamicObjectNoCameraCol(objectid types.ObjectID)
	MoveDynamicObject(objectid types.ObjectID, pos types.Vector3, speed float32, rot types.Vector3)
	StopDynamicObject(objectid types.ObjectID)
	IsDynamicObjectMoving(objectid types.ObjectID) bool
	AttachCameraToDynamicObject(playerid int32, objectid types.ObjectID)
	AttachDynamicObjectToObject(objectid types.ObjectID, attachtoid int32, offset types.Vector3, rot types.Vector3, syncrotation bool)
	AttachDynamicObjectToPlayer(objectid types.ObjectID, playerid int32, offset types.Vector3, rot types.Vector3)
	AttachDynamicObjectToVehicle(objectid types.ObjectID, vehicleid int32, offset types.Vector3, rot types.Vector3)
	EditDynamicObject(playerid int32, objectid types.ObjectID)

	// Material methods
	IsDynamicObjectMaterialUsed(objectid types.ObjectID, materialindex int32) bool
	RemoveDynamicObjectMaterial(objectid types.ObjectID, materialindex int32)
	GetDynamicObjectMaterial(objectid types.ObjectID, materialindex int32) (modelid int32, txdname string, texturename string, materialcolor int32)
	SetDynamicObjectMaterial(objectid types.ObjectID, materialindex int32, modelid int32, txdname string, texturename string, materialcolor int32)
	IsDynamicObjectMaterialTextUsed(objectid types.ObjectID, materialindex int32) bool
	RemoveDynamicObjectMaterialText(objectid types.ObjectID, materialindex int32)
	GetDynamicObjectMaterialText(objectid types.ObjectID, materialindex int32) (text string, materialsize types.MaterialSize, fontface string, fontsize int32, bold bool, fontcolor int32, backcolor int32, textalignment int32)
	SetDynamicObjectMaterialText(objectid types.ObjectID, materialindex int32, text string, materialsize types.MaterialSize, fontface string, fontsize int32, bold bool, fontcolor int32, backcolor int32, textalignment int32)
	GetPlayerCameraTargetDynObject(playerid int32) types.ObjectID

	// Pickups
	CreateDynamicPickup(modelid int32, pickuptype int32, pos types.Vector3, worldid int32, interiorid int32,
		playerid int32, streamdistance float32, areaid types.AreaID, priority int32) types.PickupID
	DestroyDynamicPickup(pickupid types.PickupID)
	IsValidDynamicPickup(pickupid types.PickupID) bool

	// Checkpoints
	CreateDynamicCP(pos types.Vector3, size float32, worldid int32, interiorid int32,
		playerid int32, streamdistance float32, areaid types.AreaID, priority int32) types.CheckpointID
	DestroyDynamicCP(checkpointid types.CheckpointID)
	IsValidDynamicCP(checkpointid types.CheckpointID) bool
	IsPlayerInDynamicCP(playerid int32, checkpointid types.CheckpointID) bool
	GetPlayerVisibleDynamicCP(playerid int32) types.CheckpointID

	// Race Checkpoints
	CreateDynamicRaceCP(cptype types.RaceCheckpointType, pos types.Vector3, nextpos types.Vector3, size float32, worldid int32, interiorid int32,
		playerid int32, streamdistance float32, areaid types.AreaID, priority int32) types.RaceCheckpointID
	DestroyDynamicRaceCP(checkpointid types.RaceCheckpointID)
	IsValidDynamicRaceCP(checkpointid types.RaceCheckpointID) bool
	IsPlayerInDynamicRaceCP(playerid int32, checkpointid types.RaceCheckpointID) bool
	GetPlayerVisibleDynamicRaceCP(playerid int32) types.RaceCheckpointID

	// Map Icons
	CreateDynamicMapIcon(pos types.Vector3, icontype int32, color int32, worldid int32, interiorid int32,
		playerid int32, streamdistance float32, style int32, areaid types.AreaID, priority int32) types.MapIconID
	DestroyDynamicMapIcon(iconid types.MapIconID)
	IsValidDynamicMapIcon(iconid types.MapIconID) bool

	// 3D Text Labels
	CreateDynamic3DTextLabel(text string, color int32, pos types.Vector3, drawdistance float32,
		attachedplayer int32, attachedvehicle int32, testlos bool, worldid int32, interiorid int32,
		playerid int32, streamdistance float32, areaid types.AreaID, priority int32) types.TextLabelID
	DestroyDynamic3DTextLabel(id types.TextLabelID)
	IsValidDynamic3DTextLabel(id types.TextLabelID) bool
	GetDynamic3DTextLabelText(id types.TextLabelID) string
	UpdateDynamic3DTextLabelText(id types.TextLabelID, color int32, text string)

	// Areas
	CreateDynamicCircle(center types.Vector2, size float32, worldid int32, interiorid int32, playerid int32, priority int32) types.AreaID
	CreateDynamicCylinder(pos types.Vector2, minz float32, maxz float32, size float32, worldid int32, interiorid int32, playerid int32, priority int32) types.AreaID
	CreateDynamicSphere(pos types.Vector3, size float32, worldid int32, interiorid int32, playerid int32, priority int32) types.AreaID
	CreateDynamicRectangle(min types.Vector2, max types.Vector2, worldid int32, interiorid int32, playerid int32, priority int32) types.AreaID
	CreateDynamicCuboid(min types.Vector3, max types.Vector3, worldid int32, interiorid int32, playerid int32, priority int32) types.AreaID
	CreateDynamicCube(min types.Vector3, max types.Vector3, worldid int32, interiorid int32, playerid int32, priority int32) types.AreaID
	CreateDynamicPolygon(points []types.Vector2, minz float32, maxz float32, worldid int32, interiorid int32, playerid int32, priority int32) types.AreaID
	DestroyDynamicArea(areaid types.AreaID)
	IsValidDynamicArea(areaid types.AreaID) bool
	GetDynamicAreaType(areaid types.AreaID) int32
	GetDynamicPolygonPoints(areaid types.AreaID) []types.Vector2
	GetDynamicPolygonNumberPoints(areaid types.AreaID) int32
	IsPlayerInDynamicArea(playerid int32, areaid types.AreaID, recheck bool) bool
	IsPlayerInAnyDynamicArea(playerid int32, recheck bool) bool
	IsAnyPlayerInDynamicArea(areaid types.AreaID, recheck bool) bool
	IsAnyPlayerInAnyDynamicArea(recheck bool) bool
	GetPlayerDynamicAreas(playerid int32) []types.AreaID
	GetPlayerNumberDynamicAreas(playerid int32) int32
	IsPointInDynamicArea(areaid types.AreaID, point types.Vector3) bool
	IsPointInAnyDynamicArea(point types.Vector3) bool
	IsLineInDynamicArea(areaid types.AreaID, start types.Vector3, end types.Vector3) bool
	IsLineInAnyDynamicArea(start types.Vector3, end types.Vector3) bool
	GetDynamicAreasForPoint(point types.Vector3) []types.AreaID
	GetNumberDynamicAreasForPoint(point types.Vector3) int32
	GetDynamicAreasForLine(start types.Vector3, end types.Vector3) []types.AreaID
	GetNumberDynamicAreasForLine(start types.Vector3, end types.Vector3) int32
	AttachDynamicAreaToObject(areaid types.AreaID, objectid types.ObjectID, objecttype int32, playerid int32, offset types.Vector3)
	AttachDynamicAreaToPlayer(areaid types.AreaID, playerid int32, offset types.Vector3)
	AttachDynamicAreaToVehicle(areaid types.AreaID, vehicleid int32, offset types.Vector3)
	ToggleDynAreaSpectateMode(areaid types.AreaID, toggle bool)
	IsToggleDynAreaSpectateMode(areaid types.AreaID) bool

	// Actors
	CreateDynamicActor(modelid int32, pos types.Vector3, rotation float32, invulnerable bool, health float32,
		worldid int32, interiorid int32, playerid int32, streamdistance float32, areaid types.AreaID, priority int32) types.ActorID
	DestroyDynamicActor(actorid types.ActorID)
	IsValidDynamicActor(actorid types.ActorID) bool
	IsDynamicActorStreamedIn(actorid types.ActorID, forplayerid int32) bool
	GetDynamicActorVirtualWorld(actorid types.ActorID) int32
	SetDynamicActorVirtualWorld(actorid types.ActorID, vworld int32)
	GetDynamicActorAnimation(actorid types.ActorID) (animlib string, animname string, fdelta float32, loop bool, lockx bool, locky bool, freeze bool, time int32)
	ApplyDynamicActorAnimation(actorid types.ActorID, animlib string, animname string, fdelta float32, loop bool, lockx bool, locky bool, freeze bool, time int32)
	ClearDynamicActorAnimations(actorid types.ActorID)
	GetDynamicActorFacingAngle(actorid types.ActorID) float32
	SetDynamicActorFacingAngle(actorid types.ActorID, angle float32)
	GetDynamicActorPos(actorid types.ActorID) types.Vector3
	SetDynamicActorPos(actorid types.ActorID, pos types.Vector3)
	GetDynamicActorHealth(actorid types.ActorID) float32
	SetDynamicActorHealth(actorid types.ActorID, health float32)
	SetDynamicActorInvulnerable(actorid types.ActorID, invulnerable bool)
	IsDynamicActorInvulnerable(actorid types.ActorID) bool
	GetPlayerTargetDynamicActor(playerid int32) types.ActorID
	GetPlayerCameraTargetDynActor(playerid int32) types.ActorID

	// Extended creation functions
	CreateDynamicObjectEx(modelid int32, pos types.Vector3, rot types.Vector3, streamdistance float32, drawdistance float32,
		worlds []int32, interiors []int32, players []int32, areas []types.AreaID, priority int32) types.ObjectID
	CreateDynamicPickupEx(modelid int32, pickuptype int32, pos types.Vector3, streamdistance float32,
		worlds []int32, interiors []int32, players []int32, areas []types.AreaID, priority int32) types.PickupID
	CreateDynamicCPEx(pos types.Vector3, size float32, streamdistance float32,
		worlds []int32, interiors []int32, players []int32, areas []types.AreaID, priority int32) types.CheckpointID
	CreateDynamicRaceCPEx(cptype int32, pos types.Vector3, nextpos types.Vector3, size float32, streamdistance float32,
		worlds []int32, interiors []int32, players []int32, areas []types.AreaID, priority int32) types.RaceCheckpointID
	CreateDynamicMapIconEx(pos types.Vector3, icontype int32, color int32, style int32, streamdistance float32,
		worlds []int32, interiors []int32, players []int32, areas []types.AreaID, priority int32) types.MapIconID
	CreateDynamic3DTextLabelEx(text string, color int32, pos types.Vector3, drawdistance float32,
		attachedplayer int32, attachedvehicle int32, testlos bool, streamdistance float32,
		worlds []int32, interiors []int32, players []int32, areas []types.AreaID, priority int32) types.TextLabelID
	CreateDynamicCircleEx(center types.Vector2, size float32,
		worlds []int32, interiors []int32, players []int32, priority int32) types.AreaID
	CreateDynamicCylinderEx(pos types.Vector2, minz float32, maxz float32, size float32,
		worlds []int32, interiors []int32, players []int32, priority int32) types.AreaID
	CreateDynamicSphereEx(pos types.Vector3, size float32,
		worlds []int32, interiors []int32, players []int32, priority int32) types.AreaID
	CreateDynamicRectangleEx(min types.Vector2, max types.Vector2,
		worlds []int32, interiors []int32, players []int32, priority int32) types.AreaID
	CreateDynamicCuboidEx(min types.Vector3, max types.Vector3,
		worlds []int32, interiors []int32, players []int32, priority int32) types.AreaID
	CreateDynamicCubeEx(min types.Vector3, max types.Vector3,
		worlds []int32, interiors []int32, players []int32, priority int32) types.AreaID
	CreateDynamicPolygonEx(points []types.Vector2, minz float32, maxz float32,
		worlds []int32, interiors []int32, players []int32, priority int32) types.AreaID
	CreateDynamicActorEx(modelid int32, pos types.Vector3, rotation float32, invulnerable bool, health float32, streamdistance float32,
		worlds []int32, interiors []int32, players []int32, areas []types.AreaID, priority int32) types.ActorID
}

type Callbacks interface {
	// Object callbacks
	OnDynamicObjectMoved(objectid types.ObjectID)
	OnPlayerEditDynamicObject(playerid int32, objectid types.ObjectID, response int32, pos types.Vector3, rot types.Vector3)
	OnPlayerSelectDynamicObject(playerid int32, objectid types.ObjectID, modelid int32, pos types.Vector3)
	OnPlayerShootDynamicObject(playerid int32, weaponid int32, objectid types.ObjectID, pos types.Vector3)

	// Pickup callbacks
	OnPlayerPickUpDynamicPickup(playerid int32, pickupid types.PickupID)

	// Checkpoint callbacks
	OnPlayerEnterDynamicCP(playerid int32, checkpointid types.CheckpointID)
	OnPlayerLeaveDynamicCP(playerid int32, checkpointid types.CheckpointID)
	OnPlayerEnterDynamicRaceCP(playerid int32, checkpointid types.RaceCheckpointID)
	OnPlayerLeaveDynamicRaceCP(playerid int32, checkpointid types.RaceCheckpointID)

	// Area callbacks
	OnPlayerEnterDynamicArea(playerid int32, areaid types.AreaID)
	OnPlayerLeaveDynamicArea(playerid int32, areaid types.AreaID)

	// Actor callbacks
	OnPlayerGiveDamageDynamicActor(playerid int32, actorid types.ActorID, amount float32, weaponid int32, bodypart int32)
	OnDynamicActorStreamIn(actorid types.ActorID, forplayerid int32)
	OnDynamicActorStreamOut(actorid types.ActorID, forplayerid int32)

	// General streaming callbacks
	Streamer_OnItemStreamIn(streamType types.StreamerType, id types.StreamerTag, forplayerid int32)
	Streamer_OnItemStreamOut(streamType types.StreamerType, id types.StreamerTag, forplayerid int32)
	Streamer_OnPluginError(err string)
}
