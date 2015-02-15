package gocmdo

import (
	"strconv"
)

// Game
type Game struct {
	id          string
	PlayerOneId string
	PlayerTwoId string
	Moves       []string
}

func (self *Game) AddMove(move string) {
	self.Moves = append(self.Moves, move)
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

func (self *GamesRepo) AddMove(id string, move string) *Game {
	g := self.FindGameById(id)
	if g != nil {
		g.AddMove(move)
	}
	return g
}

func (self *GamesRepo) FindGameById(id string) *Game {
	for _, g := range self.Games {
		if g.id == id {
			return g
		}
	}
	return nil
}
