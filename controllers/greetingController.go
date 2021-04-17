package controllers

import (
	"log"
	"xnok/slack-go-demo/views"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

// We create a sctucture to let us use dependency injection
type GreetingController struct {
	EventHandler *socketmode.SocketmodeHandler
}

func NewGreetingController(eventhandler *socketmode.SocketmodeHandler) GreetingController {
	c := GreetingController{
		EventHandler: eventhandler,
	}

	// App Home (2)
	c.EventHandler.HandleEventsAPI(
		slackevents.AppMention,
		c.reactToMention,
	)

	// App Home (2)
	c.EventHandler.HandleEventsAPI(
		slackevents.MemberJoinedChannel,
		c.postGreetingMessage,
	)

	return c

}

func (c *GreetingController) postGreetingMessage(evt *socketmode.Event, clt *socketmode.Client) {
	// we need to cast our socketmode.Event into slackevents.AppHomeOpenedEvent
	evt_api, _ := evt.Data.(slackevents.EventsAPIEvent)
	evt_member_join, ok := evt_api.InnerEvent.Data.(*slackevents.MemberJoinedChannelEvent)

	clt.Ack(*evt.Request)

	if ok != true {
		log.Printf("ERROR converting event to slackevents.AppMentionEvent: %v", ok)
	}

	userInfo, err := clt.GetApiClient().GetUserInfo(evt_member_join.User)

	if err != nil {
		log.Printf("ERROR unable to retrive user info: %v", err)
	}

	// create the view using block-kit
	blocks := views.GreetingMessage(userInfo.Name)

	// Post greeting message (3)
	// We get the Api client from `clt`
	_, err = clt.GetApiClient().PostEphemeral(
		evt_member_join.Channel,
		evt_member_join.User,
		slack.MsgOptionBlocks(blocks...),
	)

	//Handle errors
	if err != nil {
		log.Printf("ERROR postGreetingMessage: %v", err)
	}
}

func (c *GreetingController) reactToMention(evt *socketmode.Event, clt *socketmode.Client) {
	// we need to cast our socketmode.Event into slackevents.AppMentionEvent
	evt_api, _ := evt.Data.(slackevents.EventsAPIEvent)
	evt_app_mention, ok := evt_api.InnerEvent.Data.(*slackevents.AppMentionEvent)

	clt.Ack(*evt.Request)

	if ok != true {
		log.Printf("ERROR converting event to slackevents.MemberJoinedChannelEvent: %v", ok)
	}

	userInfo, err := clt.GetApiClient().GetUserInfo(evt_app_mention.User)

	if err != nil {
		log.Printf("ERROR unable to retrive user info: %v", err)
	}

	// create the view using block-kit
	blocks := views.GreetingMessage(userInfo.Name)

	// Post greeting message (3) in User's App Home
	// Pass a user's ID as the value of channel to post to that user's App Home
	// We get the Api client from `clt`
	_, _, err = clt.GetApiClient().PostMessage(
		evt_app_mention.User,
		slack.MsgOptionBlocks(blocks...),
	)

	//Handle errors
	if err != nil {
		log.Printf("ERROR reactToMention: %v", err)
	}
}
