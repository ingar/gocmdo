package gocmdo

import (
	"fmt"
)

func validUser(name string) (valid bool) {
	for _, user := range Users {
		if user.Name == name {
			valid = true
			break
		}
	}
	return
}

func newgame(args []string) (response string) {
	if validUser(args[0]) && validUser(args[1]) {
		game := GamesRepository.NewGame(args[0], args[1])
		response = fmt.Sprintf("Game created: %v", *game)
		return
	}
	response = fmt.Sprintf("Invalid user id")
	return
}

func init() {
	registerHandler("newgame", newgame)
}
