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

import "math"

// Vector3 represents a position or rotation in 3D space
type Vector3 struct {
	X float32
	Y float32
	Z float32
}

// NewVector3 creates a new Vector3
func NewVector3(x, y, z float32) Vector3 {
	return Vector3{X: x, Y: y, Z: z}
}

// Add adds two vectors
func (v Vector3) Add(other Vector3) Vector3 {
	return Vector3{
		X: v.X + other.X,
		Y: v.Y + other.Y,
		Z: v.Z + other.Z,
	}
}

// Sub subtracts two vectors
func (v Vector3) Sub(other Vector3) Vector3 {
	return Vector3{
		X: v.X - other.X,
		Y: v.Y - other.Y,
		Z: v.Z - other.Z,
	}
}

// Mul multiplies a vector by a scalar
func (v Vector3) Mul(scalar float32) Vector3 {
	return Vector3{
		X: v.X * scalar,
		Y: v.Y * scalar,
		Z: v.Z * scalar,
	}
}

// DistanceTo calculates the distance to another vector
func (v Vector3) DistanceTo(other Vector3) float32 {
	dx := v.X - other.X
	dy := v.Y - other.Y
	dz := v.Z - other.Z
	return float32(math.Sqrt(float64(dx*dx + dy*dy + dz*dz)))
}

// DistanceToSquared calculates the squared distance (more efficient)
func (v Vector3) DistanceToSquared(other Vector3) float32 {
	dx := v.X - other.X
	dy := v.Y - other.Y
	dz := v.Z - other.Z
	return dx*dx + dy*dy + dz*dz
}

// Length returns the magnitude of the vector
func (v Vector3) Length() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z)))
}

// Normalize returns a normalized vector (magnitude = 1)
func (v Vector3) Normalize() Vector3 {
	length := v.Length()
	if length == 0 {
		return Vector3{}
	}
	return Vector3{
		X: v.X / length,
		Y: v.Y / length,
		Z: v.Z / length,
	}
}

// IsInRange checks if the vector is within a range of another vector
func (v Vector3) IsInRange(other Vector3, range_ float32) bool {
	return v.DistanceToSquared(other) <= range_*range_
}

// Zero returns a zero vector (0,0,0)
func Zero() Vector3 {
	return Vector3{0, 0, 0}
}

// Forward returns a vector that points forward (0,1,0)
func Forward() Vector3 {
	return Vector3{0, 1, 0}
}

// Up returns a vector that points up (0,0,1)
func Up() Vector3 {
	return Vector3{0, 0, 1}
}

// Right returns a vector that points right (1,0,0)
func Right() Vector3 {
	return Vector3{1, 0, 0}
}
