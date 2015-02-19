package gocmdo

import (
	"errors"
	"strconv"
	"strings"
)

// Game
type Coordinates struct {
	row    int
	column int
}

type Game struct {
	id          string
	PlayerOneId string
	PlayerTwoId string
	Moves       []Coordinates
}

func (self *Game) AddMove(move string) error {
	if c, err := self.validMove(move); err == nil {
		self.Moves = append(self.Moves, c)
		return nil
	}
	return errors.New("Invalid move")
}

func (self *Game) validMove(move string) (c Coordinates, err error) {
	err = nil
	tokens := strings.Split(move, ",")
	c.row, _ = strconv.Atoi(tokens[0])
	c.column, _ = strconv.Atoi(tokens[1])
	return
}

type BoardGrid [19][19]string

func (self *Game) Board() (b BoardGrid) {
	b = BoardGrid{}

	for idx, move := range self.Moves {
		s := "W"
		if idx%2 == 0 {
			s = "B"
		}
		b[move.column][move.row] = s
	}
	return
}

// GamesRepo
type GamesRepo struct {
	Games []*Game
}

var GamesRepository GamesRepo

func (self *GamesRepo) NewGame(playerOneId string, playerTwoId string) *Game {
	g := Game{id: strconv.Itoa(len(self.Games) + 1), PlayerOneId: playerOneId, PlayerTwoId: playerTwoId}
	self.Games = append(self.Games, &g)
	return &g
}

func (self *GamesRepo) AllGames() []*Game {
	return self.Games
}

func (self *GamesRepo) AddMove(id string, move string) (g *Game, err error) {
	g, err = self.FindGameById(id)
	if err == nil {
		err = g.AddMove(move)
	}
	return
}

func (self *GamesRepo) FindGameById(id string) (*Game, error) {
	for _, g := range self.Games {
		if g.id == id {
			return g, nil
		}
	}
	return nil, errors.New("Game not found")
}
