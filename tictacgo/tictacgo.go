package tictacgo

import (
	"errors"
	"fmt"
)

const (
	boardMask uint16 = 0x01FF     // board layout: 876
	                             //               543
	                             //               210
	xoMask uint16 = 0x0200
)

var wins = [...]uint16 {
	0b111_000_000,
	0b000_111_000,
	0b000_000_111,
	0b100_100_100,
	0b010_010_010,
	0b001_001_001,
	0b100_010_001,
	0b001_010_100,
}

type Player struct {
	Name string
	current uint16
	winMask uint16
}

func InitPlayer (p *Player, n string, isx bool){
	p.Name = n
	if isx {
		p.current = xoMask
	} else {
		p.current = 0
	}
	p.winMask = 0
}

func Winner (p1, p2 *Player) string {
	if boardMask & (p1.current | p2.current) == boardMask {
		return "draw"
	}
	for _, w := range wins {
		if (p1.current & w) == w {
			p1.winMask = w
			return p1.Name
		} else if (p2.current & w) == w {
			p2.winMask = w
			return p2.Name
		}
	}
    return "none"
}

func Move (p1, p2 *Player, cell uint) error {
	if cell > 9 {
		return errors.New("position out of bounds")
	}
	var pos uint16 = 1 << cell
	isOpen := (p1.current | p2.current) & pos == 0
	if isOpen {
		p1.current |= pos
	} else {
		return errors.New("position is already occupied")
	}
	return nil
}

func RenderBoard (p1, p2 Player){
	p1glyph := "X"
	p2glyph := "O"
	if p1.current & xoMask != xoMask {
		p1glyph = "O"
		p2glyph = "X"
	}
	wm := p1.winMask
	if p2.winMask != 0 {
		wm = p2.winMask
	}
    for rm := 0; rm < 3; rm++ {
    	cells := make( []string, 3, 3)
	    for cm := 0; cm < 3; cm++ {
	    	pos := 3 * rm + cm
	    	var cellMask uint16 = 1 << pos
	    	if wm != 0 {
				if wm & cellMask == cellMask {
					cells[cm] = "W"
				} else {
					cells[cm] = " "
				}
			} else {
				if p1.current&cellMask == cellMask {
					cells[cm] = p1glyph
				} else if p2.current&cellMask == cellMask {
					cells[cm] = p2glyph
				} else {
					cells[cm] = fmt.Sprintf("%d", pos+1)
				}
			}
		}
		fmt.Printf(" %s | %s | %s\n", cells[0], cells[1], cells[2])
	    fmt.Println("---+---+---")
	}

}