package constant

const (
	EventObjectMoved        EventType = "OnDynamicObjectMoved"
	EventPlayerEditObject   EventType = "OnPlayerEditDynamicObject"
	EventPlayerSelectObject EventType = "OnPlayerSelectDynamicObject"
	EventPlayerShootObject  EventType = "OnPlayerShootDynamicObject"
	EventPickupPickup       EventType = "OnPlayerPickUpDynamicPickup"
	EventEnterCheckpoint    EventType = "OnPlayerEnterDynamicCP"
	EventLeaveCheckpoint    EventType = "OnPlayerLeaveDynamicCP"
	EventEnterArea          EventType = "OnPlayerEnterDynamicArea"
	EventLeaveArea          EventType = "OnPlayerLeaveDynamicArea"
	EventActorStreamIn      EventType = "OnDynamicActorStreamIn"
	EventActorStreamOut     EventType = "OnDynamicActorStreamOut"
	EventItemStreamIn       EventType = "Streamer_OnItemStreamIn"
	EventItemStreamOut      EventType = "Streamer_OnItemStreamOut"
	EventPluginError        EventType = "Streamer_OnPluginError"
)
