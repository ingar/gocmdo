package gocmdo

import (
	"fmt"
)

func newgame(args []string) string {
	game := GamesRepository.NewGame(args[0], args[1])
	return fmt.Sprintf("%v", game)
}

func init() {
	registerHandler("newgame", newgame)
}
