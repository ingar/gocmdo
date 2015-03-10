package gocmdo

import (
	"github.com/ingar/barglebot"
	"fmt"
)

func cmdShow(message barglebot.Message) (resp string, err error) {
	var game *Game
	if game, err = LoadGame(message.Args()[0]); err != nil {
		return
	}

	resp = fmt.Sprintf("%v\n", game.Game)

	var url string
	if url, err = game.URL(); err != nil {
		return
	}

	resp += url
	return
}

func init() {
	registerHandler("show", CommandHandler(cmdShow))
}
