package streamer

// #cgo CFLAGS: -I${SRCDIR}/modules/streamer/src
// #cgo LDFLAGS: -L${SRCDIR}/modules/streamer/lib -lstreamer
//
// #include <stdlib.h>
// #include "modules/streamer/src/streamer_wrapper.h"
import "C"
import "unsafe"

// NativeImpl implementa la interfaz Native llamando a funciones C++ nativas
type NativeImpl struct{}

func (n *NativeImpl) Streamer_GetChunkTickRate(streamerType StreamerType, playerid int32) int32 {
	return int32(C.Streamer_GetChunkTickRate(C.StreamerType(streamerType), C.int(playerid)))
}

func (n *NativeImpl) Streamer_SetChunkTickRate(streamerType StreamerType, tickrate int32, playerid int32) {
	C.Streamer_SetChunkTickRate(C.StreamerType(streamerType), C.int(tickrate), C.int(playerid))
}

func (n *NativeImpl) Streamer_GetChunkSize(streamerType StreamerType) float32 {
	return float32(C.Streamer_GetChunkSize(C.StreamerType(streamerType)))
}

func (n *NativeImpl) Streamer_SetChunkSize(streamerType StreamerType, size float32) {
	C.Streamer_SetChunkSize(C.StreamerType(streamerType), C.float(size))
}

func (n *NativeImpl) Streamer_GetMaxItems(streamerType StreamerType) int32 {
	return int32(C.Streamer_GetMaxItems(C.StreamerType(streamerType)))
}

func (n *NativeImpl) Streamer_SetMaxItems(streamerType StreamerType, maxItems int32) {
	C.Streamer_SetMaxItems(C.StreamerType(streamerType), C.int(maxItems))
}

func (n *NativeImpl) Streamer_GetVisibleItems(streamerType StreamerType, playerid int32) int32 {
	return int32(C.Streamer_GetVisibleItems(C.StreamerType(streamerType), C.int(playerid)))
}

func (n *NativeImpl) Streamer_SetVisibleItems(streamerType StreamerType, visibleItems int32, playerid int32) {
	C.Streamer_SetVisibleItems(C.StreamerType(streamerType), C.int(visibleItems), C.int(playerid))
}

func (n *NativeImpl) Streamer_GetRadiusMultiplier(streamerType StreamerType, playerid int32) float32 {
	return float32(C.Streamer_GetRadiusMultiplier(C.StreamerType(streamerType), C.int(playerid)))
}

func (n *NativeImpl) Streamer_SetRadiusMultiplier(streamerType StreamerType, multiplier *float32, playerid int32) {
	C.Streamer_SetRadiusMultiplier(C.StreamerType(streamerType), (*C.float)(multiplier), C.int(playerid))
}

func (n *NativeImpl) Streamer_GetTypePriority(types []int32, maxtypes int32) int32 {
	//return int32(C.Streamer_GetTypePriority(C.int(types[0]), C.int(maxtypes)))
	panic("implement me")
}

func (n *NativeImpl) Streamer_SetTypePriority(types []int32, maxtypes int32) {
	//C.Streamer_SetTypePriority(C.int(types[0]), C.int(maxtypes))
}

func (n *NativeImpl) Streamer_GetCellDistance() float32 {
	return float32(C.Streamer_GetCellDistance())
}

func (n *NativeImpl) Streamer_SetCellDistance(distance float32) {
	C.Streamer_SetCellDistance(C.float(distance))
}

func (n *NativeImpl) Streamer_GetCellSize() float32 {
	return float32(C.Streamer_GetCellSize())
}

func (n *NativeImpl) Streamer_SetCellSize(size float32) {
	C.Streamer_SetCellSize(C.float(size))
}

func (n *NativeImpl) Streamer_ToggleItemStatic(streamerType StreamerType, id StreamerTag, toggle bool) {
	toggleVal := 0
	if toggle {
		toggleVal = 1
	}
	C.Streamer_ToggleItemStatic(C.StreamerType(streamerType), C.StreamerTag(id), C.int(toggleVal))
}

func (n *NativeImpl) Streamer_IsToggleItemStatic(streamerType StreamerType, id StreamerTag) bool {
	panic("implement me")
}

func (n *NativeImpl) Streamer_ToggleItemInvAreas(streamerType StreamerType, id StreamerTag, toggle bool) {
	toggleVal := 0
	if toggle {
		toggleVal = 1
	}
	C.Streamer_ToggleItemInvAreas(C.StreamerType(streamerType), C.StreamerTag(id), C.int(toggleVal))
}

func (n *NativeImpl) Streamer_IsToggleItemInvAreas(streamerType StreamerType, id StreamerTag) bool {
	panic("implement me")
}

func (n *NativeImpl) Streamer_ToggleItemCallbacks(streamerType StreamerType, id StreamerTag, toggle bool) {
	toggleVal := 0
	if toggle {
		toggleVal = 1
	}
	C.Streamer_ToggleItemCallbacks(C.StreamerType(streamerType), C.StreamerTag(id), C.int(toggleVal))
}

func (n *NativeImpl) Streamer_IsToggleItemCallbacks(streamerType StreamerType, id StreamerTag) bool {
	panic("implement me")
}

func (n *NativeImpl) Streamer_ToggleErrorCallback(toggle bool) {
	toggleVal := 0
	if toggle {
		toggleVal = 1
	}
	C.Streamer_ToggleErrorCallback(C.int(toggleVal))
}

func (n *NativeImpl) Streamer_IsToggleErrorCallback() bool {
	panic("implement me")
}

