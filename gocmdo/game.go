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
	c.column, _ = strconv.Atoi(tokens[0])
	c.row, _ = strconv.Atoi(tokens[1])
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

func (self *GamesRepo) FindGameById(id string) (g *Game, err error) {
	for _, g = range self.Games {
		if g.id == id {
			return
		}
	}
	err = errors.New("Game not found")
	return
}

// debug function to seed a game
func init() {
	g := Game{id: "1", PlayerOneId: "ingar", PlayerTwoId: "ingar"}
	g.AddMove("3,3")
	g.AddMove("5,2")

	GamesRepository.Games = append(GamesRepository.Games, &g)
}
