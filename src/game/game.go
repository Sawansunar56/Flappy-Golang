package game

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/Sawansunar56/flappy-bird-golang/src/player"
	"github.com/Sawansunar56/flappy-bird-golang/src/utils"
	"github.com/gdamore/tcell/v2"
)

func randomNumberGenerator(min, max int, screen tcell.Screen) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 100; i++ {
		randomNumber := r.Intn(max-min+1) + min
		screen.SetContent(20+i, 30, rune(randomNumber), nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
	}
}

// Aligns a new top box in the display
func renderTopBox(screen tcell.Screen, x, y int, display [utils.Wide][utils.High]rune) [utils.Wide][utils.High]rune {
	for i := 1; i <= y; i++ {
		display[x][y-i] = utils.VerticalLine
	}

	display[x][y] = utils.BottomLeft
	for i := 1; i <= 10; i++ {
		display[x+i][y] = utils.HorizontalLine
	}

	x = x + 10
	display[x][y] = utils.BottomRight
	for i := 1; i <= y; i++ {
		display[x][y-i] = utils.VerticalLine
	}

	return display
}

// Aligns a new bottom box into the display
func renderBottomBox(screen tcell.Screen, x, y int, display [utils.Wide][utils.High]rune) [utils.Wide][utils.High]rune {
	_, height := screen.Size()

	for i := height - 1; i > y+10; i-- {
		display[x][i] = utils.VerticalLine
	}

	y = y + 10
	display[x][y] = utils.TopLeft
	for i := 1; i <= 10; i++ {
		display[x+i][y] = utils.HorizontalLine
	}

	x = x + 10
	display[x][y] = utils.TopRight
	for i := height - 1; i > y; i-- {
		display[x][i] = utils.VerticalLine
	}

	return display
}

func renderNewBox(screen tcell.Screen, display [utils.Wide][utils.High]rune, count int) [utils.Wide][utils.High]rune {
	if count == 35 {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		randomNumber := r.Intn(30-3+1) + 3
		display = renderTopBox(screen, 130, randomNumber, display)
		display = renderBottomBox(screen, 130, randomNumber, display)
	}
	return display
}

func renderBack(display [utils.Wide][utils.High]rune, width, height int) [utils.Wide][utils.High]rune {
	for i := 0; i < width-1; i++ {
		for j := 0; j < height; j++ {
			display[i][j] = display[i+1][j]
		}
	}
	return display
}

func Game(screen tcell.Screen) {
	width, height := screen.Size()
	var display [utils.Wide][utils.High]rune

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

	downBias := 1
	count := 30
	var player player.Player
	collisionCondition := true
	player.SetValues(20, 20)
	display = player.Init(display)

	for {
		select {
		case ev := <-events:
			switch ev.(type) {
			case *tcell.EventKey:
				keyEvent := ev.(*tcell.EventKey)

				if keyEvent.Key() == tcell.KeyCtrlC {
					return
				} else if keyEvent.Key() == tcell.KeyUp {
					screen.Clear()
					// IMPLEMENT THE COLLISION CONDITION HERE>
					display, collisionCondition = player.UpMovement(display)
					for i := 0; i < width; i++ {
						for j := 0; j < height; j++ {
							screen.SetContent(i, j, display[i][j], nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
						}
					}
					screen.Show()
                    if collisionCondition == false {
                        return
                    }
				} else if keyEvent.Key() == tcell.KeyDown {
					screen.Clear()
					display, collisionCondition = player.DownMovement(display, height)
					for i := 0; i < width; i++ {
						for j := 0; j < height; j++ {
							screen.SetContent(i, j, display[i][j], nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
						}
					}
					screen.Show()
                    // if collisionCondition {
                    //     return
                    // }
				}
			}
		default:
			screen.Clear()

			display = renderNewBox(screen, display, count)
			if count == 35 {
				count = 0
			}
			display = renderBack(display, width, height)
			display = player.Forward(display)
			if downBias == 7 {
				display, collisionCondition = player.DownMovement(display, height)
				downBias = 1
			}
			for i := 0; i < width; i++ {
				for j := 0; j < height; j++ {
					screen.SetContent(i, j, display[i][j], nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
				}
			}
			screen.Show()
			count += 1
			downBias += 1
			time.Sleep(100 * time.Millisecond)
		}
	}
}
