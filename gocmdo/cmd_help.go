package gocmdo

func help(args []string) string {
	return "Go HELP yourself at http://www.google.com"
}

func init() {
	registerHandler("help", help)
}
