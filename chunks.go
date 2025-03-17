package streamer

// ChunkStreamer maneja el streaming por pedazos para mejorar el rendimiento
type ChunkStreamer struct {
	core                  *Core
	chunkSize             map[int]int
	chunkStreamingEnabled bool
	streamInCallbacks     []StreamCallback
	streamOutCallbacks    []StreamCallback
}

// NewChunkStreamer crea un nuevo streamer de chunks
func NewChunkStreamer(core *Core) *ChunkStreamer {
	cs := &ChunkStreamer{
		core:               core,
		chunkSize:          make(map[int]int),
		streamInCallbacks:  make([]StreamCallback, 0),
		streamOutCallbacks: make([]StreamCallback, 0),
	}

	// Valores por defecto
	cs.chunkSize[StreamerTypeObject] = 10
	cs.chunkSize[StreamerTypeMapIcon] = 10
	cs.chunkSize[StreamerType3DTextLabel] = 10

	return cs
}

// StreamMapIcons implementa el streaming de iconos de mapa para un jugador
// StreamMapIcons transmite íconos del mapa para un jugador
func (c *ChunkStreamer) StreamMapIcons(player *Player, automatic bool) {
	// Actualizar solo si no es automático o si toca actualizar según el tick rate
	shouldProcess := !automatic
	if automatic {
		player.ChunkTickCount[StreamerTypeMapIcon]++
		shouldProcess = player.ChunkTickCount[StreamerTypeMapIcon] >= player.ChunkTickRate[StreamerTypeMapIcon]
	}
	if shouldProcess {
		chunkCount := 0

		// Procesar íconos que deben eliminarse
		if len(player.RemovedMapIcons) > 0 {
			// Convertir a slice para poder modificarlo durante la iteración
			removedIcons := make([]int, 0, len(player.RemovedMapIcons))
			for iconID := range player.RemovedMapIcons {
				removedIcons = append(removedIcons, iconID)
			}

			for _, iconID := range removedIcons {
				// Limitar el procesamiento por chunk si es automático
				shouldProcess = automatic
				if automatic {
					chunkCount++
					shouldProcess = chunkCount > c.chunkSize[StreamerTypeMapIcon]
				}
				if shouldProcess {
					break
				}

				// Buscar el ícono en los íconos internos del jugador
				if internalID, exists := player.InternalMapIcons[iconID]; exists {
					// Eliminar el ícono del cliente
					//c.core.SAMP.RemovePlayerMapIcon(player.PlayerId, internalID)

					// Verificar callbacks de streaming
					if mapIcon, exists := c.core.Data.MapIcons[iconID]; exists && mapIcon.StreamCallbacks {
						c.streamOutCallbacks = append(c.streamOutCallbacks,
							StreamCallback{
								ItemType: StreamerTypeMapIcon,
								ItemID:   iconID,
								PlayerID: player.PlayerId,
							})
					}

					// Liberar el identificador y eliminar de los íconos internos
					player.MapIconIdentifier.Remove(internalID, len(player.InternalMapIcons))
					delete(player.InternalMapIcons, iconID)
				}

				// Eliminar de la lista de íconos a eliminar
				delete(player.RemovedMapIcons, iconID)
			}
		} else {
			// Procesar íconos descubiertos
			for _, item := range player.DiscoveredMapIcons.items {
				priority := item.Priority
				distance := item.Distance
				mapIconID := item.ItemID
				mapIcon := item.ItemPtr

				// Limitar el procesamiento por chunk si es automático
				shouldProcess = automatic
				if automatic {
					chunkCount++
					shouldProcess = chunkCount > c.chunkSize[StreamerTypeMapIcon]
				}

				if shouldProcess {
					break
				}

				// Verificar si el ícono ya está en los íconos internos
				if _, exists := player.InternalMapIcons[mapIconID]; exists {
					player.DiscoveredMapIcons.Remove(mapIconID)
					continue
				}

				// Si alcanzamos el máximo de íconos visibles, intentar reemplazar uno menos prioritario
				if len(player.InternalMapIcons) == player.MaxVisibleMapIcons {
					// Buscar el ícono menos prioritario entre los existentes
					if !player.ExistingMapIcons.IsEmpty() {
						lastPriority, lastDistance, lastIconID, lastIcon := player.ExistingMapIcons.GetLast()

						// Reemplazar si el nuevo tiene mayor prioridad o está más cerca
						if lastPriority < priority || (lastDistance > StreamerStaticDistanceCutoff &&
							distance < lastDistance) {

							// Eliminar el ícono menos prioritario
							if internalID, exists := player.InternalMapIcons[lastIconID]; exists {
								//c.core.SAMP.RemovePlayerMapIcon(player.PlayerId, internalID)

								if mapIcon, ok := lastIcon.(*MapIcon); ok && mapIcon.StreamCallbacks {
									c.streamOutCallbacks = append(c.streamOutCallbacks,
										StreamCallback{
											ItemType: StreamerTypeMapIcon,
											ItemID:   lastIconID,
											PlayerID: player.PlayerId,
										})
								}

								player.MapIconIdentifier.Remove(internalID, len(player.InternalMapIcons))
								delete(player.InternalMapIcons, lastIconID)
							}

							// Limpiar de las estructuras
							if mapIcon, ok := lastIcon.(*MapIcon); ok && mapIcon.Cell != nil {
								delete(player.VisibleCell.MapIcons, lastIconID)
							}
							player.ExistingMapIcons.Remove(lastIconID)
						}
					}

					// Si aún estamos en el límite, no podemos añadir más íconos
					if len(player.InternalMapIcons) == player.MaxVisibleMapIcons {
						player.DiscoveredMapIcons.Clear()
						break
					}
				}

				// Obtener un nuevo identificador para el ícono
				internalID := player.MapIconIdentifier.Get()

				// Crear el ícono para el jugador
				/*c.core.SAMP.SetPlayerMapIcon(
					player.PlayerId,
					internalID,
					mapIcon.Position.X,
					mapIcon.Position.Y,
					mapIcon.Position.Z,
					mapIcon.Type,
					mapIcon.Color,
					mapIcon.Style,
				)*/

				// Notificar stream in si hay callbacks
				if mapIcon, ok := mapIcon.(*MapIcon); ok && mapIcon.StreamCallbacks {
					c.streamInCallbacks = append(c.streamInCallbacks,
						StreamCallback{
							ItemType: StreamerTypeMapIcon,
							ItemID:   mapIconID,
							PlayerID: player.PlayerId,
						})
				}

				// Registrar el ícono en los internos del jugador
				player.InternalMapIcons[mapIconID] = internalID

				// Añadir a los íconos existentes
				player.ExistingMapIcons.Insert(priority, distance, mapIconID, mapIcon)

				// Eliminar de los descubiertos
				player.DiscoveredMapIcons.Remove(mapIconID)
			}
		}

		// Reiniciar el contador de ticks para las próximas actualizaciones
		if automatic {
			player.ChunkTickCount[StreamerTypeMapIcon] = 0
		}
	}

	// Limpiar estructuras si no hay más íconos por procesar
	if player.DiscoveredMapIcons.IsEmpty() && len(player.RemovedMapIcons) == 0 {
		player.ExistingMapIcons.Clear()
		player.ProcessingChunks.Reset(StreamerTypeMapIcon)
	}
}

func (c *ChunkStreamer) PerformPlayerChunkUpdate(player *Player, automatic bool) {
	// Iterar por tipos según prioridad definida en core
	for _, itemType := range c.core.Data.TypePriority {
		switch itemType {
		case StreamerTypeObject:
			if player.ProcessingChunks.IsActive(StreamerTypeObject) {
				c.StreamObjects(player, automatic)
			}
		case StreamerTypeMapIcon:
			if player.ProcessingChunks.IsActive(StreamerTypeMapIcon) {
				c.StreamMapIcons(player, automatic)
			}
		case StreamerType3DTextLabel:
			if player.ProcessingChunks.IsActive(StreamerType3DTextLabel) {
				c.StreamTextLabels(player, automatic)
			}
		}
	}
}