func (n *NativeImpl) Streamer_ProcessActiveItems() {
	C.Streamer_ProcessActiveItems()
}

func (n *NativeImpl) Streamer_ToggleIdleUpdate(playerid int32, toggle bool) {
	toggleVal := 0
	if toggle {
		toggleVal = 1
	}
	C.Streamer_ToggleIdleUpdate(C.int(playerid), C.int(toggleVal))
}

func (n *NativeImpl) Streamer_IsToggleIdleUpdate(playerid int32) bool {
	panic("implement me")
}

func (n *NativeImpl) Streamer_ToggleCameraUpdate(playerid int32, toggle bool) {
	toggleVal := 0
	if toggle {
		toggleVal = 1
	}
	C.Streamer_ToggleCameraUpdate(C.int(playerid), C.int(toggleVal))
}

func (n *NativeImpl) Streamer_IsToggleCameraUpdate(playerid int32) bool {
	return C.Streamer_IsToggleCameraUpdate(C.int(playerid)) != 0
}

func (n *NativeImpl) Streamer_ToggleItemUpdate(playerid int32, streamerType StreamerType, toggle bool) {
	toggleVal := 0
	if toggle {
		toggleVal = 1
	}
	C.Streamer_ToggleItemUpdate(C.int(playerid), C.StreamerType(streamerType), C.int(toggleVal))
}

func (n *NativeImpl) Streamer_IsToggleItemUpdate(playerid int32, streamerType StreamerType) bool {
	return C.Streamer_IsToggleItemUpdate(C.int(playerid), C.StreamerType(streamerType)) != 0
}

func (n *NativeImpl) Streamer_GetLastUpdateTime() float32 {
	return float32(C.Streamer_GetLastUpdateTime())
}

func (n *NativeImpl) Streamer_Update(playerid int32, streamerType StreamerType) {
	C.Streamer_Update(C.int(playerid), C.StreamerType(streamerType))
}

func (n *NativeImpl) Streamer_UpdateEx(playerid int32, x float32, y float32, z float32, worldid int32, interiorid int32, streamerType StreamerType, compensatedtime int32, freezeplayer bool) {
	freezeVal := 0
	if freezeplayer {
		freezeVal = 1
	}
	C.Streamer_UpdateEx(C.int(playerid), C.float(x), C.float(y), C.float(z), C.int(worldid), C.int(interiorid), C.StreamerType(streamerType), C.int(compensatedtime), C.int(freezeVal))
}

func (n *NativeImpl) Streamer_GetFloatData(streamerType StreamerType, id StreamerTag, data int32) float32 {
	return float32(C.Streamer_GetFloatData(C.StreamerType(streamerType), C.StreamerTag(id), C.int(data)))
}

func (n *NativeImpl) Streamer_SetFloatData(streamerType StreamerType, id StreamerTag, data int32, value float32) {
	C.Streamer_SetFloatData(C.StreamerType(streamerType), C.StreamerTag(id), C.int(data), C.float(value))
}

func (n *NativeImpl) Streamer_GetIntData(streamerType StreamerType, id StreamerTag, data int32) int32 {
	return int32(C.Streamer_GetIntData(C.StreamerType(streamerType), C.StreamerTag(id), C.int(data)))
}

func (n *NativeImpl) Streamer_SetIntData(streamerType StreamerType, id StreamerTag, data int32, value int32) {
	C.Streamer_SetIntData(C.StreamerType(streamerType), C.StreamerTag(id), C.int(data), C.int(value))
}

func (n *NativeImpl) Streamer_RemoveIntData(streamerType StreamerType, id StreamerTag, data int32) {
	C.Streamer_RemoveIntData(C.StreamerType(streamerType), C.StreamerTag(id), C.int(data))
}

func (n *NativeImpl) Streamer_HasIntData(streamerType StreamerType, id StreamerTag, data int32) bool {
	return C.Streamer_HasIntData(C.StreamerType(streamerType), C.StreamerTag(id), C.int(data)) != 0
}

func (n *NativeImpl) Streamer_GetArrayData(streamerType StreamerType, id StreamerTag, data int32) []int32 {
	panic("implement me")
}

func (n *NativeImpl) Streamer_SetArrayData(streamerType StreamerType, id StreamerTag, data int32, src []int32) {
	panic("implement me")
}

func (n *NativeImpl) Streamer_IsInArrayData(streamerType StreamerType, id StreamerTag, data int32, value int32) bool {
	return C.Streamer_IsInArrayData(C.StreamerType(streamerType), C.StreamerTag(id), C.int(data), C.int(value)) != 0
}

func (n *NativeImpl) Streamer_AppendArrayData(streamerType StreamerType, id StreamerTag, data int32, value int32) {
	C.Streamer_AppendArrayData(C.StreamerType(streamerType), C.StreamerTag(id), C.int(data), C.int(value))
}

func (n *NativeImpl) Streamer_RemoveArrayData(streamerType StreamerType, id StreamerTag, data int32, value int32) {
	C.Streamer_RemoveArrayData(C.StreamerType(streamerType), C.StreamerTag(id), C.int(data), C.int(value))
}

func (n *NativeImpl) Streamer_HasArrayData(streamerType StreamerType, id StreamerTag, data int32) bool {
	return C.Streamer_HasArrayData(C.StreamerType(streamerType), C.StreamerTag(id), C.int(data)) != 0
}

func (n *NativeImpl) Streamer_GetArrayDataLength(streamerType StreamerType, id StreamerTag, data int32) int32 {
	return int32(C.Streamer_GetArrayDataLength(C.StreamerType(streamerType), C.StreamerTag(id), C.int(data)))
}

