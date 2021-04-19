package controllers

import (
	"os"
	"testing"
	"xnok/slack-go-demo/drivers"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/slacktest"
	"github.com/slack-go/slack/socketmode"
)

func setup_slacktest() (*slacktest.Server, *slack.Client) {
	// Set up the test server.
	testServer := slacktest.NewTestServer()
	go testServer.Start()

	// Setup and start the RTM.
	api := slack.New("ABCD", slack.OptionAPIURL(testServer.GetAPIURL()))

	return testServer, api
}

func TestAppHomeController_publishHomeTabView(t *testing.T) {

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

// We can use that test to reset the App home for demos
func TestAppHomeController_publishHomeTabView_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip()
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
