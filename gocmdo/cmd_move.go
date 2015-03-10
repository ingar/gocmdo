package gocmdo

import (
	"errors"
	"fmt"
	"github.com/ingar/barglebot"
	"github.com/ingar/igo"
	"log"
)

const CMD_MOVE string = "move"

func cmdMoveSyntaxErr() error {
	return errors.New(fmt.Sprintf("Error parsing move command.  Syntax: %s <gameId> <coordinates>, eg: MOVE 2 Q16", CMD_MOVE))
}

func cmdMove(message barglebot.Message) (resp string, err error) {
	args := message.Args()
	if len(args) != 2 {
		err = cmdMoveSyntaxErr()
		return
	}

	gameId, a1coords := args[0], args[1]

	var game *Game
	if game, err = LoadGame(gameId); err != nil {
		return
	}

	user := message.Sender()
	if user != game.Game.PlayerWhite && user != game.Game.PlayerBlack {
		err = errors.New(fmt.Sprintf("User '%s' is not playing in game %s vs %s", user, game.Game.PlayerBlack, game.Game.PlayerWhite))
		return
	}

	var coords igo.Coordinates
	if coords, err = igo.A1toXY(a1coords); err != nil {
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

	move := igo.Move{igo.MOVE_PLACE, color, coords}
	log.Println("Adding move", move)

	if err = game.Game.AddMove(move); err != nil {
		return
	}

	if err = SaveGame(game); err != nil {
		return
	}

	resp = game.String()
	return
}

func init() {
	registerHandler("move", CommandHandler(cmdMove))
}
