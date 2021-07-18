package controllers

import (
	"bytes"
	"encoding/json"
	"log"
	"reflect"
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

	c.EventHandler.Handle(socketmode.EventTypeErrorBadMessage, c.recoverAppHomeOpened)

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

// HomeEvent is a helper struct for recovery from err bad message
type HomeEvent struct {
	Envelope string          `json:"envelope_id"`
	Payload  json.RawMessage `json:"payload"`
}

func (c *AppHomeController) recoverAppHomeOpened(evt *socketmode.Event, clt *socketmode.Client) {

	log.Printf("Atempt to recover Bad Message %v", evt)

	/*
		We do some handling here to attempt to recover the event from a known issue with "AppHomeOpened"
		If we fail to recover the assumption is that the bad event was for a legitimate reason and not
		an internal issue with slack-go
	*/
	var e *socketmode.ErrorBadMessage
	var ok bool
	var err error

	if e, ok = evt.Data.(*socketmode.ErrorBadMessage); !ok {
		log.Printf("Bad Message Not Cast: %+v", evt)
		return
	}
	var rawBytes []byte
	if rawBytes, err = e.Message.MarshalJSON(); err != nil {
		log.Printf("Bad Message Not Marshalled. Err: %+v\n Event: %+v", err, evt)
		return
	}

	/*
		This line replaces `"state":{"values":[]}` with `"state":{"values":{}}`
		The latter is parsed correctly by slack-go while the former is what kicks off the "ErrorBadMessage"
		event in the first place. It is, unfortunately, the best way to fix this behavior until there is
		clarification from slack on intended return values and a fix is merged into slack-go
	*/
	rawMessage := bytes.Replace(rawBytes, []byte{34, 115, 116, 97, 116, 101, 34, 58, 123, 34, 118, 97, 108, 117, 101, 115, 34, 58, 91, 93, 125}, []byte{34, 115, 116, 97, 116, 101, 34, 58, 123, 34, 118, 97, 108, 117, 101, 115, 34, 58, 123, 125, 125}, 1)
	var hE HomeEvent
	if err := json.Unmarshal(rawMessage, &hE); err != nil {
		log.Printf("Raw Message Not Marshalled: %s", err)
		return
	}

	// Parse the raw json payload without verifying the token as it will fail otherwise
	var newEvent slackevents.EventsAPIEvent
	if newEvent, err = slackevents.ParseEvent(hE.Payload, slackevents.OptionNoVerifyToken()); err != nil {
		log.Printf("Bad Message Not Parsed. Err: %+v\n Inner JSON: %+v", err, rawMessage)
		return
	}

	// Plug all of our parts into an event
	fabEvent := socketmode.Event{
		Type: socketmode.EventTypeEventsAPI,
		Data: newEvent,
		Request: &socketmode.Request{ //  we need to attach the envelope ID for the request to Ack
			Type:       "events_api",
			EnvelopeID: hE.Envelope,
		},
	}

	c.EventHandler.Client.Events <- fabEvent
}

func (c *AppHomeController) publishHomeTabView(evt *socketmode.Event, clt *socketmode.Client) {
	// we need to cast our socketmode.Event into slackevents.AppHomeOpenedEvent
	evt_api, ok := evt.Data.(slackevents.EventsAPIEvent)

	if ok != true {
		log.Printf("ERROR converting event to slackevents.EventsAPIEvent")
	}

	evt_app_home_opened, ok := evt_api.InnerEvent.Data.(slackevents.AppHomeOpenedEvent)

	var user string

	if ok != true {
		log.Printf("ERROR converting inner event to slackevents.AppHomeOpenedEvent")
		//Patch the fact that we are not able to cast evt_api.InnerEvent.Data to AppHomeOpenedEvent
		user = reflect.ValueOf(evt_api.InnerEvent.Data).Elem().FieldByName("User").Interface().(string)
	} else {
		user = evt_app_home_opened.User
	}

	log.Printf("ERROR publishHomeTabView: %v", evt_app_home_opened)

	// create the view using block-kit
	view := views.AppHomeTabView()

	// Publish the view (3)
	// We get the Api client from `clt` and post our view
	_, err := clt.GetApiClient().PublishView(user, view, "")

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
