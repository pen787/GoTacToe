package main

import (
	"errors"
	"image/color"
	"log"
	"math"
	"pen787/GoTacToe/modules/gameobject"
	"pen787/GoTacToe/modules/vec"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/tinne26/etxt"
	"github.com/tinne26/fonts/liberation/lbrtserif"
)

const WINDOWSIZE_X int = 500
const WINDOWSIZE_Y int = 500

var GameObjectContainer *map[[2]uint]*gameobject.Object
var CurrentRound int
var Winner int
var AlreadyClicked mapset.Set[[2]uint]
var isDone bool

var Font *text.Face

func setUpGrid() {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			var size = vec.Vec2{
				X: 100,
				Y: 100,
			}

			var GirdMaxSize = vec.Vec2{
				X: size.X * 3,
				Y: size.Y * 3,
			}

			var marigin float32 = 10
			var gridpos vec.Vec2 = vec.Vec2{
				X: (float32(WINDOWSIZE_X) / 2) - GirdMaxSize.X/2,
				Y: (float32(WINDOWSIZE_Y) / 2) - GirdMaxSize.Y/2,
			}

			name := [2]uint{uint(row), uint(col)}
			(*GameObjectContainer)[name] = gameobject.MakeObject(
				vec.Vec2{
					X: ((size.X + marigin) * float32(col)) + gridpos.X,
					Y: ((size.Y + marigin) * float32(row)) + gridpos.Y,
				},
				vec.Vec2{X: size.X, Y: size.Y},
				color.RGBA{255, 255, 255, 255},
			)
		}
	}
}

func init() {
	GameObjectContainer = &map[[2]uint]*gameobject.Object{}
	AlreadyClicked = mapset.NewSet[[2]uint]()
	Winner = -1
	CurrentRound = 0
	isDone = false

	// Font =

	// Set up a grid
	setUpGrid()
}

type Game struct {
	text *etxt.Renderer
}

// TODO: Rewrite this code. :)
// It just works.
func checkWinner() int {
	// 3 col in any row is equal
	for row := 0; row < 3; row++ {
		box1, ok1 := (*GameObjectContainer)[[2]uint{uint(row), uint(0)}]
		box2, ok2 := (*GameObjectContainer)[[2]uint{uint(row), uint(1)}]
		box3, ok3 := (*GameObjectContainer)[[2]uint{uint(row), uint(2)}]
		if ok1 && ok2 && ok3 {
			v1 := box1.Value
			v2 := box2.Value
			v3 := box3.Value

			if v1 == v2 && v2 == v3 && v1 != -1 && v2 != -1 && v3 != -1 {
				if v1+v2+v3 == 3 {
					return 1
				} else if v1+v2+v3 == 0 {
					return 0
				}
			}
		} else {
			log.Fatalf("False to get an Object")
		}
	}

	// 3 row in any col is equal
	for col := 0; col < 3; col++ {
		box1, ok1 := (*GameObjectContainer)[[2]uint{uint(0), uint(col)}]
		box2, ok2 := (*GameObjectContainer)[[2]uint{uint(1), uint(col)}]
		box3, ok3 := (*GameObjectContainer)[[2]uint{uint(2), uint(col)}]
		if ok1 && ok2 && ok3 {
			v1 := box1.Value
			v2 := box2.Value
			v3 := box3.Value

			if v1 == v2 && v2 == v3 && v1 != -1 && v2 != -1 && v3 != -1 {
				if v1+v2+v3 == 3 {
					return 1
				} else if v1+v2+v3 == 0 {
					return 0
				}
			}
		} else {
			log.Fatalf("False to get an Object")
		}
	}

	// welp
	box1, ok1 := (*GameObjectContainer)[[2]uint{0, 0}]
	box2, ok2 := (*GameObjectContainer)[[2]uint{1, 1}]
	box3, ok3 := (*GameObjectContainer)[[2]uint{2, 2}]
	if ok1 && ok2 && ok3 {
		v1 := box1.Value
		v2 := box2.Value
		v3 := box3.Value

		if v1 == v2 && v2 == v3 && v1 != -1 && v2 != -1 && v3 != -1 {
			if v1+v2+v3 == 3 {
				return 1
			} else if v1+v2+v3 == 0 {
				return 0
			}
		}
	} else {
		log.Fatalf("False to get an Object")
	}

	box4, ok4 := (*GameObjectContainer)[[2]uint{0, 2}]
	box5, ok5 := (*GameObjectContainer)[[2]uint{2, 0}]
	if ok4 && ok2 && ok5 {
		v1 := box4.Value
		v2 := box2.Value
		v3 := box5.Value

		if v1 == v2 && v2 == v3 && v1 != -1 && v2 != -1 && v3 != -1 {
			if v1+v2+v3 == 3 {
				return 1
			} else if v1+v2+v3 == 0 {
				return 0
			}
		}
	} else {
		log.Fatalf("False to get an Object")
	}

	sum := 0
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			box, ok := (*GameObjectContainer)[[2]uint{uint(row), uint(col)}]
			if !ok {
				log.Fatalf("False to get an object : ", row, col)
			}
			if box.Value == 0 || box.Value == 1 {
				sum++
			}
		}
	}
	if sum >= 9 {
		return 2
	}

	return -1
}