func (n *NativeImpl) Streamer_GetUpperBound(streamerType StreamerType) int32 {
	return int32(C.Streamer_GetUpperBound(C.StreamerType(streamerType)))
}

func (n *NativeImpl) Streamer_GetDistanceToItem(pos Vector3, streamerType StreamerType, id StreamerTag, dimensions int32) float32 {
	return float32(C.Streamer_GetDistanceToItem(C.float(pos.X), C.float(pos.Y), C.float(pos.Z), C.StreamerType(streamerType), C.StreamerTag(id), C.int(dimensions)))
}

func (n *NativeImpl) Streamer_ToggleItem(playerid int32, streamerType StreamerType, id StreamerTag, toggle bool) {
	toggleVal := 0
	if toggle {
		toggleVal = 1
	}
	C.Streamer_ToggleItem(C.StreamerType(streamerType), C.StreamerTag(id), C.int(toggleVal))
}

func (n *NativeImpl) Streamer_IsToggleItem(playerid int32, streamerType StreamerType, id StreamerTag) bool {
	return C.Streamer_IsToggleItem(C.StreamerType(streamerType), C.StreamerTag(id)) != 0
}

func (n *NativeImpl) Streamer_ToggleAllItems(playerid int32, streamerType StreamerType, toggle bool, exceptions []int32) {
	panic("implement me")
}

func (n *NativeImpl) Streamer_GetItemInternalID(playerid int32, streamerType StreamerType, streamerid StreamerTag) int32 {
	panic("implement me")
}

func (n *NativeImpl) Streamer_GetItemStreamerID(playerid int32, streamerType StreamerType, internalid int32) StreamerTag {
	panic("implement me")
}

func (n *NativeImpl) Streamer_IsItemVisible(playerid int32, streamerType StreamerType, id StreamerTag) bool {
	panic("implement me")
}

func (n *NativeImpl) Streamer_DestroyAllVisibleItems(playerid int32, streamerType StreamerType, serverwide bool) {
	panic("implement me")
}

func (n *NativeImpl) Streamer_CountVisibleItems(playerid int32, streamerType StreamerType, serverwide bool) int32 {
	panic("implement me")
}

func (n *NativeImpl) Streamer_DestroyAllItems(streamerType StreamerType, serverwide bool) {
	panic("implement me")
}

func (n *NativeImpl) Streamer_CountItems(streamerType StreamerType, serverwide bool) int32 {
	panic("implement me")
}

func (n *NativeImpl) Streamer_GetNearbyItems(pos Vector3, streamerType StreamerType, radius float32, worldid int32) []StreamerTag {
	panic("implement me")
}

func (n *NativeImpl) Streamer_GetAllVisibleItems(playerid int32, streamerType StreamerType) []StreamerTag {
	panic("implement me")
}

func (n *NativeImpl) Streamer_GetItemPos(streamerType StreamerType, id StreamerTag) Vector3 {
	var x, y, z C.float
	C.Streamer_GetItemPos(C.StreamerType(streamerType), C.StreamerTag(id), &x, &y, &z)
	return Vector3{
		X: float32(x),
		Y: float32(y),
		Z: float32(z),
	}
}

func (n *NativeImpl) Streamer_SetItemPos(streamerType StreamerType, id StreamerTag, pos Vector3) {
	C.Streamer_SetItemPos(C.StreamerType(streamerType), C.StreamerTag(id), C.float(pos.X), C.float(pos.Y), C.float(pos.Z))
}

func (n *NativeImpl) Streamer_GetItemOffset(streamerType StreamerType, id StreamerTag) Vector3 {
	var x, y, z C.float
	C.Streamer_GetItemOffset(C.StreamerType(streamerType), C.StreamerTag(id), &x, &y, &z)
	return Vector3{
		X: float32(x),
		Y: float32(y),
		Z: float32(z),
	}
}

func (n *NativeImpl) Streamer_SetItemOffset(streamerType StreamerType, id StreamerTag, offset Vector3) {
	C.Streamer_SetItemOffset(C.StreamerType(streamerType), C.StreamerTag(id), C.float(offset.X), C.float(offset.Y), C.float(offset.Z))
}

func (n *NativeImpl) CreateDynamicObject(modelid int32, pos Vector3, rot Vector3, worldid int32, interiorid int32, playerid int32, streamdistance float32, drawdistance float32, areaid AreaID, priority int32) ObjectID {
	return ObjectID(C.CreateDynamicObject(
		C.int(modelid),
		C.float(pos.X), C.float(pos.Y), C.float(pos.Z),
		C.float(rot.X), C.float(rot.Y), C.float(rot.Z),
		C.int(worldid), C.int(interiorid), C.int(playerid),
		C.float(streamdistance), C.float(drawdistance),
		C.AreaID(areaid), C.int(priority),
	))
}

func (n *NativeImpl) DestroyDynamicObject(objectid ObjectID) {
	C.DestroyDynamicObject(C.ObjectID(objectid))
}

func (n *NativeImpl) IsValidDynamicObject(objectid ObjectID) bool {
	return C.IsValidDynamicObject(C.ObjectID(objectid)) != 0
}

func (n *NativeImpl) GetDynamicObjectPos(objectid ObjectID) Vector3 {
	var x, y, z C.float
	C.GetDynamicObjectPos(C.ObjectID(objectid), &x, &y, &z)
	return Vector3{
		X: float32(x),
		Y: float32(y),
		Z: float32(z),
	}
}

func (n *NativeImpl) SetDynamicObjectPos(objectid ObjectID, pos Vector3) {
	C.SetDynamicObjectPos(C.ObjectID(objectid), C.float(pos.X), C.float(pos.Y), C.float(pos.Z))
}

