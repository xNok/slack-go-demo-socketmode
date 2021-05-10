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

	// parse the command line
	count := 3

	// create the view using block-kit
	blocks := views.LaunchRocket(count)

	client := clt.GetApiClient()

	// Post reponse message (3) in User's App Home
	_, ts, err := client.PostMessage(
		command.ChannelID,
		slack.MsgOptionBlocks(blocks...),
		slack.MsgOptionResponseURL(command.ResponseURL, slack.ResponseTypeInChannel),
	)

	// Handle errors
	if err != nil {
		log.Printf("ERROR while sending message for /rocket: %v", err)
	}

	for i := count; i >= 0; i-- {
		// create the view using block-kit
		blocks = views.LaunchRocket(i)

		_, _, _, err = clt.GetApiClient().UpdateMessage(
			command.ChannelID,
			ts,
			slack.MsgOptionBlocks(blocks...),
			slack.MsgOptionResponseURL(command.ResponseURL, slack.ResponseTypeInChannel),
			slack.MsgOptionReplaceOriginal(command.ResponseURL),
		)

		// Handle errors
		if err != nil {
			log.Printf("ERROR while sending message for /rocket: %v", err)
		}
	}

}
