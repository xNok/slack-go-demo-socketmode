package controllers

import (
	"os"
	"testing"

	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

func TestGreetingController_postGreetingMessage(t *testing.T) {

	testServer, api := setup_slacktest()
	defer testServer.Stop()

	soccketClient := socketmode.New(
		api,
	)
	user := os.Getenv("TEST_USER")

	type args struct {
		evt *socketmode.Event
		clt *socketmode.Client
	}
	tests := []struct {
		name string
		c    *GreetingController
		args args
	}{
		{
			name: "Post greeting message when Member join channel",
			c:    &GreetingController{},
			args: args{
				evt: &socketmode.Event{
					Type: socketmode.EventTypeEventsAPI,
					Data: slackevents.EventsAPIEvent{
						Type: slackevents.CallbackEvent,
						InnerEvent: slackevents.EventsAPIInnerEvent{
							Type: string(slackevents.MemberJoinedChannel),
							Data: &slackevents.MemberJoinedChannelEvent{
								User: user,
							},
						},
					},
					Request: &socketmode.Request{
						EnvelopeID: "dummy",
					},
				},
				clt: soccketClient,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// When
			tt.c.postGreetingMessage(tt.args.evt, tt.args.clt)

			// Then -> recieve ephemeral message

		})
	}
}

func TestGreetingController_reactToMention(t *testing.T) {

	testServer, api := setup_slacktest()
	defer testServer.Stop()

	soccketClient := socketmode.New(
		api,
	)
	user := os.Getenv("TEST_USER")

	type args struct {
		evt *socketmode.Event
		clt *socketmode.Client
	}
	tests := []struct {
		name string
		c    *GreetingController
		args args
	}{
		{
			name: "Post a message in App Home when App is mentionned",
			c:    &GreetingController{},
			args: args{
				evt: &socketmode.Event{
					Type: socketmode.EventTypeEventsAPI,
					Data: slackevents.EventsAPIEvent{
						Type: slackevents.CallbackEvent,
						InnerEvent: slackevents.EventsAPIInnerEvent{
							Type: string(slackevents.AppMention),
							Data: &slackevents.AppMentionEvent{
								User: user,
							},
						},
					},
					Request: &socketmode.Request{
						EnvelopeID: "dummy",
					},
				},
				clt: soccketClient,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.reactToMention(tt.args.evt, tt.args.clt)
		})
	}
}
