package game

// #include "dyld.h"
import "C"
import "github.com/makindotcc/go-internal-game-mod/game/chat"

type Game struct {
	Chat chat.Chat
}

func New() Game {
	return Game{
		Chat: chat.New(uintptr(C.GetMainModuleAddress())),
	}
}