func (n *NativeImpl) GetDynamicObjectRot(objectid ObjectID) Vector3 {
	var x, y, z C.float
	C.GetDynamicObjectRot(C.ObjectID(objectid), &x, &y, &z)
	return Vector3{
		X: float32(x),
		Y: float32(y),
		Z: float32(z),
	}
}

func (n *NativeImpl) SetDynamicObjectRot(objectid ObjectID, rot Vector3) {
	C.SetDynamicObjectRot(C.ObjectID(objectid), C.float(rot.X), C.float(rot.Y), C.float(rot.Z))
}

func (n *NativeImpl) GetDynamicObjectNoCameraCol(objectid ObjectID) bool {
	return C.GetDynamicObjectNoCameraCol(C.ObjectID(objectid)) != 0
}

func (n *NativeImpl) SetDynamicObjectNoCameraCol(objectid ObjectID) {
	C.SetDynamicObjectNoCameraCol(C.ObjectID(objectid))
}

func (n *NativeImpl) MoveDynamicObject(objectid ObjectID, pos Vector3, speed float32, rot Vector3) {
	C.MoveDynamicObject(
		C.ObjectID(objectid),
		C.float(pos.X), C.float(pos.Y), C.float(pos.Z),
		C.float(speed),
		C.float(rot.X), C.float(rot.Y), C.float(rot.Z),
	)
}

func (n *NativeImpl) StopDynamicObject(objectid ObjectID) {
	C.StopDynamicObject(C.ObjectID(objectid))
}

func (n *NativeImpl) IsDynamicObjectMoving(objectid ObjectID) bool {
	return C.IsDynamicObjectMoving(C.ObjectID(objectid)) != 0
}

func (n *NativeImpl) AttachCameraToDynamicObject(playerid int32, objectid ObjectID) {
	panic("implement me")
}

func (n *NativeImpl) AttachDynamicObjectToObject(objectid ObjectID, attachtoid int32, offset Vector3, rot Vector3, syncrotation bool) {
	panic("implement me")
}

func (n *NativeImpl) AttachDynamicObjectToPlayer(objectid ObjectID, playerid int32, offset Vector3, rot Vector3) {
	panic("implement me")
}

func (n *NativeImpl) AttachDynamicObjectToVehicle(objectid ObjectID, vehicleid int32, offset Vector3, rot Vector3) {
	panic("implement me")
}

func (n *NativeImpl) EditDynamicObject(playerid int32, objectid ObjectID) {
	panic("implement me")
}

func (n *NativeImpl) IsDynamicObjectMaterialUsed(objectid ObjectID, materialindex int32) bool {
	panic("implement me")
}

func (n *NativeImpl) RemoveDynamicObjectMaterial(objectid ObjectID, materialindex int32) {
	panic("implement me")
}

func (n *NativeImpl) GetDynamicObjectMaterial(objectid ObjectID, materialindex int32) (modelid int32, txdname string, texturename string, materialcolor int32) {
	panic("implement me")
}

func (n *NativeImpl) SetDynamicObjectMaterial(objectid ObjectID, materialindex int32, modelid int32, txdname string, texturename string, materialcolor int32) {
	panic("implement me")
}

func (n *NativeImpl) IsDynamicObjectMaterialTextUsed(objectid ObjectID, materialindex int32) bool {
	panic("implement me")
}

func (n *NativeImpl) RemoveDynamicObjectMaterialText(objectid ObjectID, materialindex int32) {
	panic("implement me")
}

func (n *NativeImpl) GetDynamicObjectMaterialText(objectid ObjectID, materialindex int32) (text string, materialsize MaterialSize, fontface string, fontsize int32, bold bool, fontcolor int32, backcolor int32, textalignment int32) {
	panic("implement me")
}

func (n *NativeImpl) SetDynamicObjectMaterialText(objectid ObjectID, materialindex int32, text string, materialsize MaterialSize, fontface string, fontsize int32, bold bool, fontcolor int32, backcolor int32, textalignment int32) {
	panic("implement me")
}

func (n *NativeImpl) GetPlayerCameraTargetDynObject(playerid int32) ObjectID {
	panic("implement me")
}

func (n *NativeImpl) CreateDynamicPickup(modelid int32, pickuptype int32, pos Vector3, worldid int32, interiorid int32, playerid int32, streamdistance float32, areaid AreaID, priority int32) PickupID {
	return PickupID(C.CreateDynamicPickup(
		C.int(modelid), C.int(pickuptype),
		C.float(pos.X), C.float(pos.Y), C.float(pos.Z),
		C.int(worldid), C.int(interiorid), C.int(playerid),
		C.float(streamdistance), C.AreaID(areaid), C.int(priority),
	))
}

func (n *NativeImpl) DestroyDynamicPickup(pickupid PickupID) {
	C.DestroyDynamicPickup(C.PickupID(pickupid))
}

func (n *NativeImpl) IsValidDynamicPickup(pickupid PickupID) bool {
	return C.IsValidDynamicPickup(C.PickupID(pickupid)) != 0
}

func (n *NativeImpl) CreateDynamicCP(pos Vector3, size float32, worldid int32, interiorid int32, playerid int32, streamdistance float32, areaid AreaID, priority int32) CheckpointID {
	panic("implement me")
}

func (n *NativeImpl) DestroyDynamicCP(checkpointid CheckpointID) {
	panic("implement me")
}

func (n *NativeImpl) IsValidDynamicCP(checkpointid CheckpointID) bool {
	panic("implement me")
}

