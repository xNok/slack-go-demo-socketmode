package controllers

import (
	"os"
	"testing"
	"xnok/slack-go-demo/drivers"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

func TestAppHomeController_publishHomeTabView_Integration(t *testing.T) {
	// This is an Integration test
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	// Use an SLACK_BOT_TOKEN, SLACK_APP_TOKEN, TEST_USER to do our test
	godotenv.Load("../test_slack.env")

	soccketClient, _ := drivers.ConnectToSlackViaSocketmode()
	user := os.Getenv("TEST_USER")

	type args struct {
		evt *socketmode.Event
		clt *socketmode.Client
	}
	tests := []struct {
		name string
		c    AppHomeController
		args args
	}{
		{
			name: "Publish Home Tab for test User",
			c:    AppHomeController{},
			args: args{
				evt: &socketmode.Event{
					Type: socketmode.EventTypeEventsAPI,
					Data: slackevents.EventsAPIEvent{
						Type: slackevents.CallbackEvent,
						InnerEvent: slackevents.EventsAPIInnerEvent{
							Type: string(slackevents.AppHomeOpened),
							Data: slackevents.AppHomeOpenedEvent{
								User: user,
							},
						},
					},
				},
				clt: soccketClient,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.publishHomeTabView(tt.args.evt, tt.args.clt)
		})
	}
}
