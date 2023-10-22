package main

import (
	"os"

	"github.com/Sawansunar56/flappy-bird-golang/src/game"
	// "github.com/Sawansunar56/flappy-bird-golang/src/menu"
	"github.com/gdamore/tcell/v2"
)

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		// Handle error
		os.Exit(1)
	}

	err = screen.Init()
	if err != nil {
		// Handle error
		os.Exit(1)
	}
	defer screen.Fini()

	// Create a simple TUI here

    // menu.Menu(screen)
	game.Game(screen)
}
