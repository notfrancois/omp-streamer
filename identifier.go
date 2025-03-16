package streamer

// Identifier gestiona un conjunto de IDs reutilizables
type Identifier struct {
	ids        []int // Lista de IDs disponibles
	upperBound int   // Límite superior actual
	lowerBound int   // Límite inferior
	upperMax   int   // Límite superior máximo
}

// NewIdentifier crea un nuevo gestor de identificadores
func NewIdentifier(lowerBound, upperMax int) *Identifier {
	return &Identifier{
		ids:        make([]int, 0),
		upperBound: lowerBound,
		lowerBound: lowerBound,
		upperMax:   upperMax,
	}
}

// Get obtiene un identificador disponible
func (i *Identifier) Get() int {
	// Si hay IDs disponibles en la lista, usar el primero
	if len(i.ids) > 0 {
		id := i.ids[0]
		i.ids = i.ids[1:]
		return id
	}

	// Si no hay IDs disponibles, pero podemos incrementar el límite superior
	if i.upperBound < i.upperMax {
		id := i.upperBound
		i.upperBound++
		return id
	}

	// No hay IDs disponibles y no podemos incrementar el límite
	return -1
}

// Remove libera un identificador para su reutilización
func (i *Identifier) Remove(id int, currentSize int) {
	// Solo gestionar IDs dentro de los límites
	if id >= i.lowerBound && id < i.upperMax {
		// Si el ID es el límite superior - 1 y no hay elementos reutilizables
		// Podemos decrementar el límite superior
		if id == i.upperBound-1 && len(i.ids) == 0 {
			i.upperBound--
		} else {
			// Añadir el ID a la lista de disponibles
			i.ids = append(i.ids, id)
		}
	}
}

// Clear reinicia el identificador a
