package gocmdo

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ingar/gocmdo/gocmdo/firebase"
	//"github.com/ingar/igo"
	//"io"
	//"io/ioutil"
	//"net/http"
	//"os"
	//"strconv"
	"strings"
)

var gameCache = map[string]*Game{}

// Loads a game from the store
func LoadGame(id string) (g *Game, err error) {
	g, ok := gameCache[id]

	if ok {
		return
	}

	buf, err := firebase.Get(fmt.Sprintf("games/%s", id))
	if err != nil {
		return
	}

	g = &Game{}
	if err = json.Unmarshal(buf, g); err != nil {
		return
	}

	// debug
	g.Id = id

	gameCache[id] = g

	return
}

// Persists a game to the store
func SaveGame(g *Game) (err error) {
	if g.Id == "" {
		err = saveNewGame(g)
	} else {
		err = updateGame(g)
	}
	return
}

func updateGame(g *Game) (err error) {
	if g.Id == "" {
		err = errors.New("Tried to call updateGame with a game that has no Id")
		return
	}

	_, err = firebase.Patch(fmt.Sprintf("games/%s", g.Id), strings.NewReader(g.Json()))
	return
}

func addToPlayerLists(g *Game) (err error) {
	for _, player := range []string{g.Game.PlayerWhite, g.Game.PlayerBlack} {
		_, err = firebase.Put(fmt.Sprintf("players/%s/games/%s", player, g.Id), strings.NewReader("true"))
		if err != nil {
			return
		}
	}
	return
}

func saveNewGame(g *Game) (err error) {
	if g.Id != "" {
		err = errors.New("Tried to call saveNewGame with a game that has an ID")
		return
	}

	// Push an empty object to obtain an Id
	buf, err := firebase.Post("games", strings.NewReader("{}"))
	if err != nil {
		return
	}

	var postResponse map[string]string
	if err = json.Unmarshal(buf, &postResponse); err != nil {
		return
	}

	g.Id = postResponse["name"]
	if err = updateGame(g); err != nil {
		return
	}

	err = addToPlayerLists(g)
	return
}

/*
func LiveGamesForUser(userId string) (games []*Game, err error) {

	games, err
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
*/
