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

	// Register callback for the command /rocket
	c.EventHandler.HandleSlashCommand(
		"/rocket",
		c.launchRocket,
	)

	return c

}

func (c SlashCommandController) launchRocket(evt *socketmode.Event, clt *socketmode.Client) {

}
