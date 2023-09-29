package game

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/gdamore/tcell/v2"
)

const (
	wide int = 160
	high int = 50
)

func randomNumberGenerator(min, max int, screen tcell.Screen) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 100; i++ {
		randomNumber := r.Intn(max-min+1) + min
		screen.SetContent(20+i, 30, rune(randomNumber), nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
	}
}

func renderTopBox(screen tcell.Screen, x, y int, display [wide][high]rune) [wide][high]rune {
  bottomLeft := '└'
	bottomRight := '┘'
	horizontalLine := '─'
	verticalLine := '│'
  display[10][10] = 'x'
	for i := 1; i <= y; i++ {
		display[x][y-i] = verticalLine
		// screen.SetContent(x, y-i, verticalLine, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
	}

	display[x][y] = bottomLeft
	// screen.SetContent(x, y, bottomLeft, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
	for i := 1; i <= 10; i++ {
		display[x+i][y] = horizontalLine
		// screen.SetContent(x+i, y, horizontalLine, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
	}
	x = x + 10
	display[x][y] = bottomRight
	// screen.SetContent(x, y, bottomRight, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
	for i := 1; i <= y; i++ {
		display[x][y-i] = verticalLine
		// screen.SetContent(x, y-i, verticalLine, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
	}
  return display
}

func renderBottomBox(screen tcell.Screen, x, y int, display [wide][high]rune) [wide][high]rune {
	_, height := screen.Size()
	topLeft := '┌'
	topRight := '┐'
	horizontalLine := '─'
	verticalLine := '│'

	for i := height - 1; i > y+10; i-- {
		display[x][i] = verticalLine
		// screen.SetContent(x, i, verticalLine, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
	}
	y = y + 10
	display[x][y] = topLeft
	// screen.SetContent(x, y, topLeft, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
	for i := 1; i <= 10; i++ {
		display[x+i][y] = horizontalLine
		// screen.SetContent(x+i, y, horizontalLine, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
	}
	x = x + 10
	display[x][y] = topRight
	// screen.SetContent(x, y, topRight, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
	for i := height - 1; i > y; i-- {
		display[x][i] = verticalLine
		// screen.SetContent(x, i, verticalLine, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
	}
  return display
}

func platform(screen tcell.Screen, x, y int) {
}

func Game(screen tcell.Screen) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	width, height := screen.Size()
	var display [wide][high]rune

	text := strconv.Itoa(height) + " " + strconv.Itoa(width)
	x := width/2 - len(text)/2
	y := height / 2

	for i, char := range text {
		screen.SetContent(x+i, y, char, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
	}
	// screen.SetContent(x, 10, 'H', nil, tcell.StyleDefault.Foreground(tcell.ColorHotPink))
	// platform(screen, 5, 10)
	// i, j := 5, 10

	for i := 1; i <= 3; i++ {
		randomNumber := r.Intn(30-3+1) + 3
		display = renderTopBox(screen, 30*i, randomNumber, display)
		display = renderBottomBox(screen, 30*i, randomNumber, display)
	}

  // display[154][0] = 'h'
	// display = renderTopBox(screen, i, j, display)
	// display = renderBottomBox(screen, i, j, display)
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			screen.SetContent(i, j, display[i][j], nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
		}
	}

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
