# Create Slack App Home in golang using socketmode

This tutorial is inspired from the [slack documentation(https://api.slack.com/tutorials/app-home-with-modal) and feature its implementation in go using socketmode and [slack-go library](https://github.com/slack-go/slack)

## Configure your application

Refer yourself to the documentation [Setting up your app](https://api.slack.com/tutorials/app-home-with-modal#building-a-home-for-your-app---learn-how-to-create-the-app-home-view-and-use-the-modals__setting-up-your-app)


## Create the project

create a new go project and import `slack-go`

```
go mod init
go get -u github.com/slack-go/slack
```

In this tutorial I use my own fork of slack-go, because I am demonstrating the use of features that have yet been merged[#PR904](https://github.com/slack-go/slack/pull/904). 

To use a fork we need to add a replace statement in `go.mod`:

```
replace github.com/slack-go/slack => github.com/xnok/slack
```

Then we need to force that change to be taken into consideration:

```
go mod tidy
```

Create the following project structure

```
+ controllers
`- appHomeController.go
+ drivers
`- slack.go
+ views
`+ appHomeViewsAssets
  `- AppHomeView.json
   - CreateStickieNoteModal.json
   - NoteBlock.json
`- apphomeViews.go
+ main.go
```

## Divers > slack.go

In that file, we create a utility function to initialise our slack client using envirement varibles `SLACK_APP_TOKEN` and `SLACK_BOT_TOKEN`. In additiion it would be a good idea to add some validation. Slack provides two types of tokens:

```
SLACK_APP_TOKEN=xapp-xxxxxxxxx
SLACK_BOT_TOKEN=xoxb-xxxxxxxxx
```

[Slack driver code](../drivers/slack.go)

## Controllers > appHomeController.go

### Handeling events

I create a [sequence diagrame](../controllers/appHomeController.puml), inspired by [Tomomi Imura's](https://api.slack.com/tutorials/app-home-with-modal#building-a-home-for-your-app---learn-how-to-create-the-app-home-view-and-use-the-modals__setting-up-your-app), to visualy represent what we are about to code. I believe that it is very handy to keep such diagrammes alongside my code. It should make it much easier to follow this tutorial, in addition I added references to the diagrame in my code as comments. 

First create a *struct* representing our Controller to handle dependencies injection. So far we only need `socketmode.SocketmodeHandler` to register slack event we want to listen to. But in a bigger application you might have other dependencies such as **repositories** to handle database requestes.

```
type AppHomeController struct {
	EventHandler *socketmode.SocketmodeHandler
}
```

Second, create an initialisation function for our Controller. This function is in charge of registering which event we want to recieve and which function should handle that event. If you refer to the [sequence diagrame](../controllers/appHomeController.puml) this contoller needs listening to 3 events:
* An Event API called `app_home_opened` (2)
* An `interatcion` with the create button (12)
* An `interaction` with the modal submit button (22).

To register an event with `EventHandler`, you need fist to Identify what type of event we need to handle. I also suggest at this point you should get familiar with slack terminology including *Event API*, *Interaction*, *Block Action*. Therfore:
* If then event comes from [Event API](https://api.slack.com/events): use `HandleEventsAPI` function, and don't forget to subscribe to that event as explained in [Setting up your app](https://api.slack.com/tutorials/app-home-with-modal#building-a-home-for-your-app---learn-how-to-create-the-app-home-view-and-use-the-modals__setting-up-your-app)

* If it is an [interaction](https://api.slack.com/interactivity/handling): use `HandleInteraction` function
* If it is an [interaction with a block action](https://api.slack.com/reference/interaction-payloads/block-actions): use `HandleInteractionBlockAction` function

Each of those handler function works the same way. You provide the `subtype` of event you wish to listen (use autocompletion to find the one you need) and a callback function that require to argumentes 1. a pointer to the event (*socketmode.Event) 2. a pointer to the socketmode Client (*socketmode.Client). For instance this is a valide function

```
func callback(evt *socketmode.Event, clt *socketmode.Client) {}
```

This one is even better because we benefit from then dependancies injected in AppHomeController:

```
func (c *AppHomeController) callback(evt *socketmode.Event, clt *socketmode.Client) {}
```

To conclude this section, here is my initialisation function:

```
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
```

At this point our controller is ready. We only need to implement each of our three envent handeling methods:
* publishHomeTabView
* openCreateStickieNoteModal
* createStickieNote

### Implementing publishHomeTabView

### Implementing openCreateStickieNoteModal

### Implementing createStickieNote

## Views > appHomeViews.go

For the view we are going to utilise the power of [slack block-kit](https://api.slack.com/block-kit) to its full capacity. Therfore to create a view in `json` by storing the `Payload` provided by block-kit into files. Then we can load them when needed. To achive that I decided to use the newest *golang 1.16* feature created help manage static assets, namely [embed](https://github.com/akmittal/go-embed). 

This way we manage our slack Application like a simple web MVC application, with the views stored as static assests. It make updating our App much easier using block-kit and copy pasting the result.

#### Publish the App Home View

[Block-kit](https://app.slack.com/block-kit-builder/T0B5XJYR2#%7B%22type%22:%22home%22,%22blocks%22:%5B%7B%22type%22:%22section%22,%22text%22:%7B%22type%22:%22mrkdwn%22,%22text%22:%22*Welcome!*%20%5CnThis%20is%20a%20home%20for%20Stickers%20app.%20You%20can%20add%20small%20notes%20here!%22%7D,%22accessory%22:%7B%22type%22:%22button%22,%22action_id%22:%22add_note%22,%22text%22:%7B%22type%22:%22plain_text%22,%22text%22:%22Add%20a%20Stickie%22%7D%7D%7D,%7B%22type%22:%22divider%22%7D%5D%7D)

#### Opening a modal dialog

[Block-kit](https://app.slack.com/block-kit-builder/T0B5XJYR2#%7B%22title%22:%7B%22type%22:%22plain_text%22,%22text%22:%22Create%20a%20stickie%20note%22,%22emoji%22:true%7D,%22submit%22:%7B%22type%22:%22plain_text%22,%22text%22:%22Create%22,%22emoji%22:true%7D,%22type%22:%22modal%22,%22blocks%22:%5B%7B%22type%22:%22input%22,%22block_id%22:%22note01%22,%22element%22:%7B%22type%22:%22plain_text_input%22,%22placeholder%22:%7B%22type%22:%22plain_text%22,%22text%22:%22Take%20a%20note...%20%22%7D,%22multiline%22:true%7D,%22label%22:%7B%22type%22:%22plain_text%22,%22text%22:%22Note%22%7D%7D,%7B%22type%22:%22input%22,%22element%22:%7B%22type%22:%22static_select%22,%22action_id%22:%22color%22,%22options%22:%5B%7B%22text%22:%7B%22type%22:%22plain_text%22,%22text%22:%22yellow%22%7D,%22value%22:%22yellow%22%7D,%7B%22text%22:%7B%22type%22:%22plain_text%22,%22text%22:%22blue%22%7D,%22value%22:%22blue%22%7D%5D%7D,%22label%22:%7B%22type%22:%22plain_text%22,%22text%22:%22Color%22%7D%7D%5D%7D)

#### Updating the App Home view


