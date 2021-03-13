package controllers

import (
	"log"
	"time"
	"xnok/slack-go-demo/views"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

// We create a sctucture to let us use dependency injection
type AppHomeController struct {
	EventHandler *socketmode.SocketmodeHandler
}

func NewAppHomeController(eventhandler *socketmode.SocketmodeHandler) AppHomeController {
	c := AppHomeController{
		EventHandler: eventhandler,
	}

	// App Home (2)
	c.EventHandler.HandleEventsAPI(
		slackevents.AppHomeOpened,
		c.publishHomeTabView,
	)

	// Create Stickie note Triggered (12)
	c.EventHandler.HandleInteractionBlockAction(
		views.AddStockieNoteActionID,
		c.openCreateStickieNoteModal,
	)

	// Create Stickie note Submitted (22)
	c.EventHandler.HandleInteraction(
		slack.InteractionTypeViewSubmission,
		c.createStickieNote,
	)

	return c

}

func (c *AppHomeController) publishHomeTabView(evt *socketmode.Event, clt *socketmode.Client) {
	// we need to cast our socketmode.Event into slackevents.AppHomeOpenedEvent
	evt_api, _ := evt.Data.(slackevents.EventsAPIEvent)
	evt_app_home_opened, _ := evt_api.InnerEvent.Data.(slackevents.AppHomeOpenedEvent)

	// create the view using block-kit
	view := views.AppHomeTabView()

	// Publish the view (3)
	// We get the Api client from `clt` and post our view
	_, err := clt.GetApiClient().PublishView(evt_app_home_opened.User, view, "")

	//Handle errors
	if err != nil {
		log.Printf("ERROR publishHomeTabView: %v", err)
	}
}

func (c *AppHomeController) openCreateStickieNoteModal(evt *socketmode.Event, clt *socketmode.Client) {
	// we need to cast our socketmode.Event
	interaction := evt.Data.(slack.InteractionCallback)

	// Make sure to respond to the server to avoid an error
	clt.Ack(*evt.Request)

	// create the view using block-kit
	view := views.CreateStickieNoteModal()

	// Open Modal (13)
	_, err := clt.GetApiClient().OpenView(interaction.TriggerID, view)

	//Handle errors
	if err != nil {
		log.Printf("ERROR openCreateStickieNoteModal: %v", err)
	}

}

func (c *AppHomeController) createStickieNote(evt *socketmode.Event, clt *socketmode.Client) {
	// we need to cast our socketmode.Event into slack.InteractionCallback
	view_submission := evt.Data.(slack.InteractionCallback)

	// Make sure to respond to the server to avoid an error
	clt.Ack(*evt.Request)

	// Create the model
	note := views.StickieNote{
		Description: view_submission.View.State.Values[views.ModalDescriptionBlockID][views.ModalDescriptionActionID].Value,
		Color:       view_submission.View.State.Values[views.ModalColorBlockID][views.ModalColorActionID].SelectedOption.Value,
		Timestamp:   time.Unix(time.Now().Unix(), 0).String(),
	}

	// create the view using block-kit
	view := views.AppHomeCreateStickieNote(note)

	// Publish the view (23)
	// We get the Api client from `clt` and post our view
	_, err := clt.GetApiClient().PublishView(view_submission.User.ID, view, "")

	//Handle errors
	if err != nil {
		log.Printf("ERROR createStickieNote: %v", err)
	}
}
