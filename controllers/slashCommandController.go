package controllers

import (
	"log"
	"xnok/slack-go-demo/views"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

// We create a sctucture to let us use dependency injection
type SlashCommandController struct {
	EventHandler *socketmode.SocketmodeHandler
}

func NewSlashCommandController(eventhandler *socketmode.SocketmodeHandler) SlashCommandController {
	// we need to cast our socketmode.Event into a SlashCommand
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
	// we need to cast our socketmode.Event into a Slash Command
	command, ok := evt.Data.(slack.SlashCommand)

	if ok != true {
		log.Printf("ERROR converting event to Slash Command: %v", ok)
	}

	// Make sure to respond to the server to avoid an error
	clt.Ack(*evt.Request)

	// create the view using block-kit
	blocks := views.GreetingMessage("bob")

	// Post greeting message (3) in User's App Home
	// Pass a user's ID as the value of channel to post to that user's App Home
	// We get the Api client from `clt`
	_, _, err := clt.GetApiClient().PostMessage(
		command.ChannelID,
		slack.MsgOptionBlocks(blocks...),
	)

	//Handle errors
	if err != nil {
		log.Printf("ERROR while sending message for /rocket: %v", err)
	}
}
