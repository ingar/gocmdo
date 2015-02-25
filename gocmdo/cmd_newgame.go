package gocmdo

import (
	"errors"
	"fmt"
	"github.com/ingar/barglebot"
)

func validUser(name string) (err error) {
	for _, user := range Users {
		if user.Name == name {
			return
		}
	}
	err = errors.New(fmt.Sprintf("'%s' is an invalid user", name))
	return
}

func cmdNewGame(message barglebot.Message) (resp string, err error) {
	args := message.Args()

	for _, user := range args[:2] {
		if err = validUser(user); err != nil {
			return
		}
	}

	game := GamesRepository.NewGame(args[0], args[1])
	resp = fmt.Sprintf("Game created: %v", *game)
	return
}

func init() {
	registerHandler("newgame", CommandHandler(cmdNewGame))
}
