package commands

import "github.com/ingar/gocmdo/bot"

func init() {
	bot.RegisterHandler("help", func(args []string) string {
		return "Go HELP yourself at http://www.google.com"
	})
}
