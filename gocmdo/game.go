package gocmdo

import (
	"encoding/json"
	"fmt"
	"github.com/ingar/igo"
)

type Game struct {
	id   string
	Game *igo.Game
}

func (self Game) String() (s string) {
	var board *igo.Board
	var err error

	if board, err = self.Game.Board(); err != nil {
		return err.Error()
	}

	pointStateToSymbol := map[igo.PointState]string{
		igo.POINT_STATE_BLACK: "B",
		igo.POINT_STATE_WHITE: "W",
		igo.POINT_STATE_BLANK: ".",
	}

	s += fmt.Sprintf("%v\n", self.Game)
	s += "```"
	for row := 18; row >= 0; row-- {
		s += fmt.Sprintf("%2d ", row)
		for col := 0; col < 19; col++ {
			ps, err := board.IntersectionState(igo.Coordinates{col, row})
			if err != nil {
				return fmt.Sprintf("%v", err)
			}
			s += fmt.Sprintf(" %v ", pointStateToSymbol[ps])
		}
		s += "\n"
	}
	s += "    0  1  2  3  4  5  6  7  8  9  10 11 12 13 14 15 16 17 18\n"
	s += "```"
	return
}

func (self Game) Json() (out string) {
	if bytes, err := json.Marshal(self); err == nil {
		out = string(bytes)
	}
	return
}
