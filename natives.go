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

type Native interface {
	Streamer_GetTickRate() int32
	Streamer_SetTickRate(tickrate int32)
	Streamer_GetPlayerTickRate(playerid int32) int32
	Streamer_SetPlayerTickRate(playerid int32, tickrate int32)
	Streamer_ToggleChunkStreaming(toggle bool)
	Streamer_IsToggleChunkStreaming() bool
	Streamer_GetChunkTickRate(streamerType StreamerType, playerid int32) int32
	Streamer_SetChunkTickRate(streamerType StreamerType, tickrate int32, playerid int32)
	Streamer_GetChunkSize(streamerType StreamerType) float32
	Streamer_SetChunkSize(streamerType StreamerType, size float32)
	Streamer_GetMaxItems(streamerType StreamerType) int32
	Streamer_SetMaxItems(streamerType StreamerType, maxItems int32)
	Streamer_GetVisibleItems(streamerType StreamerType, playerid int32) int32
	Streamer_SetVisibleItems(streamerType StreamerType, visibleItems int32, playerid int32)
	Streamer_GetRadiusMultiplier(streamerType StreamerType, playerid int32) float32
	Streamer_SetRadiusMultiplier(streamerType StreamerType, multiplier *float32, playerid int32)
	Streamer_GetTypePriority(types []int32, maxtypes int32) int32
	Streamer_SetTypePriority(types []int32, maxtypes int32)
	Streamer_GetCellDistance() float32
	Streamer_SetCellDistance(distance float32)
	Streamer_GetCellSize() float32
	Streamer_SetCellSize(size float32)
	Streamer_ToggleItemStatic(streamerType StreamerType, id StreamerTag, toggle bool)
	Streamer_IsToggleItemStatic(streamerType StreamerType, id StreamerTag) bool
	Streamer_ToggleItemInvAreas(streamerType StreamerType, id StreamerTag, toggle bool)
	Streamer_IsToggleItemInvAreas(streamerType StreamerType, id StreamerTag) bool
	Streamer_ToggleItemCallbacks(streamerType StreamerType, id StreamerTag, toggle bool)
	Streamer_IsToggleItemCallbacks(streamerType StreamerType, id StreamerTag) bool
	Streamer_ToggleErrorCallback(toggle bool)
	Streamer_IsToggleErrorCallback() bool

	// Updates
	Streamer_ProcessActiveItems()
	Streamer_ToggleIdleUpdate(playerid int32, toggle bool)
	Streamer_IsToggleIdleUpdate(playerid int32) bool
	Streamer_ToggleCameraUpdate(playerid int32, toggle bool)
	Streamer_IsToggleCameraUpdate(playerid int32) bool
	Streamer_ToggleItemUpdate(playerid int32, streamerType StreamerType, toggle bool)
	Streamer_IsToggleItemUpdate(playerid int32, streamerType StreamerType) bool
	Streamer_GetLastUpdateTime() float32
	Streamer_Update(playerid int32, streamerType StreamerType)
	Streamer_UpdateEx(playerid int32, x float32, y float32, z float32, worldid int32, interiorid int32, streamerType StreamerType, compensatedtime int32, freezeplayer bool)

	// Data manipulation
	Streamer_GetFloatData(streamerType StreamerType, id StreamerTag, data int32) float32
	Streamer_SetFloatData(streamerType StreamerType, id StreamerTag, data int32, value float32)
	Streamer_GetIntData(streamerType StreamerType, id StreamerTag, data int32) int32
	Streamer_SetIntData(streamerType StreamerType, id StreamerTag, data int32, value int32)
	Streamer_RemoveIntData(streamerType StreamerType, id StreamerTag, data int32)
	Streamer_HasIntData(streamerType StreamerType, id StreamerTag, data int32) bool

	// Array data manipulation
	Streamer_GetArrayData(streamerType StreamerType, id StreamerTag, data int32) []int32
	Streamer_SetArrayData(streamerType StreamerType, id StreamerTag, data int32, src []int32)
	Streamer_IsInArrayData(streamerType StreamerType, id StreamerTag, data int32, value int32) bool
	Streamer_AppendArrayData(streamerType StreamerType, id StreamerTag, data int32, value int32)
	Streamer_RemoveArrayData(streamerType StreamerType, id StreamerTag, data int32, value int32)
	Streamer_HasArrayData(streamerType StreamerType, id StreamerTag, data int32) bool
	Streamer_GetArrayDataLength(streamerType StreamerType, id StreamerTag, data int32) int32
	Streamer_GetUpperBound(streamerType StreamerType) int32

	// Miscellaneous

	// Item manipulation
	Streamer_GetDistanceToItem(pos Vector3, streamerType StreamerType, id StreamerTag, dimensions int32) float32
	Streamer_ToggleItem(playerid int32, streamerType StreamerType, id StreamerTag, toggle bool)
	Streamer_IsToggleItem(playerid int32, streamerType StreamerType, id StreamerTag) bool
	Streamer_ToggleAllItems(playerid int32, streamerType StreamerType, toggle bool, exceptions []int32)
	Streamer_GetItemInternalID(playerid int32, streamerType StreamerType, streamerid StreamerTag) int32
	Streamer_GetItemStreamerID(playerid int32, streamerType StreamerType, internalid int32) StreamerTag
	Streamer_IsItemVisible(playerid int32, streamerType StreamerType, id StreamerTag) bool

	// Bulk operations
	Streamer_DestroyAllVisibleItems(playerid int32, streamerType StreamerType, serverwide bool)
	Streamer_CountVisibleItems(playerid int32, streamerType StreamerType, serverwide bool) int32
	Streamer_DestroyAllItems(streamerType StreamerType, serverwide bool)
	Streamer_CountItems(streamerType StreamerType, serverwide bool) int32

	// Position and offset
	Streamer_GetNearbyItems(pos Vector3, streamerType StreamerType, radius float32, worldid int32) []StreamerTag
	Streamer_GetAllVisibleItems(playerid int32, streamerType StreamerType) []StreamerTag
	Streamer_GetItemPos(streamerType StreamerType, id StreamerTag) Vector3
	Streamer_SetItemPos(streamerType StreamerType, id StreamerTag, pos Vector3)
	Streamer_GetItemOffset(streamerType StreamerType, id StreamerTag) Vector3
	Streamer_SetItemOffset(streamerType StreamerType, id StreamerTag, offset Vector3)

	// Objects
	CreateDynamicObject(modelid int32, pos Vector3, rot Vector3, worldid int32, interiorid int32,
		playerid int32, streamdistance float32, drawdistance float32, areaid AreaID, priority int32) ObjectID
	DestroyDynamicObject(objectid ObjectID)
	IsValidDynamicObject(objectid ObjectID) bool
	GetDynamicObjectPos(objectid ObjectID) Vector3
	SetDynamicObjectPos(objectid ObjectID, pos Vector3)
	GetDynamicObjectRot(objectid ObjectID) Vector3
	SetDynamicObjectRot(objectid ObjectID, rot Vector3)
	GetDynamicObjectNoCameraCol(objectid ObjectID) bool
	SetDynamicObjectNoCameraCol(objectid ObjectID)
	MoveDynamicObject(objectid ObjectID, pos Vector3, speed float32, rot Vector3)
	StopDynamicObject(objectid ObjectID)
	IsDynamicObjectMoving(objectid ObjectID) bool
	AttachCameraToDynamicObject(playerid int32, objectid ObjectID)
	AttachDynamicObjectToObject(objectid ObjectID, attachtoid int32, offset Vector3, rot Vector3, syncrotation bool)
	AttachDynamicObjectToPlayer(objectid ObjectID, playerid int32, offset Vector3, rot Vector3)
	AttachDynamicObjectToVehicle(objectid ObjectID, vehicleid int32, offset Vector3, rot Vector3)
	EditDynamicObject(playerid int32, objectid ObjectID)

	// Material methods
	IsDynamicObjectMaterialUsed(objectid ObjectID, materialindex int32) bool
	RemoveDynamicObjectMaterial(objectid ObjectID, materialindex int32)
	GetDynamicObjectMaterial(objectid ObjectID, materialindex int32) (modelid int32, txdname string, texturename string, materialcolor int32)
	SetDynamicObjectMaterial(objectid ObjectID, materialindex int32, modelid int32, txdname string, texturename string, materialcolor int32)
	IsDynamicObjectMaterialTextUsed(objectid ObjectID, materialindex int32) bool
	RemoveDynamicObjectMaterialText(objectid ObjectID, materialindex int32)
	GetDynamicObjectMaterialText(objectid ObjectID, materialindex int32) (text string, materialsize MaterialSize, fontface string, fontsize int32, bold bool, fontcolor int32, backcolor int32, textalignment int32)
	SetDynamicObjectMaterialText(objectid ObjectID, materialindex int32, text string, materialsize MaterialSize, fontface string, fontsize int32, bold bool, fontcolor int32, backcolor int32, textalignment int32)
	GetPlayerCameraTargetDynObject(playerid int32) ObjectID

	// Pickups
	CreateDynamicPickup(modelid int32, pickuptype int32, pos Vector3, worldid int32, interiorid int32,
		playerid int32, streamdistance float32, areaid AreaID, priority int32) PickupID
	DestroyDynamicPickup(pickupid PickupID)
	IsValidDynamicPickup(pickupid PickupID) bool

	// Checkpoints
	CreateDynamicCP(pos Vector3, size float32, worldid int32, interiorid int32,
		playerid int32, streamdistance float32, areaid AreaID, priority int32) CheckpointID
	DestroyDynamicCP(checkpointid CheckpointID)
	IsValidDynamicCP(checkpointid CheckpointID) bool
	IsPlayerInDynamicCP(playerid int32, checkpointid CheckpointID) bool
	GetPlayerVisibleDynamicCP(playerid int32) CheckpointID

	// Race Checkpoints
	CreateDynamicRaceCP(cptype RaceCheckpointType, pos Vector3, nextpos Vector3, size float32, worldid int32, interiorid int32,
		playerid int32, streamdistance float32, areaid AreaID, priority int32) RaceCheckpointID
	DestroyDynamicRaceCP(checkpointid RaceCheckpointID)
	IsValidDynamicRaceCP(checkpointid RaceCheckpointID) bool
	IsPlayerInDynamicRaceCP(playerid int32, checkpointid RaceCheckpointID) bool
	GetPlayerVisibleDynamicRaceCP(playerid int32) RaceCheckpointID

	// Map Icons
	CreateDynamicMapIcon(pos Vector3, icontype int32, color int32, worldid int32, interiorid int32,
		playerid int32, streamdistance float32, style int32, areaid AreaID, priority int32) MapIconID
	DestroyDynamicMapIcon(iconid MapIconID)
	IsValidDynamicMapIcon(iconid MapIconID) bool

	// 3D Text Labels
	CreateDynamic3DTextLabel(text string, color int32, pos Vector3, drawdistance float32,
		attachedplayer int32, attachedvehicle int32, testlos bool, worldid int32, interiorid int32,
		playerid int32, streamdistance float32, areaid AreaID, priority int32) TextLabelID
	DestroyDynamic3DTextLabel(id TextLabelID)
	IsValidDynamic3DTextLabel(id TextLabelID) bool
	GetDynamic3DTextLabelText(id TextLabelID) string
	UpdateDynamic3DTextLabelText(id TextLabelID, color int32, text string)

	// Areas
	CreateDynamicCircle(center Vector2, size float32, worldid int32, interiorid int32, playerid int32, priority int32) AreaID
	CreateDynamicCylinder(pos Vector2, minz float32, maxz float32, size float32, worldid int32, interiorid int32, playerid int32, priority int32) AreaID
	CreateDynamicSphere(pos Vector3, size float32, worldid int32, interiorid int32, playerid int32, priority int32) AreaID
	CreateDynamicRectangle(min Vector2, max Vector2, worldid int32, interiorid int32, playerid int32, priority int32) AreaID
	CreateDynamicCuboid(min Vector3, max Vector3, worldid int32, interiorid int32, playerid int32, priority int32) AreaID
	CreateDynamicCube(min Vector3, max Vector3, worldid int32, interiorid int32, playerid int32, priority int32) AreaID
	CreateDynamicPolygon(points []Vector2, minz float32, maxz float32, worldid int32, interiorid int32, playerid int32, priority int32) AreaID
	DestroyDynamicArea(areaid AreaID)
	IsValidDynamicArea(areaid AreaID) bool
	GetDynamicAreaType(areaid AreaID) int32
	GetDynamicPolygonPoints(areaid AreaID) []Vector2
	GetDynamicPolygonNumberPoints(areaid AreaID) int32
	IsPlayerInDynamicArea(playerid int32, areaid AreaID, recheck bool) bool
	IsPlayerInAnyDynamicArea(playerid int32, recheck bool) bool
	IsAnyPlayerInDynamicArea(areaid AreaID, recheck bool) bool
	IsAnyPlayerInAnyDynamicArea(recheck bool) bool
	GetPlayerDynamicAreas(playerid int32) []AreaID
	GetPlayerNumberDynamicAreas(playerid int32) int32
	IsPointInDynamicArea(areaid AreaID, point Vector3) bool
	IsPointInAnyDynamicArea(point Vector3) bool
	IsLineInDynamicArea(areaid AreaID, start Vector3, end Vector3) bool
	IsLineInAnyDynamicArea(start Vector3, end Vector3) bool
	GetDynamicAreasForPoint(point Vector3) []AreaID
	GetNumberDynamicAreasForPoint(point Vector3) int32
	GetDynamicAreasForLine(start Vector3, end Vector3) []AreaID
	GetNumberDynamicAreasForLine(start Vector3, end Vector3) int32
	AttachDynamicAreaToObject(areaid AreaID, objectid ObjectID, objecttype int32, playerid int32, offset Vector3)
	AttachDynamicAreaToPlayer(areaid AreaID, playerid int32, offset Vector3)
	AttachDynamicAreaToVehicle(areaid AreaID, vehicleid int32, offset Vector3)
	ToggleDynAreaSpectateMode(areaid AreaID, toggle bool)
	IsToggleDynAreaSpectateMode(areaid AreaID) bool

	// Actors
	CreateDynamicActor(modelid int32, pos Vector3, rotation float32, invulnerable bool, health float32,
		worldid int32, interiorid int32, playerid int32, streamdistance float32, areaid AreaID, priority int32) ActorID
	DestroyDynamicActor(actorid ActorID)
	IsValidDynamicActor(actorid ActorID) bool
	IsDynamicActorStreamedIn(actorid ActorID, forplayerid int32) bool
	GetDynamicActorVirtualWorld(actorid ActorID) int32
	SetDynamicActorVirtualWorld(actorid ActorID, vworld int32)
	GetDynamicActorAnimation(actorid ActorID) (animlib string, animname string, fdelta float32, loop bool, lockx bool, locky bool, freeze bool, time int32)
	ApplyDynamicActorAnimation(actorid ActorID, animlib string, animname string, fdelta float32, loop bool, lockx bool, locky bool, freeze bool, time int32)
	ClearDynamicActorAnimations(actorid ActorID)
	GetDynamicActorFacingAngle(actorid ActorID) float32
	SetDynamicActorFacingAngle(actorid ActorID, angle float32)
	GetDynamicActorPos(actorid ActorID) Vector3
	SetDynamicActorPos(actorid ActorID, pos Vector3)
	GetDynamicActorHealth(actorid ActorID) float32
	SetDynamicActorHealth(actorid ActorID, health float32)
	SetDynamicActorInvulnerable(actorid ActorID, invulnerable bool)
	IsDynamicActorInvulnerable(actorid ActorID) bool
	GetPlayerTargetDynamicActor(playerid int32) ActorID
	GetPlayerCameraTargetDynActor(playerid int32) ActorID

	// Extended creation functions
	CreateDynamicObjectEx(modelid int32, pos Vector3, rot Vector3, streamdistance float32, drawdistance float32,
		worlds []int32, interiors []int32, players []int32, areas []AreaID, priority int32) ObjectID
	CreateDynamicPickupEx(modelid int32, pickuptype int32, pos Vector3, streamdistance float32,
		worlds []int32, interiors []int32, players []int32, areas []AreaID, priority int32) PickupID
	CreateDynamicCPEx(pos Vector3, size float32, streamdistance float32,
		worlds []int32, interiors []int32, players []int32, areas []AreaID, priority int32) CheckpointID
	CreateDynamicRaceCPEx(cptype int32, pos Vector3, nextpos Vector3, size float32, streamdistance float32,
		worlds []int32, interiors []int32, players []int32, areas []AreaID, priority int32) RaceCheckpointID
	CreateDynamicMapIconEx(pos Vector3, icontype int32, color int32, style int32, streamdistance float32,
		worlds []int32, interiors []int32, players []int32, areas []AreaID, priority int32) MapIconID
	CreateDynamic3DTextLabelEx(text string, color int32, pos Vector3, drawdistance float32,
		attachedplayer int32, attachedvehicle int32, testlos bool, streamdistance float32,
		worlds []int32, interiors []int32, players []int32, areas []AreaID, priority int32) TextLabelID
	CreateDynamicCircleEx(center Vector2, size float32,
		worlds []int32, interiors []int32, players []int32, priority int32) AreaID
	CreateDynamicCylinderEx(pos Vector2, minz float32, maxz float32, size float32,
		worlds []int32, interiors []int32, players []int32, priority int32) AreaID
	CreateDynamicSphereEx(pos Vector3, size float32,
		worlds []int32, interiors []int32, players []int32, priority int32) AreaID
	CreateDynamicRectangleEx(min Vector2, max Vector2,
		worlds []int32, interiors []int32, players []int32, priority int32) AreaID
	CreateDynamicCuboidEx(min Vector3, max Vector3,
		worlds []int32, interiors []int32, players []int32, priority int32) AreaID
	CreateDynamicCubeEx(min Vector3, max Vector3,
		worlds []int32, interiors []int32, players []int32, priority int32) AreaID
	CreateDynamicPolygonEx(points []Vector2, minz float32, maxz float32,
		worlds []int32, interiors []int32, players []int32, priority int32) AreaID
	CreateDynamicActorEx(modelid int32, pos Vector3, rotation float32, invulnerable bool, health float32, streamdistance float32,
		worlds []int32, interiors []int32, players []int32, areas []AreaID, priority int32) ActorID
}

