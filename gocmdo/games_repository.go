package gocmdo

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ingar/igo"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type GamesDump map[string]Game

// GamesRepo
type GamesRepo struct {
	Games GamesDump
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

func (self *GamesRepo) Persist(g *Game) (err error) {
	url := fmt.Sprintf("https://gocmdo-dev.firebaseio.com/games.json?auth=%s", os.Getenv("FIREBASE_SECRET"))

	req, err := http.NewRequest("POST", url, strings.NewReader(g.Json()))

	if err != nil {
		return
	}

	fmt.Println(req)
	fmt.Println(g.Json())

	client := http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return
	}

	if _, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}

	return
}

func (self *GamesRepo) LoadAll() (err error) {
	url := fmt.Sprintf("https://gocmdo-dev.firebaseio.com/games.json?auth=%s", os.Getenv("FIREBASE_SECRET"))

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return
	}

	client := http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return
	}

	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}

	fmt.Println("LOADALL")
	fmt.Println(string(body))

	var allGames GamesDump
	if err = json.Unmarshal(body, &allGames); err != nil {
		fmt.Println("Error:", err)
	} else {
		for _, g := range allGames {
			fmt.Println(g)
		}
	}

	return
}

func init() {
	GamesRepository.LoadAll()
}
