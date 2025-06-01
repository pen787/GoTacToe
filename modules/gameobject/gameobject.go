package gameobject

import (
	"image/color"
	"pen787/GoTacToe/modules/vec"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Renderer interface {
	Render()
}

type Object struct {
	Position *vec.Vec2
	Size     *vec.Vec2
	Color    *color.RGBA
	Value    int
}

func (gobj *Object) IsPointInside(point *vec.Vec2) bool {
	xmin := gobj.Position.X
	xmax := xmin + gobj.Size.X
	ymin := gobj.Position.Y
	ymax := ymin + gobj.Size.Y

	if point.X <= xmax && point.X >= xmin && point.Y >= ymin && point.Y <= ymax {
		return true
	}
	return false
}

func (gobj *Object) IsMouseInside() bool {
	x, y := ebiten.CursorPosition()
	return gobj.IsPointInside(&vec.Vec2{X: float32(x), Y: float32(y)})
}

func (gobj *Object) Render(dist *ebiten.Image) {
	vector.DrawFilledRect(dist, gobj.Position.X, gobj.Position.Y, gobj.Size.X, gobj.Size.Y, gobj.Color, true)
}

func MakeObject(position vec.Vec2, size vec.Vec2, color color.RGBA) *Object {
	return &Object{
		Position: &position,
		Size:     &size,
		Color:    &color,
		Value:    -1,
	}
}
