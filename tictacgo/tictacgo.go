package tictacgo

import (
	"errors"
	"fmt"
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

type Game struct {
	xplayer *Player
	oplayer *Player
}

func (g *Game) SetNames (xp, op string) {
	g.xplayer = &Player {xp, 0, 0 }
	g.oplayer = &Player{op, 0, 0 }
}

func (g Game) PlayerName (isx bool) string {
	if isx {
		return g.xplayer.Name
	} else {
		return g.oplayer.Name
	}
}

func (g Game) Winner () string {
	possible := 8
	for _, w := range wins {
		xp := g.xplayer.current & w
		op := g.oplayer.current & w
		if xp == w {
			g.xplayer.winMask = w
			return g.xplayer.Name
		} else if op == w {
			g.oplayer.winMask = w
			return g.oplayer.Name
		} else if xp != 0 && op != 0{
			possible--
		}
	}
	if possible == 0 {
		return "draw"
	}
    return "none"
}

func (g *Game) Move (isx bool, cell uint) error {
	if cell > 9 {
		return errors.New("position out of bounds")
	}
	var pos uint16 = 1 << cell
	isOpen := (g.xplayer.current | g.oplayer.current) & pos == 0
	p := g.oplayer
	if isx {
		p = g.xplayer
	}
	if isOpen {
		p.current |= pos
	} else {
		return errors.New("position is already occupied")
	}
	return nil
}

func RenderBoard(g Game){
	xglyph := "X"
	oglyph := "O"
	wm := g.xplayer.winMask
	if g.oplayer.winMask != 0 {
		wm = g.oplayer.winMask
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
				if g.xplayer.current & cellMask == cellMask {
					cells[cm] = xglyph
				} else if g.oplayer.current & cellMask == cellMask {
					cells[cm] = oglyph
				} else {
					cells[cm] = fmt.Sprintf("%d", pos+1)
				}
			}
		}
		fmt.Printf(" %s | %s | %s\n", cells[0], cells[1], cells[2])
	    fmt.Println("---+---+---")
	}

}