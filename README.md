# Building Slack Bots in Golang

This project demonstrates how to build a Slackbot in Golang; it uses the [slack-go](https://github.com/slack-go/slack) library and communicates with slack using the [socket mode](https://api.slack.com/apis/connections/socket).

Working on this project inspired me to write a couple article that you can read here or on [Medium](https://medium.com/@couedeloalexandre)

* Article 1 : [Manage Static Assets in Golang](./docs/1_go_1_16_embeded.md) - [Medium Version](https://couedeloalexandre.medium.com/manage-static-assets-with-embed-golang-1-16-75c89c3eea39)
* Article 2 : [Handler and Middleware design pattern in Golang](./docs/2_middleware_design_pattern.md) - [Medium Version](https://medium.com/codex/handler-and-middleware-design-pattern-in-golang-de23ec452fce)
* WIP: Article 3 [Diagrams as code 3 must have tools](./docs/3_diagrame_as_code.md)
* WIP: Article 4: [Building a home for your app 🏡, Revisited in Go](./docs/building_a_home.md)

References:
* [Building a home for your app 🏡](https://api.slack.com/tutorials/app-home-with-modal)

## Test the project

Create a file `test_slack.env` with the following variables:

```
SLACK_BOT_TOKEN=xoxb-xxxxxxxxxxx
SLACK_APP_TOKEN=xapp-1-xxxxxxxxx
```

Run the application

```
go run main.go
```

## Show cases


### Greetings (AppMentionEvent)
![](./docs/assets/greeting.gif)

![](./out/controllers/greetingController/greetingController.png)