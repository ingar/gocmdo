package gocmdo

import (
	"github.com/ingar/barglebot"
	"strings"
	"strconv"
	"fmt"
	"errors"
	"github.com/ingar/igo"
)

const CMD_MOVE string = "move"

func cmdMoveSyntaxErr() error {
	return errors.New(fmt.Sprintf("Error parsing move command.  Syntax: %s <gameId> <x>,<y>", CMD_MOVE))
}

func cmdMove(message barglebot.Message) (resp string, err error) {
	args := message.Args()
	if len(args) != 2 {
		err = cmdMoveSyntaxErr()
		return
	}

	gameId, coordinates := args[0], args[1]

	var game *Game
	if game, err = GamesRepository.FindGameById(gameId); err != nil {
		return
	}

	user := message.Sender()
	if user != game.Game.PlayerWhite && user != game.Game.PlayerBlack {
		err = errors.New(fmt.Sprintf("User '%s' is not playing in game %s vs %s", user, game.Game.PlayerBlack, game.Game.PlayerWhite))
		return
	}

	tokens := strings.Split(coordinates, ",")
	if len(tokens) != 2 {
		err = cmdMoveSyntaxErr()
		return
	}

	var x int
	if x, err = strconv.Atoi(tokens[0]); err != nil {
		err = cmdMoveSyntaxErr()
		return
	}

	var y int
	if y, err = strconv.Atoi(tokens[1]); err != nil {
		err = cmdMoveSyntaxErr()
		return
	}

	color := igo.COLOR_BLACK
	if user == game.Game.PlayerWhite {
		color = igo.COLOR_WHITE
	}

	// Allow users to play themselves
	if game.Game.PlayerBlack == game.Game.PlayerWhite {
		if game.Game.IsWhitesTurn() {
			color = igo.COLOR_WHITE
		} else {
			color = igo.COLOR_BLACK
		}
	}

	move := igo.Move{igo.MOVE_PLACE, color, igo.Coordinates{x, y}}

	err = game.Game.AddMove(move)
	resp = "Move added"
	return
}

func init() {
	registerHandler("move", CommandHandler(cmdMove))
}
