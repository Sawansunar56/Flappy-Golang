package menu

import (
	"github.com/Sawansunar56/flappy-bird-golang/src/utils"
	"github.com/gdamore/tcell/v2"
)

// Renders the outer border of the menu screen
func menuBrackets(screen tcell.Screen) {
	screen.SetContent(0, 0, utils.TopLeft, nil, tcell.StyleDefault.Foreground(tcell.Color101))
	for i := 1; i < utils.ScreenWidth; i++ {
		screen.SetContent(
			i,
			0,
			utils.HorizontalLine,
			nil,
			tcell.StyleDefault.Foreground(tcell.Color101),
		)
		screen.SetContent(
			i,
			utils.ScreenHeight,
			utils.HorizontalLine,
			nil,
			tcell.StyleDefault.Foreground(tcell.Color101),
		)
	}
	screen.SetContent(
		utils.ScreenWidth,
		0,
		utils.TopRight,
		nil,
		tcell.StyleDefault.Foreground(tcell.Color101),
	)
	for i := 1; i < utils.ScreenHeight; i++ {
		screen.SetContent(
			0,
			i,
			utils.VerticalLine,
			nil,
			tcell.StyleDefault.Foreground(tcell.Color101),
		)
		screen.SetContent(
			utils.ScreenWidth,
			i,
			utils.VerticalLine,
			nil,
			tcell.StyleDefault.Foreground(tcell.Color101),
		)
	}
	screen.SetContent(
		0,
		utils.ScreenHeight,
		utils.BottomLeft,
		nil,
		tcell.StyleDefault.Foreground(tcell.Color101),
	)
	screen.SetContent(
		utils.ScreenWidth,
		utils.ScreenHeight,
		utils.BottomRight,
		nil,
		tcell.StyleDefault.Foreground(tcell.Color101),
	)
}

// PRINTS: HELLO
func header(screen tcell.Screen) {
    // Currently I believe the best way to print a gigantic header is by making
    // string that holds it in and then rendering the string line by line.
}

// Main game Menu
func Menu(screen tcell.Screen) {
	menuBrackets(screen)
	header(screen)
	text := "Hello There"
	x := utils.ScreenWidth/2 - len(text)
	y := utils.ScreenHeight / 2
	for _, i := range text {
		screen.SetContent(x, y, i, nil, tcell.StyleDefault.Foreground(tcell.Color101))
		x += 1
	}

	screen.Show()

	for {
		event := screen.PollEvent()
		switch event.(type) {
		case *tcell.EventKey:
			keyEvent := event.(*tcell.EventKey)

			if keyEvent.Key() == tcell.KeyCtrlC {
				return
			}
		}

	}
}
