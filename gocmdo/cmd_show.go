package gocmdo

import "github.com/ingar/barglebot"

func cmdShow(message barglebot.Message) (resp string, err error) {
	var game *Game
	if game, err = LoadGame(message.Args()[0]); err != nil {
		return
	}

	resp = game.String()
	return
}

func init() {
	registerHandler("show", CommandHandler(cmdShow))
}
