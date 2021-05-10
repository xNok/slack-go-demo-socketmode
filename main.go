package main

import (
	"os"
	"xnok/slack-go-demo/controllers"
	"xnok/slack-go-demo/drivers"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/slack-go/slack/socketmode"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	err := godotenv.Load("./test_slack.env")
	if err != nil {
		log.Fatal().Msg("Error loading .env file")
	}

	// Instanciate deps
	client, err := drivers.ConnectToSlackViaSocketmode()
	if err != nil {
		log.Error().
			Str("error", err.Error()).
			Msg("Unable to connect to slack")

		os.Exit(1)
	}

	// Inject Deps in router
	socketmodeHandler := socketmode.NewsSocketmodeHandler(client)

	// This if for Separate articles and demos. You can run there separatly or all together

	// Build a Slack App Home in Golang Using Socket Mode
	controllers.NewAppHomeController(socketmodeHandler)
	// Properly Welcome Users in Slack with Golang using Socket Mode
	controllers.NewGreetingController(socketmodeHandler)
	// Build Slack Slash Command in Golang Using Socket Mode
	controllers.NewSlashCommandController(socketmodeHandler)

	socketmodeHandler.RunEventLoop()

}
