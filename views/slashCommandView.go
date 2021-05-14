package views

import (
	"embed"
	"encoding/json"
	"io/ioutil"

	"github.com/slack-go/slack"
)

//go:embed slackCommandAssets/*
var slashCommandAssets embed.FS

func LaunchRocketAnnoncement(number int) []slack.Block {
	// we need a stuct to hold template arguments
	type args struct {
		Number int
	}

	tpl := renderTemplate(slashCommandAssets, "slackCommandAssets/annnoncement.json", args{Number: number})

	// we convert the view into a message struct
	view := slack.Msg{}

	str, _ := ioutil.ReadAll(&tpl)
	json.Unmarshal(str, &view)

	// We only return the block because of the way the PostEphemeral function works
	// we are going to use slack.MsgOptionBlocks in the controller
	return view.Blocks.BlockSet

}

func LaunchRocket(number int) []slack.Block {

	// we need a stuct to hold template arguments
	type args struct {
		Number int
	}

	tpl := renderTemplate(slashCommandAssets, "slackCommandAssets/rocket.json", args{Number: number})

	// we convert the view into a message struct
	view := slack.Msg{}

	str, _ := ioutil.ReadAll(&tpl)
	json.Unmarshal(str, &view)

	// We only return the block because of the way the PostEphemeral function works
	// we are going to use slack.MsgOptionBlocks in the controller
	return view.Blocks.BlockSet
}
