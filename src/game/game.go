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

// Aligns a new top box in the display
func renderTopBox(screen tcell.Screen, x, y int, display [wide][high]rune) [wide][high]rune {
	bottomLeft := '└'
	bottomRight := '┘'
	horizontalLine := '─'
	verticalLine := '│'

	for i := 1; i <= y; i++ {
		display[x][y-i] = verticalLine
	}

	display[x][y] = bottomLeft
	for i := 1; i <= 10; i++ {
		display[x+i][y] = horizontalLine
	}

	x = x + 10
	display[x][y] = bottomRight
	for i := 1; i <= y; i++ {
		display[x][y-i] = verticalLine
	}

	return display
}

// Aligns a new bottom box into the display
func renderBottomBox(screen tcell.Screen, x, y int, display [wide][high]rune) [wide][high]rune {
	_, height := screen.Size()
	topLeft := '┌'
	topRight := '┐'
	horizontalLine := '─'
	verticalLine := '│'

	for i := height - 1; i > y+10; i-- {
		display[x][i] = verticalLine
	}

	y = y + 10
	display[x][y] = topLeft
	for i := 1; i <= 10; i++ {
		display[x+i][y] = horizontalLine
	}

	x = x + 10
	display[x][y] = topRight
	for i := height - 1; i > y; i-- {
		display[x][i] = verticalLine
	}

	return display
}

func platform(screen tcell.Screen, x, y int) {
}

func renderNewBox(screen tcell.Screen, display [wide][high]rune, count int) [wide][high]rune {
	if count == 30 {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		randomNumber := r.Intn(30-3+1) + 3
		display = renderTopBox(screen, 130, randomNumber, display)
		display = renderBottomBox(screen, 130, randomNumber, display)
	}
	return display
}

func renderBack(display [wide][high]rune, width, height int) [wide][high]rune {
	for i := 0; i < width-1; i++ {
		for j := 0; j < height; j++ {
			display[i][j] = display[i+1][j]
		}
	}
	return display
}

func Game(screen tcell.Screen) {
	width, height := screen.Size()
	var display [wide][high]rune

	text := strconv.Itoa(height) + " " + strconv.Itoa(width)
	x := width/2 - len(text)/2
	y := height / 2

	for i, char := range text {
		screen.SetContent(x+i, y, char, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
	}

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			screen.SetContent(i, j, display[i][j], nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
		}
	}

	screen.Show()

	events := make(chan tcell.Event)

	go func() {
		for {
			ev := screen.PollEvent()
			events <- ev
		}
	}()

	count := 30

	for {
		select {
		case ev := <-events:
			switch ev.(type) {
			case *tcell.EventKey:
				// Handle key events
				keyEvent := ev.(*tcell.EventKey)
				if keyEvent.Key() == tcell.KeyCtrlC {
					return // Exit the application on Ctrl+C
				} else if keyEvent.Rune() == 'i' {
					screen.Clear()
					display = renderBack(display, width, height)
					for i := 0; i < width; i++ {
						for j := 0; j < height; j++ {
							screen.SetContent(i, j, display[i][j], nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
						}
					}
					screen.Show()
				} else if keyEvent.Rune() == 'x' {
					screen.Clear()
					display = renderNewBox(screen, display, count)

					for i := 0; i < width; i++ {
						for j := 0; j < height; j++ {
							screen.SetContent(i, j, display[i][j], nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
						}
					}
					screen.Show()
				}
			}
		default:
			screen.Clear()
			display = renderNewBox(screen, display, count)
			if count == 30 {
				count = 0
			}
			display = renderBack(display, width, height)
			for i := 0; i < width; i++ {
				for j := 0; j < height; j++ {
					screen.SetContent(i, j, display[i][j], nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
				}
			}
			screen.Show()
			count += 1
			time.Sleep(100 * time.Millisecond)
		}
	}
}
