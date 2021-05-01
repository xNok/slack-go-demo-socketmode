package controllers

import (
	"github.com/slack-go/slack/socketmode"
)

// We create a sctucture to let us use dependency injection
type SlashCommandController struct {
	EventHandler *socketmode.SocketmodeHandler
}

func NewSlashCommandController(eventhandler *socketmode.SocketmodeHandler) SlashCommandController {

	c := SlashCommandController{
		EventHandler: eventhandler,
	}

	return c

}
