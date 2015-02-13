package bot

import (
	"fmt"
	"github.com/ingar/barglebot"
	"github.com/ingar/barglebot/transport/slack"
	"os"
	"strings"
)

var handlers = map[string]func([]string) string{}

func debug(s string) {
	fmt.Println("[GOCMDO]", s)
}

func RegisterHandler(command string, handler func([]string) string) {
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
		handleIncoming(<-incomingCommands)
	}
}
