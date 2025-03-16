package streamer

import (
	"time"
)

// StartAutomaticUpdate inicia el proceso de actualización automática
func (s *Streamer) StartAutomaticUpdate() {
	// Verificar si hay interfaces registradas
	if len(s.core.Data.Interfaces) == 0 {
		return
	}

	currentTime := time.Now()
	
	// Si hay jugadores, procesar sus actualizaciones
	if len(s.core.Data.Players) > 0 {
		updatedActiveItems := false
		
		// Iterar por todos los jugadores
		for playerID, player := range s.core.Data.Players {
			// Si el streaming por chunks está habilitado y el jugador está procesando chunks
			if s.core.ChunkStreamer.chunkStreamingEnabled && !player.ProcessingChunks.IsEmpty() {
				s.core.ChunkStreamer.PerformPlayerChunkUpdate(player, true)
			} else {
				// Incrementar el contador de ticks y verificar si es tiempo de actualizar
				player.TickCount++
				if player.TickCount >= player.TickRate {
					// Si aún no hemos actualizado los items activos, hacerlo ahora
					if !updatedActiveItems {
						s.ProcessActiveItems()
						updatedActiveItems = true
					}
					
					// Realizar la actualización del jugador
					if !player.DelayedUpdate {
						s.PerformPlayerUpdate(player, true)
					} else {
						s.StartManualUpdate(player, player.DelayedUpdateType)
					}
					
					// Reiniciar el contador de ticks
					player.TickCount = 0
					
					// Actualizar el jugador en el mapa ya que modificamos sus valores
					s.core.Data.Players[playerID] = player
				}
			}
		}
	} else {
		// Si no hay jugadores, solo actualizar los items activos
		s.ProcessActiveItems()
	}
	
	// Actualizar el tiempo de la última actualización
	s.lastUpdateTime = float32(time.Since(s.lastUpdateTimePoint).Seconds())
	s.lastUpdateTimePoint = currentTime
}

// PerformPlayerUpdate realiza una actualización para un jugador específico
func (s *Streamer) PerformPlayerUpdate(player *Player, automatic bool) {
	// Obtener la posición actual del jugador
	position := player.Position
	delta := Vector3f{}
	
	// Si es una actualización automática, calcular delta para detectar movimiento
	if automatic {
		// Solo procesar si el jugador está conectado
		if !s.core.SAMP.IsPlayerConnected(player.PlayerId) {
			return
		}
		
		// Verificar si el jugador debe actualizarse estando inactivo
		if !player.UpdateWhenIdle && s.IsPlayerIdle(player)) {
			return
		}
		
		// Actualizar interior y mundo virtual del jugador
		player.InteriorId = s.core.SAMP.GetPlayerInterior(player.PlayerId)
		player.WorldId = s.core.SAMP.GetPlayerVirtualWorld(player.PlayerId)
		
		// Obtener posición actual
		newPosition := s.core.SAMP.GetPlayerPos(player.PlayerId)
		
		// Calcular el delta (movimiento)
		delta = Vector3f{
			X: newPosition.X - player.Position.X,
			Y: newPosition.Y - player.Position.Y,
			Z: newPosition.Z - player.Position.Z,
		}
		
		// Actualizar la posición del jugador
		player.Position = newPosition
	}
	
	// Si el jugador no está actualizando su cámara o si es una actualización manual
	if !player.UpdateCameraPosition || !automatic {
		// Obtener las celdas cercanas al jugador
		cells := s.core.Grid.GetCellsInRange(player.Position.X, player.Position.Y, 
			s.core.Data.CellDistance, player.WorldId, player.InteriorId)
		
		// Procesar cada tipo de elemento si está habilitado para el jugador
		for streamType := 0; streamType < StreamerMaxTypes; streamType++ {
			if player.EnabledItems[streamType] {
				switch streamType {
				case StreamerTypeObject:
					if len(s.core.Data.Objects) > 0 {
						s.DiscoverObjects(player, cells)
					}
				case StreamerTypePickup:
					if len(s.core.Data.Pickups) > 0 {
						s.ProcessPickups(player, cells)
					}
				case StreamerTypeCP:
					if len(s.core.Data.Checkpoints) > 0 {
						s.ProcessCheckpoints(player, cells)
					}
				case StreamerTypeRaceCP:
					if len(s.core.Data.RaceCheckpoints) > 0 {
						s.ProcessRaceCheckpoints(player, cells)
					}
				case StreamerTypeMapIcon:
					if len(s.core.Data.MapIcons) > 0 {
						s.DiscoverMapIcons(player, cells)
					}
				case StreamerType3DTextLabel:
					if len(s.core.Data.TextLabels) > 0 {
						s.DiscoverTextLabels(player, cells)
					}
				case StreamerTypeArea:
					if len(s.core.Data.Areas) > 0 {
						// Si hay delta (movimiento), usar posición anterior para el procesamiento
						if !delta.IsZero() {
							player.Position = position
						}
						
						s.ProcessAreas(player, cells)
						
						// Restaurar posición si hubo movimiento
						if !delta.IsZero() {
							player.Position = player.Position.Add(delta)
						}
					}
				}
			}
		}
		
		// Restaurar posición original si hubo movimiento
		if !delta.IsZero() {
			player.Position = position
		}
	}
}

