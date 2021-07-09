package controllers

import (
	"log"
	"xnok/slack-go-demo/views"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

// We create a sctucture to let us use dependency injection
type ShortcutController struct {
	EventHandler *socketmode.SocketmodeHandler
}

func NewShortcutController(eventhandler *socketmode.SocketmodeHandler) ShortcutController {
	c := ShortcutController{
		EventHandler: eventhandler,
	}

	// Global shortcut
	c.EventHandler.HandleInteraction(
		slack.InteractionTypeShortcut,
		c.globalshortcut,
	)

	return c

}

func (c *ShortcutController) globalshortcut(evt *socketmode.Event, clt *socketmode.Client) {
	// we need to cast our socketmode.Event into slackevents.AppHomeOpenedEvent
	shortcut := evt.Data.(slack.InteractionCallback)

	// Make sure to respond to the server to avoid an error
	clt.Ack(*evt.Request)

	// create the view using block-kit
	view := views.CreateStickieNoteModal()

	// Open Modal (13)
	_, err := clt.GetApiClient().OpenView(shortcut.TriggerID, view)

	//Handle errors
	if err != nil {
		log.Printf("ERROR openCreateStickieNoteModal: %v", err)
	}

}
