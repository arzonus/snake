package main

import (
	"github.com/arzonus/snake/internal/game"
	"github.com/hajimehoshi/ebiten"
	"log"
)

func main() {
	g, err := game.New()
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetRunnableInBackground(true)

	if err := ebiten.Run(game.Update(g), game.ScreenWidth, game.ScreenHeight, 1, "Snake"); err != nil {
		log.Fatal(err)
	}
}
