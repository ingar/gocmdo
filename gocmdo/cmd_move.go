package gocmdo

import (
	"fmt"
)

func init() {
	registerHandler("move", func(args []string) string {
		gameId := args[0]
		game := GamesRepository.FindGameById(gameId)
		if game != nil {
			game.AddMove(args[1])
			return fmt.Sprintf("%v", game)
		}
		return fmt.Sprintf("Couldn't find game with id: %s", gameId)
	})
}
