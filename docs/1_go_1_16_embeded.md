# Manage Static Assets with `embed` (Golang 1.16)

## Overview

Golang 1.16 new package `embed` manages static assets, embedding them in the application binary and letting you use them easily. Any files from a package or package subdirectory can be "embedded" and retrieved as a variable of type `string` or `bytes[]`.

```
import _ "embed"

//go:embed hello.txt
var s strings

//go:embed hello.txt
var b []bytes
```

Besides, you can retrieve your embedded files with a variable of type `FS`. You can define which file needs to be "embedded" in your application using a [glob](https://man7.org/linux/man-pages/man7/glob.7.html) pathname.

```
import "embed"

//go:embed assets/*
var f embed.FS
```

[Official Documentation](https://golang.org/pkg/embed/)

## Use case: Using embed in your SlackBot

Given how easy it is to create and edit your messages using [Block-kit Builder](https://app.slack.com/block-kit-builder/T0B5XJYR2), I believe that the most convenient method to design and maintain your SlackBot is to save the design created with Block-kit as a `json` payload. The new `embed` package is the perfect feature for our use case. In terms of design pattern, those `json` payloads represent the *View* a classical [MVC application](https://en.wikipedia.org/wiki/Model%E2%80%93view%E2%80%93controller). We can send those messages as is or use some templating to add data to them.

In my [Slack demo application](https://github.com/xNok/slack-go-demo-socketmode), I have used this method in all my `Views` in combination with [go markup language](https://golang.org/pkg/text/template/). In this section, I will be demonstrating how to manage a greeting message designed with Block-kit. I will only focus on the `View` part of the application, ignoring the implementation of `Model` and `Controller`. Nevertheless, feel free to peak at them in my git repository; Also, I am planning a more generic article on creating SlackBots using Golang, covering those details.

### Create a Message with Block-kit

In this step no code required! Go to [Block Kit Builder](https://app.slack.com/block-kit-builder/T0B5XJYR2#%7B%22blocks%22:%5B%7B%22type%22:%22section%22,%22text%22:%7B%22type%22:%22mrkdwn%22,%22text%22:%22Hi%20David%20:wave:%22%7D%7D,%7B%22type%22:%22section%22,%22text%22:%7B%22type%22:%22mrkdwn%22,%22text%22:%22Great%20to%20see%20you%20here!%20App%20helps%20you%20to%20stay%20up-to-date%20with%20your%20meetings%20and%20events%20right%20here%20within%20Slack.%20These%20are%20just%20a%20few%20things%20which%20you%20will%20be%20able%20to%20do:%22%7D%7D,%7B%22type%22:%22section%22,%22text%22:%7B%22type%22:%22mrkdwn%22,%22text%22:%22%E2%80%A2%20Schedule%20meetings%20%5Cn%20%E2%80%A2%20Manage%20and%20update%20attendees%20%5Cn%20%E2%80%A2%20Get%20notified%20about%20changes%20of%20your%20meetings%22%7D%7D,%7B%22type%22:%22section%22,%22text%22:%7B%22type%22:%22mrkdwn%22,%22text%22:%22But%20before%20you%20can%20do%20all%20these%20amazing%20things,%20we%20need%20you%20to%20connect%20your%20calendar%20to%20App.%20Simply%20click%20the%20button%20below:%22%7D%7D,%7B%22type%22:%22actions%22,%22elements%22:%5B%7B%22type%22:%22button%22,%22text%22:%7B%22type%22:%22plain_text%22,%22text%22:%22Connect%20account%22,%22emoji%22:true%7D,%22value%22:%22click_me_123%22%7D%5D%7D%5D%7D). Customize your template, copy the `json` payload and, save it to a file. [greeting.json](../views/greetingViews/greeting.json) in my case.

Next edit this template to make it customizeable using go [markup language](https://golang.org/pkg/text/template/). For instance, I want to add the name of the user that recieve the message, then the text for the message will likke like this:

```
Hi {{ .User }} :wave:
```

After rendering the template I would expect (if my user is called David)

```
Hi David :wave:
```

### Render the Message

First, lets use `embed` and declare a variable `greetingAssets` that refers to our asset folder.

```go
import (
	"embed"
)

//go:embed greetingViews/*
var greetingAssets embed.FS
```

Second, let’s create a function that takes the `user` name as a string and returns a slack of `slack.Block`. Those slack.Block(s) represent the blocks we have created with Block-kit and saved into the file `greetingViews/greeting.json`

```go
func GreetingMessage(user string) []slack.Block {

	// [TODO]: parse the template `greetingViews/greeting.json`

	view := slack.Msg{}

	// [TODO]: unmarshal the template into slack.Msg{}

	return view.Blocks.BlockSet
}
```

Next, we want to render `greetingViews/greeting.json` using `greetingAssets` and the user name provided as input for our function. To do so I created a small utility function because it might be reused accross our application. This function takes as arguments a variable of type `fs.FS` such as `greetingAssets`, the path of the file to use as a template and, a variable of type `interface{}` that represent any struct that contains data to interpolate in the template.

[utils.go](../views/utils.go)
```go
func renderTemplate(fs fs.FS, file string, args interface{}) bytes.Buffer {

	var tpl bytes.Buffer

	// read the block-kit definition as a go template
	t, err := template.ParseFS(fs, file)
	if err != nil {
		panic(err)
	}

	// render the template using provided datas
	err = t.Execute(&tpl, args)
	if err != nil {
		panic(err)
	}

	return tpl
}

```

Finally, we put all the pieces together:
* read the block-kit definition as a go template and interpolate data
* unmarshal the template into slack.Msg{}

[greetingViews.go](../views/greetingViews.go)
```go
func GreetingMessage(user string) []slack.Block {

	// we need a stuct to hold template arguments
	type args struct {
		User string
	}

	tpl := renderTemplate(greetingAssets, "greetingViews/greeting.json", args{User: user})

	// we convert the view into a message struct
	view := slack.Msg{}

	str, _ := ioutil.ReadAll(&tpl)
	json.Unmarshal(str, &view)

	// We only return the block because of the way the PostEphemeral function works
	// we are going to use slack.MsgOptionBlocks in the controller
	return view.Blocks.BlockSet
}

```

## 3 Articles tackeling `embed` in a different context

* [Working with Embed in Go 1.16 Versions](https://lakefs.io/working-with-embed-in-go/)
* [Golang embed static assets in binary (with React build example)](https://www.akmittal.dev/posts/go-embed-files/)
* [How to Use //go:embed](https://blog.carlmjohnson.net/post/2021/how-to-use-go-embed/)