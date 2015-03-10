package main

import (
	"github.com/ingar/gocmdo/gocmdo"
)

func main() {
	go gocmdo.StartGobanServer()
	gocmdo.Run()
	/*
		g, _ := gocmdo.LoadGame("-JizYROIBEPIP6jI7iOR")
		g.Game.PlayerWhite = "arglebargle"
		gocmdo.SaveGame(g)
	*/
}
