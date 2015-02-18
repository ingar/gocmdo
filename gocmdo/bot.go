package gocmdo

import (
	"fmt"
	"github.com/ingar/barglebot"
	"github.com/ingar/barglebot/transport/slack"
	"os"
	"strings"
)

// A CommandHandler handles incoming bot commands
type CommandHandler func([]string) string

var handlers = map[string]CommandHandler{}

func debug(s string) {
	fmt.Println("[GOCMDO]", s)
}

func registerHandler(command string, handler func([]string) string) {
	handlers[command] = handler
}

func handleIncoming(message barglebot.Message) {
	debug("Handling incoming message")

	tokens := strings.Split(message.Text(), " ")

	if handler, ok := handlers[strings.ToLower(tokens[0])]; ok {
		message.Respond(handler(tokens[1:]))
	} else {
		debug(fmt.Sprintf("Unknown command: %s", message.DebugDump()))
	}
}

func Run() {
	incomingCommands := make(chan barglebot.Message)
	go slack.Connect(os.Getenv("SLACK_BOT_API_KEY"), incomingCommands)
	for {
		go handleIncoming(<-incomingCommands)
	}
}