func (n *NativeImpl) IsPlayerInDynamicCP(playerid int32, checkpointid CheckpointID) bool {
	panic("implement me")
}

func (n *NativeImpl) GetPlayerVisibleDynamicCP(playerid int32) CheckpointID {
	panic("implement me")
}

func (n *NativeImpl) CreateDynamicRaceCP(cptype RaceCheckpointType, pos Vector3, nextpos Vector3, size float32, worldid int32, interiorid int32, playerid int32, streamdistance float32, areaid AreaID, priority int32) RaceCheckpointID {
	panic("implement me")
}

func (n *NativeImpl) DestroyDynamicRaceCP(checkpointid RaceCheckpointID) {
	panic("implement me")
}

func (n *NativeImpl) IsValidDynamicRaceCP(checkpointid RaceCheckpointID) bool {
	panic("implement me")
}

func (n *NativeImpl) IsPlayerInDynamicRaceCP(playerid int32, checkpointid RaceCheckpointID) bool {
	panic("implement me")
}

func (n *NativeImpl) GetPlayerVisibleDynamicRaceCP(playerid int32) RaceCheckpointID {
	panic("implement me")
}

func (n *NativeImpl) CreateDynamicMapIcon(pos Vector3, icontype int32, color int32, worldid int32, interiorid int32, playerid int32, streamdistance float32, style int32, areaid AreaID, priority int32) MapIconID {
	panic("implement me")
}

func (n *NativeImpl) DestroyDynamicMapIcon(iconid MapIconID) {
	panic("implement me")
}

func (n *NativeImpl) IsValidDynamicMapIcon(iconid MapIconID) bool {
	panic("implement me")
}

func (n *NativeImpl) CreateDynamic3DTextLabel(text string, color int32, pos Vector3, drawdistance float32, attachedplayer int32, attachedvehicle int32, testlos bool, worldid int32, interiorid int32, playerid int32, streamdistance float32, areaid AreaID, priority int32) TextLabelID {
	panic("implement me")
}

func (n *NativeImpl) DestroyDynamic3DTextLabel(id TextLabelID) {
	panic("implement me")
}

func (n *NativeImpl) IsValidDynamic3DTextLabel(id TextLabelID) bool {
	panic("implement me")
}

func (n *NativeImpl) GetDynamic3DTextLabelText(id TextLabelID) string {
	panic("implement me")
}

func (n *NativeImpl) UpdateDynamic3DTextLabelText(id TextLabelID, color int32, text string) {
	panic("implement me")
}

func (n *NativeImpl) CreateDynamicCircle(center Vector2, size float32, worldid int32, interiorid int32, playerid int32, priority int32) AreaID {
	return AreaID(C.CreateDynamicCircle(
		C.float(center.X), C.float(center.Y),
		C.float(size),
		C.int(worldid), C.int(interiorid), C.int(playerid),
		C.int(priority),
	))
}

func (n *NativeImpl) CreateDynamicCylinder(pos Vector2, minz float32, maxz float32, size float32, worldid int32, interiorid int32, playerid int32, priority int32) AreaID {
	panic("implement me")
}

func (n *NativeImpl) CreateDynamicSphere(pos Vector3, size float32, worldid int32, interiorid int32, playerid int32, priority int32) AreaID {
	panic("implement me")
}

func (n *NativeImpl) CreateDynamicRectangle(min Vector2, max Vector2, worldid int32, interiorid int32, playerid int32, priority int32) AreaID {
	panic("implement me")
}

func (n *NativeImpl) CreateDynamicCuboid(min Vector3, max Vector3, worldid int32, interiorid int32, playerid int32, priority int32) AreaID {
	panic("implement me")
}

func (n *NativeImpl) CreateDynamicCube(min Vector3, max Vector3, worldid int32, interiorid int32, playerid int32, priority int32) AreaID {
	panic("implement me")
}

func (n *NativeImpl) CreateDynamicPolygon(points []Vector2, minz float32, maxz float32, worldid int32, interiorid int32, playerid int32, priority int32) AreaID {
	if len(points) == 0 {
		return 0
	}

	pointsArray := C.malloc(C.size_t(len(points) * 8)) // 8 bytes per Vector2 (2 floats)
	defer C.free(pointsArray)

	for i, point := range points {
		*(*C.float)(unsafe.Pointer(uintptr(pointsArray) + uintptr(i*8))) = C.float(point.X)
		*(*C.float)(unsafe.Pointer(uintptr(pointsArray) + uintptr(i*8+4))) = C.float(point.Y)
	}

	return AreaID(C.CreateDynamicPolygon(
		(*C.float)(pointsArray), C.int(len(points)),
		C.float(minz), C.float(maxz),
		C.int(worldid), C.int(interiorid), C.int(playerid),
		C.int(priority),
	))
}

func (n *NativeImpl) DestroyDynamicArea(areaid AreaID) {
	panic("implement me")
}

func (n *NativeImpl) IsValidDynamicArea(areaid AreaID) bool {
	panic("implement me")
}

func (n *NativeImpl) GetDynamicAreaType(areaid AreaID) int32 {
	panic("implement me")
}

func (n *NativeImpl) GetDynamicPolygonPoints(areaid AreaID) []Vector2 {
	panic("implement me")
}

func (n *NativeImpl) GetDynamicPolygonNumberPoints(areaid AreaID) int32 {
	panic("implement me")
}

func (n *NativeImpl) IsPlayerInDynamicArea(playerid int32, areaid AreaID, recheck bool) bool {
	panic("implement me")
}

