package gocmdo

import (
	"errors"
	"strconv"
)

// Game
type Game struct {
	id          string
	PlayerOneId string
	PlayerTwoId string
	Moves       []string
}

func (self *Game) AddMove(move string) error {
	if self.validMove(move) {
		self.Moves = append(self.Moves, move)
		return nil
	}
	return errors.New("Invalid move")
}

func (self *Game) validMove(move string) bool {
	return false
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
