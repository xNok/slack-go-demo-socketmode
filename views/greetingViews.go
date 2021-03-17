package views

import (
	"bytes"
	"embed"
	"encoding/json"
	"html/template"
	"io/ioutil"

	"github.com/slack-go/slack"
)

//go:embed greetingViews/*
var greetingAssets embed.FS

func GreetingMessage(user string) []slack.Block {

	// read the block-kit definition as a go template
	t, err := template.ParseFS(greetingAssets, "greetingViews/greeting.json")
	if err != nil {
		panic(err)
	}
	var tpl bytes.Buffer

	// we need a stuct to hold template arguments
	type args struct {
		User string
	}

	err = t.Execute(&tpl, args{User: user})
	if err != nil {
		panic(err)
	}

	view := slack.Msg{}

	str, _ := ioutil.ReadAll(&tpl)
	json.Unmarshal(str, &view)

	return view.Blocks.BlockSet
}