type Callbacks interface {
	// Object callbacks
	OnDynamicObjectMoved(objectid ObjectID)
	OnPlayerEditDynamicObject(playerid int32, objectid ObjectID, response int32, pos Vector3, rot Vector3)
	OnPlayerSelectDynamicObject(playerid int32, objectid ObjectID, modelid int32, pos Vector3)
	OnPlayerShootDynamicObject(playerid int32, weaponid int32, objectid ObjectID, pos Vector3)

	// Pickup callbacks
	OnPlayerPickUpDynamicPickup(playerid int32, pickupid PickupID)

	// Checkpoint callbacks
	OnPlayerEnterDynamicCP(playerid int32, checkpointid CheckpointID)
	OnPlayerLeaveDynamicCP(playerid int32, checkpointid CheckpointID)
	OnPlayerEnterDynamicRaceCP(playerid int32, checkpointid RaceCheckpointID)
	OnPlayerLeaveDynamicRaceCP(playerid int32, checkpointid RaceCheckpointID)

	// Area callbacks
	OnPlayerEnterDynamicArea(playerid int32, areaid AreaID)
	OnPlayerLeaveDynamicArea(playerid int32, areaid AreaID)

	// Actor callbacks
	OnPlayerGiveDamageDynamicActor(playerid int32, actorid ActorID, amount float32, weaponid int32, bodypart int32)
	OnDynamicActorStreamIn(actorid ActorID, forplayerid int32)
	OnDynamicActorStreamOut(actorid ActorID, forplayerid int32)

	// General streaming callbacks
	Streamer_OnItemStreamIn(streamType StreamerType, id StreamerTag, forplayerid int32)
	Streamer_OnItemStreamOut(streamType StreamerType, id StreamerTag, forplayerid int32)
	Streamer_OnPluginError(err string)
}
