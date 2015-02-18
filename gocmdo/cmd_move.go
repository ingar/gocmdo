package gocmdo

import (
	"fmt"
)

func cmdMove(args []string) string {
	gameId := args[0]
	move := args[1]

	game, err := GamesRepository.AddMove(gameId, move)

	if err != nil {
		return fmt.Sprintf("%v", err)
	}
	return fmt.Sprintf("%v", game)
}

func init() {
	registerHandler("move", cmdMove)
}
