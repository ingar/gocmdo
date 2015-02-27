package gocmdo

import (
	//"fmt"
	"github.com/ingar/barglebot"
)

func cmdGame(message barglebot.Message) (resp string, err error) {
	/*
		games := GamesRepository.AllGames()
		for _, g := range games {
			resp += fmt.Sprintf("Game %s: %v\n", g.id, g.Game)
		}
	*/
	return
}

func init() {
	registerHandler("games", CommandHandler(cmdGame))
}