// ExecuteCallbacks ejecuta las callbacks de streaming pendientes
func (s *Streamer) ExecuteCallbacks() {
	// Procesar callbacks de salida de áreas
	if len(s.areaLeaveCallbacks) > 0 {
		// Copiar y limpiar la lista de callbacks para iterar de forma segura
		callbacks := s.areaLeaveCallbacks
		s.areaLeaveCallbacks = nil
		
		// Procesar en orden inverso (como en el código original)
		for i := len(callbacks) - 1; i >= 0; i-- {
			callback := callbacks[i]
			
			// Verificar si el área existe
			area, exists := s.core.Data.Areas[callback.ItemID]
			if exists {
				// Llamar a la callback en todas las interfaces
				for _, iface := range s.core.Data.Interfaces {
					// En Go no tenemos AMX, así que usamos el sistema de callbacks propio
					if s.core.Callbacks != nil {
						s.core.Callbacks.OnPlayerLeaveDynamicArea(int32(callback.PlayerID), types.AreaID(callback.ItemID))
					}
				}
			}
		}
	}
	
	// Procesar callbacks de entrada a áreas
	if len(s.areaEnterCallbacks) > 0 {
		callbacks := s.areaEnterCallbacks
		s.areaEnterCallbacks = nil
		
		for i := len(callbacks) - 1; i >= 0; i-- {
			callback := callbacks[i]
			
			area, exists := s.core.Data.Areas[callback.ItemID]
			if exists {
				for _, iface := range s.core.Data.Interfaces {
					if s.core.Callbacks != nil {
						s.core.Callbacks.OnPlayerEnterDynamicArea(int32(callback.PlayerID), types.AreaID(callback.ItemID))
					}
				}
			}
		}
	}
	
	// Procesar callbacks de StreamIn
	if len(s.streamInCallbacks) > 0 {
		callbacks := s.streamInCallbacks
		s.streamInCallbacks = nil
		
		for _, callback := range callbacks {
			if s.core.Callbacks != nil {
				s.core.Callbacks.Streamer_OnItemStreamIn(types.StreamerType(callback.ItemType), 
					types.StreamerTag(callback.ItemID), int32(callback.PlayerID))
			}
		}
	}
	
	// Procesar callbacks de StreamOut
	if len(s.streamOutCallbacks) > 0 {
		callbacks := s.streamOutCallbacks
		s.streamOutCallbacks = nil
		
		for _, callback := range callbacks {
			if s.core.Callbacks != nil {
				s.core.Callbacks.Streamer_OnItemStreamOut(types.StreamerType(callback.ItemType), 
					types.StreamerTag(callback.ItemID), int32(callback.PlayerID))
			}
		}
	}
}

// IsPlayerIdle verifica si un jugador está inactivo
func (s *Streamer) IsPlayerIdle(player *Player) bool {
	// Obtener posición actual del jugador
	currentPos := s.core.SAMP.GetPlayerPos(player.PlayerId)
	
	// Si la posición no cambió y no actualizó su cámara, está inactivo
	if player.LastUpdatePos.X == currentPos.X && 
	   player.LastUpdatePos.Y == currentPos.Y && 
	   player.LastUpdatePos.Z == currentPos.Z {
		return true
	}
	
	// Actualizar última posición conocida
	player.LastUpdatePos = currentPos
	return false
}

