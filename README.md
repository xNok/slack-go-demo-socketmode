# Building Slack Bots using socketmode and go

* Article 1 [Manage Static Assets with `embed` (Golang 1.16)](./docs/1_go_1_16_embeded.md)
* Article 2 [Handler and Middleware design pattern in Golang](./docs/2_middleware_design_pattern.md)
* WIP: Article 3: [Building a home for your app 🏡, Revisited in Go](./docs/building_a_home.md)

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