package main

import "github.com/hajimehoshi/ebiten/v2"

type Game struct{}

// Update вызывается на каждый tick
func (g *Game) Update() error {
	return nil
}

// Draw вызывается на каждый frame
func (g *Game) Draw(screen *ebiten.Image) {

}

// Layout is user for scaling the game screen
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {

	g := &Game{}
	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
