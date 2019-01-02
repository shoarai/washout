// Copyright © 2018 shoarai

package washout

// A Vector is a vector in 3D axis.
type Vector struct {
	X, Y, Z float64
}

// Plus adds a vector.
func (vec Vector) Plus(vector Vector) Vector {
	return Vector{
		vec.X + vector.X,
		vec.Y + vector.Y,
		vec.Z + vector.Z}
}

// Multi multiples a value.
func (vec Vector) Multi(val float64) Vector {
	return Vector{
		vec.X * val,
		vec.Y * val,
		vec.Z * val}
}