func (n *NativeImpl) IsPlayerInAnyDynamicArea(playerid int32, recheck bool) bool {
	panic("implement me")
}

func (n *NativeImpl) IsAnyPlayerInDynamicArea(areaid AreaID, recheck bool) bool {
	panic("implement me")
}

func (n *NativeImpl) IsAnyPlayerInAnyDynamicArea(recheck bool) bool {
	panic("implement me")
}

func (n *NativeImpl) GetPlayerDynamicAreas(playerid int32) []AreaID {
	panic("implement me")
}

func (n *NativeImpl) GetPlayerNumberDynamicAreas(playerid int32) int32 {
	panic("implement me")
}

func (n *NativeImpl) IsPointInDynamicArea(areaid AreaID, point Vector3) bool {
	panic("implement me")
}

func (n *NativeImpl) IsPointInAnyDynamicArea(point Vector3) bool {
	panic("implement me")
}

func (n *NativeImpl) IsLineInDynamicArea(areaid AreaID, start Vector3, end Vector3) bool {
	panic("implement me")
}

func (n *NativeImpl) IsLineInAnyDynamicArea(start Vector3, end Vector3) bool {
	panic("implement me")
}

func (n *NativeImpl) GetDynamicAreasForPoint(point Vector3) []AreaID {
	panic("implement me")
}

func (n *NativeImpl) GetNumberDynamicAreasForPoint(point Vector3) int32 {
	panic("implement me")
}

func (n *NativeImpl) GetDynamicAreasForLine(start Vector3, end Vector3) []AreaID {
	panic("implement me")
}

func (n *NativeImpl) GetNumberDynamicAreasForLine(start Vector3, end Vector3) int32 {
	panic("implement me")
}

func (n *NativeImpl) AttachDynamicAreaToObject(areaid AreaID, objectid ObjectID, objecttype int32, playerid int32, offset Vector3) {
	panic("implement me")
}

func (n *NativeImpl) AttachDynamicAreaToPlayer(areaid AreaID, playerid int32, offset Vector3) {
	panic("implement me")
}

func (n *NativeImpl) AttachDynamicAreaToVehicle(areaid AreaID, vehicleid int32, offset Vector3) {
	panic("implement me")
}

func (n *NativeImpl) ToggleDynAreaSpectateMode(areaid AreaID, toggle bool) {
	panic("implement me")
}

func (n *NativeImpl) IsToggleDynAreaSpectateMode(areaid AreaID) bool {
	panic("implement me")
}

func (n *NativeImpl) CreateDynamicActor(modelid int32, pos Vector3, rotation float32, invulnerable bool, health float32, worldid int32, interiorid int32, playerid int32, streamdistance float32, areaid AreaID, priority int32) ActorID {
	invulnerableVal := 0
	if invulnerable {
		invulnerableVal = 1
	}

	return ActorID(C.CreateDynamicActor(
		C.int(modelid),
		C.float(pos.X), C.float(pos.Y), C.float(pos.Z),
		C.float(rotation),
		C.int(invulnerableVal), C.float(health),
		C.int(worldid), C.int(interiorid), C.int(playerid),
		C.float(streamdistance), C.AreaID(areaid), C.int(priority),
	))
}

func (n *NativeImpl) DestroyDynamicActor(actorid ActorID) {
	C.DestroyDynamicActor(C.ActorID(actorid))
}

func (n *NativeImpl) IsValidDynamicActor(actorid ActorID) bool {
	return C.IsValidDynamicActor(C.ActorID(actorid)) != 0
}

func (n *NativeImpl) IsDynamicActorStreamedIn(actorid ActorID, forplayerid int32) bool {
	panic("implement me")
}

func (n *NativeImpl) GetDynamicActorVirtualWorld(actorid ActorID) int32 {
	panic("implement me")
}

func (n *NativeImpl) SetDynamicActorVirtualWorld(actorid ActorID, vworld int32) {
	panic("implement me")
}

func (n *NativeImpl) GetDynamicActorAnimation(actorid ActorID) (animlib string, animname string, fdelta float32, loop bool, lockx bool, locky bool, freeze bool, time int32) {
	panic("implement me")
}

func (n *NativeImpl) ApplyDynamicActorAnimation(actorid ActorID, animlib string, animname string, fdelta float32, loop bool, lockx bool, locky bool, freeze bool, time int32) {
	loopVal, lockxVal, lockyVal, freezeVal := 0, 0, 0, 0

	if loop {
		loopVal = 1
	}
	if lockx {
		lockxVal = 1
	}
	if locky {
		lockyVal = 1
	}
	if freeze {
		freezeVal = 1
	}

	canimlib := C.CString(animlib)
	defer C.free(unsafe.Pointer(canimlib))

	canimname := C.CString(animname)
	defer C.free(unsafe.Pointer(canimname))

	C.ApplyDynamicActorAnimation(
		C.ActorID(actorid),
		canimlib, canimname,
		C.float(fdelta),
		C.int(loopVal), C.int(lockxVal), C.int(lockyVal), C.int(freezeVal),
		C.int(time),
	)
}

func (n *NativeImpl) ClearDynamicActorAnimations(actorid ActorID) {
	panic("implement me")
}

func (n *NativeImpl) GetDynamicActorFacingAngle(actorid ActorID) float32 {
	panic("implement me")
}

func (n *NativeImpl) SetDynamicActorFacingAngle(actorid ActorID, angle float32) {
	panic("implement me")
}

func (n *NativeImpl) GetDynamicActorPos(actorid ActorID) Vector3 {
	panic("implement me")
}