func (g *Game) Update() error {

	// When mouse clicked
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) && !isDone {
		for row := 0; row < 3; row++ {
			for col := 0; col < 3; col++ {
				box, ok := (*GameObjectContainer)[[2]uint{uint(row), uint(col)}]
				if !ok {
					log.Fatalf("False to get an object : ", row, col)
					return errors.New("failed to get an object")
				}
				if box.IsMouseInside() && !AlreadyClicked.Contains([2]uint{uint(row), uint(col)}) {
					if CurrentRound == 0 {
						box.Color = &color.RGBA{255, 0, 0, 255}
						box.Value = 0
						CurrentRound = 1
					} else if CurrentRound == 1 {
						box.Color = &color.RGBA{0, 255, 0, 255}
						box.Value = 1
						CurrentRound = 0
					}
					AlreadyClicked.Add([2]uint{uint(row), uint(col)})
				}
			}
		}
		Winner = checkWinner()
		if Winner == 1 || Winner == 0 || Winner == 2 {
			isDone = true
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{100, 100, 100, 25})

	if !isDone {
		if CurrentRound == 0 {
			g.text.Draw(screen, "Red to go.", WINDOWSIZE_X/2, 30)
		} else if CurrentRound == 1 {
			g.text.Draw(screen, "Green to go.", WINDOWSIZE_X/2, 30)
		}
	} else {
		if Winner == 0 {
			g.text.Draw(screen, "Red Win.", WINDOWSIZE_X/2, 30)
		} else if Winner == 1 {
			g.text.Draw(screen, "Green Win.", WINDOWSIZE_X/2, 30)
		} else {
			g.text.Draw(screen, "Tie.", WINDOWSIZE_X/2, 30)
		}
	}

	for _, Object := range *GameObjectContainer {
		Object.Render(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	scale := ebiten.Monitor().DeviceScaleFactor()
	g.text.SetScale(scale)

	canvasWidth := int(math.Ceil(float64(WINDOWSIZE_X) * scale))
	canvasHeight := int(math.Ceil(float64(WINDOWSIZE_Y) * scale))

	return canvasWidth, canvasHeight
}

func main() {
	renderer := etxt.NewRenderer()
	renderer.SetFont(lbrtserif.Font())
	renderer.Utils().SetCache8MiB()

	renderer.SetColor(color.RGBA{255, 255, 255, 255})
	renderer.SetAlign(etxt.Center)
	renderer.SetSize(32)

	ebiten.SetWindowSize(WINDOWSIZE_X, WINDOWSIZE_Y)
	ebiten.SetWindowTitle("TicTacGo")
	log.Println("Program Started!")
	if err := ebiten.RunGame(&Game{text: renderer}); err != nil {
		log.Fatal(err)
	}

}
