package gocmdo

import (
	"errors"
	"strconv"
	"github.com/ingar/igo"
	"fmt"
)

type Game struct {
	id string
	Game *igo.Game
}

func (self Game) String() (s string) {
	var board *igo.Board
	var err error

	if board, err = self.Game.Board(); err != nil {
		return err.Error()
	}

	pointStateToSymbol := map[igo.PointState]string {
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

// GamesRepo
type GamesRepo struct {
	Games []*Game
}
var GamesRepository GamesRepo

func (self *GamesRepo) NewGame(playerOneId string, playerTwoId string) *Game {
	g := Game{strconv.Itoa(len(self.Games) + 1), &igo.Game{PlayerBlack: playerOneId, PlayerWhite: playerTwoId, BoardSize: 19}}
	self.Games = append(self.Games, &g)
	return &g
}

func (self *GamesRepo) AllGames() []*Game {
	return self.Games
}

func (self *GamesRepo) FindGameById(id string) (g *Game, err error) {
	for _, g = range self.Games {
		if g.id == id {
			return
		}
	}
	err = errors.New(fmt.Sprintf("Game %s not found", id))
	return
}

// debug function to seed a game
func init() {
	g := GamesRepository.NewGame("ingar", "ingar")
	g.Game.AddMove(igo.Move{igo.MOVE_PLACE, igo.COLOR_BLACK, igo.Coordinates{3, 3}})
	g.Game.AddMove(igo.Move{igo.MOVE_PLACE, igo.COLOR_WHITE, igo.Coordinates{5, 2}})
	g.Game.AddMove(igo.Move{igo.MOVE_PLACE, igo.COLOR_BLACK, igo.Coordinates{4, 2}})
	g.Game.AddMove(igo.Move{igo.MOVE_PLACE, igo.COLOR_WHITE, igo.Coordinates{4, 4}})
	g.Game.AddMove(igo.Move{igo.MOVE_PLACE, igo.COLOR_BLACK, igo.Coordinates{3, 1}})
	g.Game.AddMove(igo.Move{igo.MOVE_PLACE, igo.COLOR_WHITE, igo.Coordinates{10, 10}})
	g.Game.AddMove(igo.Move{igo.MOVE_PLACE, igo.COLOR_BLACK, igo.Coordinates{2, 2}})
}