func (n *NativeImpl) SetDynamicActorPos(actorid ActorID, pos Vector3) {
	panic("implement me")
}

func (n *NativeImpl) GetDynamicActorHealth(actorid ActorID) float32 {
	panic("implement me")
}

func (n *NativeImpl) SetDynamicActorHealth(actorid ActorID, health float32) {
	panic("implement me")
}

func (n *NativeImpl) SetDynamicActorInvulnerable(actorid ActorID, invulnerable bool) {
	panic("implement me")
}

func (n *NativeImpl) IsDynamicActorInvulnerable(actorid ActorID) bool {
	panic("implement me")
}

func (n *NativeImpl) GetPlayerTargetDynamicActor(playerid int32) ActorID {
	return ActorID(C.GetPlayerTargetDynamicActor(C.int(playerid)))
}

func (n *NativeImpl) GetPlayerCameraTargetDynActor(playerid int32) ActorID {
	return ActorID(C.GetPlayerCameraTargetDynActor(C.int(playerid)))
}

func (n *NativeImpl) CreateDynamicObjectEx(modelid int32, pos Vector3, rot Vector3, streamdistance float32, drawdistance float32, worlds []int32, interiors []int32, players []int32, areas []AreaID, priority int32) ObjectID {
	worldsPtr, worldsLen := convertSliceToIntArray(worlds)
	if worldsPtr != nil {
		defer C.free(unsafe.Pointer(worldsPtr))
	}

	interiorsPtr, interiorsLen := convertSliceToIntArray(interiors)
	if interiorsPtr != nil {
		defer C.free(unsafe.Pointer(interiorsPtr))
	}

	playersPtr, playersLen := convertSliceToIntArray(players)
	if playersPtr != nil {
		defer C.free(unsafe.Pointer(playersPtr))
	}

	// Convert areas to int32 slice
	areasInt := make([]int32, len(areas))
	for i, a := range areas {
		areasInt[i] = int32(a)
	}

	areasPtr, areasLen := convertSliceToIntArray(areasInt)
	if areasPtr != nil {
		defer C.free(unsafe.Pointer(areasPtr))
	}

	return ObjectID(C.CreateDynamicObjectEx(
		C.int(modelid),
		C.float(pos.X), C.float(pos.Y), C.float(pos.Z),
		C.float(rot.X), C.float(rot.Y), C.float(rot.Z),
		C.float(streamdistance), C.float(drawdistance),
		worldsPtr, worldsLen,
		interiorsPtr, interiorsLen,
		playersPtr, playersLen,
		(*C.int)(unsafe.Pointer(areasPtr)), areasLen,
		C.int(priority),
	))
}

func (n *NativeImpl) CreateDynamicPickupEx(modelid int32, pickuptype int32, pos Vector3, streamdistance float32, worlds []int32, interiors []int32, players []int32, areas []AreaID, priority int32) PickupID {
	worldsPtr, worldsLen := convertSliceToIntArray(worlds)
	if worldsPtr != nil {
		defer C.free(unsafe.Pointer(worldsPtr))
	}

	interiorsPtr, interiorsLen := convertSliceToIntArray(interiors)
	if interiorsPtr != nil {
		defer C.free(unsafe.Pointer(interiorsPtr))
	}

	playersPtr, playersLen := convertSliceToIntArray(players)
	if playersPtr != nil {
		defer C.free(unsafe.Pointer(playersPtr))
	}

	// Convert areas to int32 slice
	areasInt := make([]int32, len(areas))
	for i, a := range areas {
		areasInt[i] = int32(a)
	}

	areasPtr, areasLen := convertSliceToIntArray(areasInt)
	if areasPtr != nil {
		defer C.free(unsafe.Pointer(areasPtr))
	}

	return PickupID(C.CreateDynamicPickupEx(
		C.int(modelid), C.int(pickuptype),
		C.float(pos.X), C.float(pos.Y), C.float(pos.Z),
		C.float(streamdistance),
		worldsPtr, worldsLen,
		interiorsPtr, interiorsLen,
		playersPtr, playersLen,
		(*C.int)(unsafe.Pointer(areasPtr)), areasLen,
		C.int(priority),
	))
}

func (n *NativeImpl) CreateDynamicCPEx(pos Vector3, size float32, streamdistance float32, worlds []int32, interiors []int32, players []int32, areas []AreaID, priority int32) CheckpointID {
	panic("implement me")
}

func (n *NativeImpl) CreateDynamicRaceCPEx(cptype int32, pos Vector3, nextpos Vector3, size float32, streamdistance float32, worlds []int32, interiors []int32, players []int32, areas []AreaID, priority int32) RaceCheckpointID {
	panic("implement me")
}

func (n *NativeImpl) CreateDynamicMapIconEx(pos Vector3, icontype int32, color int32, style int32, streamdistance float32, worlds []int32, interiors []int32, players []int32, areas []AreaID, priority int32) MapIconID {
	panic("implement me")
}

func (n *NativeImpl) CreateDynamic3DTextLabelEx(text string, color int32, pos Vector3, drawdistance float32, attachedplayer int32, attachedvehicle int32, testlos bool, streamdistance float32, worlds []int32, interiors []int32, players []int32, areas []AreaID, priority int32) TextLabelID {
	panic("implement me")
}

