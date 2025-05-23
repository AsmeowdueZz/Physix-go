package vector

import "math"

// Vector represents a 2D vector.
type Vector struct {
	X, Y float64
}

// Add performs vector addition.
func (v Vector) Add(other Vector) Vector {
	return Vector{v.X + other.X, v.Y + other.Y}
}
// NewVector creates a new vector with the given x and y components.
func NewVector(x, y float64) Vector {
	return Vector{X: x, Y: y}
}

// InnerProduct performs vector inner product or Dot product.
func (v Vector) InnerProduct(other Vector) float64 {
	return v.X * other.X+ v.Y * other.Y
}

// Subtract performs vector subtraction.
func (v Vector) Sub(other Vector) Vector {
	return Vector{v.X - other.X, v.Y - other.Y}
}

// Scale multiplies the vector by a scalar.
func (v Vector) Scale(scalar float64) Vector {
	return Vector{v.X * scalar, v.Y * scalar}
}

// Magnitude calculates the magnitude (length) of the vector.
func (v Vector) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Normalize normalizes the vector to have a magnitude of 1.
func (v Vector) Normalize() Vector {
	magnitude := v.Magnitude()
	if magnitude == 0 {
		return Vector{}
	}
	return v.Scale(1 / magnitude)
}


// Calculate distance between to 2 heads of Vector
func Distance(v1 Vector , v2 Vector) float64 {
	return math.Sqrt(     ((v2.X-v1.X)*(v2.X-v1.X))    +    ((v2.Y-v1.Y)*(v2.Y-v1.Y))    )
}


// Orthogonal computes the perpendicular vector of a given vector.
func Orthogonal(v Vector) Vector {
	return Vector{-v.Y, v.X}
}

func ComponentAlong(A, B Vector) Vector {
    if B.Magnitude() == 0 {
        return Vector{0, 0} // Handle zero-length vector case
    }
    scale := A.InnerProduct(B) / B.InnerProduct(B) // (A · B) / (|B|^2)
    return B.Scale(scale)        // Scale B to get projection
}
