package streamer

import "math"

// Vector2 representa un vector 2D con componentes float
type Vector2 struct {
	X, Y float32
}

// Vector3 representa un vector 3D con componentes float
type Vector3 struct {
	X, Y, Z float32
}

// IsZero verifica si el vector es (0,0,0)
func (v Vector3) IsZero() bool {
	return v.X == 0 && v.Y == 0 && v.Z == 0
}

// SquaredNorm calcula la norma al cuadrado del vector
func (v Vector3) SquaredNorm() float32 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// Add suma otro vector
func (v Vector3) Add(other Vector3) Vector3 {
	return Vector3{
		X: v.X + other.X,
		Y: v.Y + other.Y,
		Z: v.Z + other.Z,
	}
}

// Sub resta otro vector
func (v Vector3) Sub(other Vector3) Vector3 {
	return Vector3{
		X: v.X - other.X,
		Y: v.Y - other.Y,
		Z: v.Z - other.Z,
	}
}

// Mul multiplica por un escalar
func (v Vector3) Mul(scalar float32) Vector3 {
	return Vector3{
		X: v.X * scalar,
		Y: v.Y * scalar,
		Z: v.Z * scalar,
	}
}

// Distance calcula la distancia euclidiana entre dos puntos
func Distance(a, b Vector3) float32 {
	dx := a.X - b.X
	dy := a.Y - b.Y
	dz := a.Z - b.Z
	return float32(math.Sqrt(float64(dx*dx + dy*dy + dz*dz)))
}

// ComparableDistance calcula la distancia al cuadrado (para comparaciones eficientes)
func ComparableDistance(a, b Vector3) float32 {
	dx := a.X - b.X
	dy := a.Y - b.Y
	dz := a.Z - b.Z
	return dx*dx + dy*dy + dz*dz
}

// Box2d representa una caja 2D
type Box2d struct {
	Min, Max Vector2
}

// Box3d representa una caja 3D
type Box3d struct {
	Min, Max Vector3
}

// Polygon2d representa un polígono 2D
type Polygon2d struct {
	Points []Vector2
}
