package gocmdo

import (
	"fmt"
	"github.com/ingar/barglebot"
	"github.com/ingar/barglebot/transport/slack"
	"os"
	"strings"
)

// A CommandHandler handles incoming bot commands
type CommandHandler func(barglebot.Message) (string, error)

var handlers = map[string]CommandHandler{}

var Users []slack.User

func debug(s string) {
	fmt.Println("[GOCMDO]", s)
}

func registerHandler(command string, handler CommandHandler) {
	handlers[command] = handler
}

func handleIncoming(message barglebot.Message) {
	command := strings.ToLower(message.Tokens()[0])
	if handler, ok := handlers[command]; ok {
		response, err := handler(message)
		if err != nil {
			message.Respond(err.Error())
		} else {
			message.Respond(response)
		}
	} else {
		debug(fmt.Sprintf("Unknown command: %s", message.DebugDump()))
	}
}

func Run() {
	incomingCommands := make(chan barglebot.Message)
	Users = slack.Connect(os.Getenv("SLACK_BOT_API_KEY"), incomingCommands)
	for {
		go handleIncoming(<-incomingCommands)
	}
}
