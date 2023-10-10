package player

import (
	"github.com/Sawansunar56/flappy-bird-golang/src/utils"
)

type Player struct {
	xPosition int
	yPosition int
}

func (p *Player) SetValues(x, y int) {
	p.xPosition = x
	p.yPosition = y
}

func (p *Player) GetValues() (x, y int) {
	x = p.xPosition
	y = p.yPosition
	return x, y
}

func (p *Player) UpMovement(display [utils.Wide][utils.High]rune) ([utils.Wide][utils.High]rune, bool) {
	collision := true
	display[p.xPosition][p.yPosition] = ' '
	if p.yPosition > 0 {
		p.yPosition -= 1
		if display[p.xPosition][p.yPosition] != ' ' {
			collision = false
		}
	}
	display[p.xPosition][p.yPosition] = '󱗆'

	return display, collision
}

func (p *Player) DownMovement(display [utils.Wide][utils.High]rune, height int) ([utils.Wide][utils.High]rune, bool) {
	collision := true
	display[p.xPosition][p.yPosition] = ' '

	if p.yPosition < height-1 {
		p.yPosition += 1

		if display[p.xPosition][p.yPosition] != ' ' {
			collision = false
		}
	}
	display[p.xPosition][p.yPosition] = '󱗆'

	return display, collision
}

func (p *Player) Init(display [utils.Wide][utils.High]rune) [utils.Wide][utils.High]rune {
	display[p.xPosition][p.yPosition] = '󱗆'
	return display
}

func (p *Player) Forward(display [utils.Wide][utils.High]rune) [utils.Wide][utils.High]rune {
	display[p.xPosition][p.yPosition] = '󱗆'
	display[p.xPosition-1][p.yPosition] = ' '
	return display
}

