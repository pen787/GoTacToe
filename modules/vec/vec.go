package vec

import "math"

type Vec2 struct {
	X float32
	Y float32
}

func (v *Vec2) Add(x *Vec2) Vec2 {
	return Vec2{v.X + x.X, v.Y + v.Y}
}

func (v *Vec2) Sub(x *Vec2) Vec2 {
	return Vec2{v.X - x.X, v.Y - v.Y}
}

func (v *Vec2) Mul(x *Vec2) Vec2 {
	return Vec2{v.X * x.X, v.Y * x.Y}
}

func (v *Vec2) Scale(x float32) {
	v.X *= x
	v.Y *= x
}

func (v *Vec2) Invert() {
	v.Scale(-1)
}

func (v *Vec2) Distance(t *Vec2) float64 {
	return math.Sqrt(float64((v.X - t.X) + (v.X - t.Y)))
}
