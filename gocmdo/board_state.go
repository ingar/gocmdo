package gocmdo

import (
	"fmt"
)

type BoardGrid [19 * 19]string

func NewBoardGrid(game *Game) (b BoardGrid) {
	b = BoardGrid{}

	for idx, move := range game.Moves {
		s := "W"
		if idx%2 == 0 {
			s = "B"
		}
		b[19*move.row+move.column] = s
	}
	return
}

func (self BoardGrid) Ascii() (s string) {
	s = "```"

	for row := 18; row >= 0; row-- {
		s += fmt.Sprintf("%2d ", row)
		for col := 0; col < 19; col++ {
			val := self[row*19+col]
			if len(val) == 0 {
				val = "."
			}
			s += fmt.Sprintf(" %s ", val)
		}
		s += "\n"
	}
	s += "    0  1  2  3  4  5  6  7  8  9  10 11 12 13 14 15 16 17 18\n"
	s += "```"
	return
}
