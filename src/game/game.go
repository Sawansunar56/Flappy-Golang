package game

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/gdamore/tcell/v2"
)

func randomNumberGenerator(min, max int, screen tcell.Screen) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 100; i++ {
		randomNumber := r.Intn(max-min+1) + min
		screen.SetContent(20+i, 30, rune(randomNumber), nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
	}
}

func renderTopBox(screen tcell.Screen, x, y int) {
	bottomLeft := '└'
	bottomRight := '┘'
	horizontalLine := '─'
	verticalLine := '│'
	for i := 1; i < 10; i++ {
		screen.SetContent(x, y-i, verticalLine, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
	}

	screen.SetContent(x, y, bottomLeft, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
	for i := 1; i <= 10; i++ {
		screen.SetContent(x+i, y, horizontalLine, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
	}
	x = x + 10
	screen.SetContent(x, y, bottomRight, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
	for i := 1; i < 10; i++ {
		screen.SetContent(x, y-i, verticalLine, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
	}
}

func renderBottomBox(screen tcell.Screen, x, y int) {
	_, height := screen.Size()
	topLeft := '┌'
	topRight := '┐'
	horizontalLine := '─'
	verticalLine := '│'

	for i := height - 5; i > y+10; i-- {
		screen.SetContent(x, i, verticalLine, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
	}
	y = y + 10
	screen.SetContent(x, y, topLeft, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
	for i := 1; i <= 10; i++ {
		screen.SetContent(x+i, y, horizontalLine, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
	}
	x = x + 10
	screen.SetContent(x, y, topRight, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
	for i := height - 5; i > y; i-- {
		screen.SetContent(x, i, verticalLine, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
	}
}

func platform(screen tcell.Screen, x, y int) {
}

func Game(screen tcell.Screen) {
	// Example: Display "Hello, TCell!" in the center of the screen
	width, height := screen.Size()

	text := strconv.Itoa(height) + " " +  strconv.Itoa(width)
	x := width/2 - len(text)/2
	y := height / 2

	for i, char := range text {
		screen.SetContent(x+i, y, char, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
	}
	// screen.SetContent(x, 10, 'H', nil, tcell.StyleDefault.Foreground(tcell.ColorHotPink))
	// platform(screen, 5, 10)

  renderTopBox(screen, 5, 10)
  renderBottomBox(screen, 5, 10)

	screen.Show()

	for {
		ev := screen.PollEvent()
		switch ev.(type) {
		case *tcell.EventKey:
			// Handle key events
			keyEvent := ev.(*tcell.EventKey)
			if keyEvent.Key() == tcell.KeyCtrlC {
				return // Exit the application on Ctrl+C
			} else if keyEvent.Rune() == 'i' {
				screen.Clear()
				screen.Show()
			}
		}
	}
}
