package chat

// #include "chat.h"
import "C"
import (
	"unsafe"

	"github.com/makindotcc/go-internal-game-mod/game/offset"
)

type Channel int

const (
	ChannelTeam = 0
	ChannelAll  = 1
)

type Chat struct {
	gameChatClient uintptr
	gamePrintChat  uintptr
}

func New(mainModule uintptr) Chat {
	return Chat{
		gameChatClient: mainModule + offset.ChatClient,
		gamePrintChat:  mainModule + offset.ChatPrint,
	}
}

func (c Chat) Print(message string) {
	cMessage := C.CString(message)
	C.printChat(C.longlong(c.gameChatClient), C.longlong(c.gamePrintChat), cMessage)
	C.free(unsafe.Pointer(cMessage))
}
