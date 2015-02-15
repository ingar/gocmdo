package gocmdo

import (
	"fmt"
)

func allgames(args []string) string {
	games := GamesRepository.AllGames()
	out := "["
	for _, g := range games {
		out += fmt.Sprintf("%v", *g)
	}
	out += "]"
	return out
}

func init() {
	registerHandler("allgames", allgames)
}
