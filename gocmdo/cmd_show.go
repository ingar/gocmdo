package gocmdo

import ()

func cmdShow(args []string) string {
	g, err := GamesRepository.FindGameById(args[0])
	if err == nil {
		return NewBoardGrid(g).Ascii()
	}
	return err.Error()
}

func init() {
	registerHandler("show", cmdShow)
}
