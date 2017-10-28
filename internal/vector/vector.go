package vector

type Vector struct {
	X, Y, Z float64
}

func (vec Vector) Multi(val float64) Vector {
	return Vector{
		vec.X * val,
		vec.Y * val,
		vec.Z * val}
}

func (vec Vector) Plus(vector Vector) Vector {
	return Vector{
		vec.X + vector.X,
		vec.Y + vector.Y,
		vec.Z + vector.Z}
}