func (n *NativeImpl) CreateDynamicCircleEx(center Vector2, size float32, worlds []int32, interiors []int32, players []int32, priority int32) AreaID {
	worldsPtr, worldsLen := convertSliceToIntArray(worlds)
	if worldsPtr != nil {
		defer C.free(unsafe.Pointer(worldsPtr))
	}

	interiorsPtr, interiorsLen := convertSliceToIntArray(interiors)
	if interiorsPtr != nil {
		defer C.free(unsafe.Pointer(interiorsPtr))
	}

	playersPtr, playersLen := convertSliceToIntArray(players)
	if playersPtr != nil {
		defer C.free(unsafe.Pointer(playersPtr))
	}

	return AreaID(C.CreateDynamicCircleEx(
		C.float(center.X), C.float(center.Y),
		C.float(size),
		worldsPtr, worldsLen,
		interiorsPtr, interiorsLen,
		playersPtr, playersLen,
		C.int(priority),
	))
}

func (n *NativeImpl) CreateDynamicCylinderEx(pos Vector2, minz float32, maxz float32, size float32, worlds []int32, interiors []int32, players []int32, priority int32) AreaID {
	panic("implement me")
}

func (n *NativeImpl) CreateDynamicSphereEx(pos Vector3, size float32, worlds []int32, interiors []int32, players []int32, priority int32) AreaID {
	panic("implement me")
}

func (n *NativeImpl) CreateDynamicRectangleEx(min Vector2, max Vector2, worlds []int32, interiors []int32, players []int32, priority int32) AreaID {
	panic("implement me")
}

func (n *NativeImpl) CreateDynamicCuboidEx(min Vector3, max Vector3, worlds []int32, interiors []int32, players []int32, priority int32) AreaID {
	panic("implement me")
}

func (n *NativeImpl) CreateDynamicCubeEx(min Vector3, max Vector3, worlds []int32, interiors []int32, players []int32, priority int32) AreaID {
	panic("implement me")
}

func (n *NativeImpl) CreateDynamicPolygonEx(points []Vector2, minz float32, maxz float32, worlds []int32, interiors []int32, players []int32, priority int32) AreaID {
	panic("implement me")
}

func (n *NativeImpl) CreateDynamicActorEx(modelid int32, pos Vector3, rotation float32, invulnerable bool, health float32, streamdistance float32, worlds []int32, interiors []int32, players []int32, areas []AreaID, priority int32) ActorID {
	invulnerableVal := 0
	if invulnerable {
		invulnerableVal = 1
	}

	worldsPtr, worldsLen := convertSliceToIntArray(worlds)
	if worldsPtr != nil {
		defer C.free(unsafe.Pointer(worldsPtr))
	}

	interiorsPtr, interiorsLen := convertSliceToIntArray(interiors)
	if interiorsPtr != nil {
		defer C.free(unsafe.Pointer(interiorsPtr))
	}

	playersPtr, playersLen := convertSliceToIntArray(players)
	if playersPtr != nil {
		defer C.free(unsafe.Pointer(playersPtr))
	}

	// Convert areas to int32 slice
	areasInt := make([]int32, len(areas))
	for i, a := range areas {
		areasInt[i] = int32(a)
	}

	areasPtr, areasLen := convertSliceToIntArray(areasInt)
	if areasPtr != nil {
		defer C.free(unsafe.Pointer(areasPtr))
	}

	return ActorID(C.CreateDynamicActorEx(
		C.int(modelid),
		C.float(pos.X), C.float(pos.Y), C.float(pos.Z),
		C.float(rotation),
		C.int(invulnerableVal), C.float(health),
		C.float(streamdistance),
		worldsPtr, worldsLen,
		interiorsPtr, interiorsLen,
		playersPtr, playersLen,
		(*C.int)(unsafe.Pointer(areasPtr)), areasLen,
		C.int(priority),
	))
}

// NewNative crea una nueva instancia de la implementación nativa
func NewNative() Native {
	return &NativeImpl{}
}

// Streamer_GetTickRate obtiene la tasa de ticks general actual
func (n *NativeImpl) Streamer_GetTickRate() int32 {
	return int32(C.Streamer_GetTickRate())
}

// Streamer_SetTickRate establece la tasa de ticks general
func (n *NativeImpl) Streamer_SetTickRate(tickrate int32) {
	C.Streamer_SetTickRate(C.int(tickrate))
}

// Streamer_GetPlayerTickRate obtiene la tasa de ticks específica del jugador
func (n *NativeImpl) Streamer_GetPlayerTickRate(playerid int32) int32 {
	return int32(C.Streamer_GetPlayerTickRate(C.int(playerid)))
}

// Streamer_SetPlayerTickRate establece la tasa de ticks específica del jugador
func (n *NativeImpl) Streamer_SetPlayerTickRate(playerid int32, tickrate int32) {
	C.Streamer_SetPlayerTickRate(C.int(playerid), C.int(tickrate))
}

// Streamer_ToggleChunkStreaming activa/desactiva el streaming por trozos
func (n *NativeImpl) Streamer_ToggleChunkStreaming(toggle bool) {
	var toggleVal C.int
	if toggle {
		toggleVal = 1
	}
	C.Streamer_ToggleChunkStream(toggleVal)
}

// Streamer_IsToggleChunkStreaming comprueba si el streaming por trozos está activado
func (n *NativeImpl) Streamer_IsToggleChunkStreaming() bool {
	return C.Streamer_IsToggleChunkStream() != 0
}

func convertSliceToIntArray(slice []int32) (*C.int, C.int) {
	if len(slice) == 0 {
		return nil, 0
	}

	cArray := C.malloc(C.size_t(len(slice)) * C.size_t(unsafe.Sizeof(C.int(0))))
	sliceHeader := (*[1 << 30]C.int)(cArray)

	for i, v := range slice {
		sliceHeader[i] = C.int(v)
	}

	return (*C.int)(cArray), C.int(len(slice))
}
