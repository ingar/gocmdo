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
	return fmt.Sprintf("%v", NewBoardGrid(game).Ascii())
}

func init() {
	registerHandler("move", cmdMove)
}