// StartManualUpdate inicia una actualización manual para un jugador
func (s *Streamer) StartManualUpdate(player *Player, streamType int) {
	// Guardar estado actual de los items habilitados
	enabledItems := player.EnabledItems
	
	// Verificar si hay actualización retrasada
	if player.DelayedUpdate {
		if player.DelayedUpdateTime.Before(time.Now()) {
			// Si es necesario, descongelar al jugador
			if player.DelayedUpdateFreeze {
				s.core.SAMP.TogglePlayerControllable(player.PlayerId, true)
			}
			player.DelayedUpdate = false
		}
	}
	
	// Si se especificó un tipo específico
	if streamType >= 0 && streamType < StreamerMaxTypes {
		// Si el streaming por chunks está habilitado, limpiar datos de descubrimiento
		if s.core.ChunkStreamer.chunkStreamingEnabled {
			switch streamType {
			case StreamerTypeObject:
				player.DiscoveredObjects.Clear()
				player.ExistingObjects.Clear()
				player.ProcessingChunks.Reset(StreamerTypeObject)
			case StreamerTypeMapIcon:
				player.DiscoveredMapIcons.Clear()
				player.ExistingMapIcons.Clear()
				player.ProcessingChunks.Reset(StreamerTypeMapIcon)
			case StreamerType3DTextLabel:
				player.DiscoveredTextLabels.Clear()
				player.ExistingTextLabels.Clear()
				player.ProcessingChunks.Reset(StreamerType3DTextLabel)
			}
		}
		
		// Configurar solo el tipo específico para actualizar
		player.EnabledItems = make([]bool, StreamerMaxTypes)
		player.EnabledItems[streamType] = true
	} else if s.core.ChunkStreamer.chunkStreamingEnabled {
		// Limpiar todo si no se especificó un tipo
		player.DiscoveredMapIcons.Clear()
		player.DiscoveredObjects.Clear()
		player.DiscoveredTextLabels.Clear()
		player.ExistingMapIcons.Clear()
		player.ExistingObjects.Clear()
		player.ExistingTextLabels.Clear()
		player.ProcessingChunks.Reset()
	}
	
	// Procesar items activos y realizar actualización
	s.ProcessActiveItems()
	s.PerformPlayerUpdate(player, false)
	
	// Si el streaming por chunks está habilitado, realizar actualización de chunks
	if s.core.ChunkStreamer.chunkStreamingEnabled {
		s.core.ChunkStreamer.PerformPlayerChunkUpdate(player, false)
	}
	
	// Restaurar items habilitados originales
	player.EnabledItems = enabledItems
}

// ProcessActiveItems procesa todos los elementos activos (objetos en movimiento, áreas adjuntas, etc.)
func (s *Streamer) ProcessActiveItems() {
	// Procesar objetos en movimiento
	if len(s.movingObjects) > 0 {
		currentTime := time.Now()
		
		for objectID, obj := range s.movingObjects {
			elapsed := currentTime.Sub(obj.MoveStartTime).Seconds()
			totalTime := float64(obj.MoveDuration) / 1000.0 // Convertir de ms a segundos
			
			if elapsed >= totalTime {
				// El objeto ha terminado de moverse
				obj.Position = obj.MoveTargetPos
				obj.Moving = false
				
				delete(s.movingObjects, objectID)
				s.core.Data.Objects[objectID] = obj
				
				// Actualizar para los jugadores
				for _, player := range s.core.Data.Players {
					if internalID, found := player.InternalObjects[objectID]; found {
						s.core.SAMP.StopPlayerObject(player.PlayerId, internalID)
						s.core.SAMP.SetPlayerObjectPos(player.PlayerId, internalID, 
							obj.Position.X, obj.Position.Y, obj.Position.Z)
					}
				}
				
				// Lanzar callback de objeto movido
				if s.core.Callbacks != nil {
					s.core.Callbacks.OnDynamicObjectMoved(types.ObjectID(objectID))
				}
			}
		}
	}
	
	// Procesar áreas adjuntas
	if len(s.attachedAreas) > 0 {
		for areaID, area := range s.attachedAreas {
			if area.AttachedObject != nil {
				// Actualizar posición basada en el objeto
				obj, exists := s.core.Data.Objects[area.AttachedObject.Object]
				if exists {
					newPos := obj.Position.Add(area.AttachedObject.Position)
					area.Position = newPos
					s.core.Data.Areas[areaID] = area
				}
			} else if area.AttachedPlayer > -1 {
				// Actualizar posición basada en el jugador
				playerPos := s.core.SAMP.GetPlayerPos(area.AttachedPlayer)
				newPos := playerPos.Add(area.AttachedObject.Position)
				area.Position = newPos
				s.core.Data.Areas[areaID] = area
			} else if area.AttachedVehicle > -1 {
				// Actualizar posición basada en el vehículo
				vehiclePos := s.core.SAMP.GetVehiclePos(area.AttachedVehicle)
				newPos := vehiclePos.Add(area.AttachedObject.Position)
				area.Position = newPos
				s.core.Data.Areas[areaID] = area
			}
		}
	}
	
	// Procesar objetos adjuntos
	// (Implementación similar para objetos adjuntos a jugadores, vehículos u otros objetos)
	if len(s.attachedObjects) > 0 {
		for objectID, obj := range s.attachedObjects {
			// Actualizar posición basada en el objeto
			objPos := s.core.Data.Objects[objectID].Position
			obj.Position = objPos
			s.core.Data.Objects[objectID] = obj
		}
	}

	// Procesar etiquetas de texto adjuntas
	// (Implementación similar para texto adjunto a jugadores, vehículos u objetos)
	if len(s.attachedTextLabels) > 0 {
		for textLabelID, textLabel := range s.attachedTextLabels {
			// Actualizar posición basada en el objeto
			objPos := s.core.Data.Objects[textLabel.AttachedObject].Position
			textLabel.Position = objPos
			s.core.Data.TextLabels[textLabelID] = textLabel
		}
	}