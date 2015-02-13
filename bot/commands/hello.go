package commands

import (
	"fmt"
	"github.com/ingar/gocmdo/bot"
)

func init() {
	bot.RegisterHandler("hello", func(args []string) string {
		image := "http://www.yosemitepark.com/Images/Header-Plan-1.1_img1.jpg"
		return fmt.Sprintf("Hello yoruself! %s", image)
	})
}
