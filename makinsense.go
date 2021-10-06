package main

import "C"
import (
	"log"

	"github.com/makindotcc/go-internal-game-mod/game"
)

func logToSyslogs() {
	log.SetOutput(Syslogger{})
}

func init() {
	logToSyslogs()
	log.Println("Initializing...")

	game := game.New()
	log.Println("Game initialized!")

	game.Chat.Print("<font color='#aabbcc'>makinsense</font> loaded ;]")
}

func main() {
	logToSyslogs()
	log.Println("ekipa es sa")
}
