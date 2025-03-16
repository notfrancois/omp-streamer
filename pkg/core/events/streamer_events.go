package events

import (
	"reflect"
	"sync"
)

type EventType string

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

type EventManager struct {
	handlers map[EventType][]EventHandler
	mutex    sync.RWMutex
}

type EventHandler func(data interface{}) bool

// NewEventManager creates a new EventManager instance
func NewEventManager() *EventManager {
	return &EventManager{
		handlers: make(map[EventType][]EventHandler),
	}
}

// RegisterHandler registers a handler for an event type
func (em *EventManager) RegisterHandler(event EventType, handler EventHandler) {
	em.mutex.Lock()
	defer em.mutex.Unlock()
	em.handlers[event] = append(em.handlers[event], handler)
}

// TriggerEvent triggers an event
func (em *EventManager) TriggerEvent(event EventType, data interface{}) bool {
	em.mutex.RLock()
	handlers := em.handlers[event]
	em.mutex.RUnlock()

	for _, handler := range handlers {
		if !handler(data) {
			return false
		}
	}
	return true
}

// Unsubscribe unsubscribes from an event type
func (em *EventManager) Unsubscribe(eventType EventType, callback EventHandler) {
	em.mutex.Lock()
	defer em.mutex.Unlock()

	callbacks := em.handlers[eventType]
	for i, cb := range callbacks {
		if reflect.ValueOf(cb).Pointer() == reflect.ValueOf(callback).Pointer() {
			em.handlers[eventType] = append(callbacks[:i], callbacks[i+1:]...)
			break
		}
	}
}
