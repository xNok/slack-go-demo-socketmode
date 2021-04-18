# Building Slack Bots in Golang

This project demonstrates how to build a Slackbot in Golang; it uses the [slack-go](https://github.com/slack-go/slack) library and communicates with slack using the [socket mode](https://api.slack.com/apis/connections/socket).

Working on this project inspired me to write a couple of articles that you can read here or on [Medium](https://medium.com/@couedeloalexandre)

* Article 1 : [Manage Static Assets in Golang](./docs/1_go_1_16_embeded.md) - [![Medium](https://img.shields.io/badge/Medium-12100E?style=for-the-badge&logo=medium&logoColor=white)](https://couedeloalexandre.medium.com/manage-static-assets-with-embed-golang-1-16-75c89c3eea39)
* Article 2 : [Handler and Middleware design pattern in Golang](./docs/2_middleware_design_pattern.md) - [![Medium](https://img.shields.io/badge/Medium-12100E?style=for-the-badge&logo=medium&logoColor=white)](https://medium.com/codex/handler-and-middleware-design-pattern-in-golang-de23ec452fce)
* Article 3 : [Diagrams as code 3 must have tools](./docs/3_diagrame_as_code.md) - [![Medium](https://img.shields.io/badge/Medium-12100E?style=for-the-badge&logo=medium&logoColor=white)](https://medium.com/geekculture/3-diagram-as-code-tools-that-combined-cover-all-your-needs-8f40f57d5cd8)
* Article 4 : [Building a home for your app 🏡, Revisited in Go](./docs/building_a_home.md) - [![Medium](https://img.shields.io/badge/Medium-12100E?style=for-the-badge&logo=medium&logoColor=white)](https://betterprogramming.pub/build-a-slack-app-home-in-golang-using-socket-mode-aff7b855bb31)
* WIP: Article 5: [You need a Dojo ⛪]()

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

## Showcases

### Greetings (AppMentionEvent)
![](./docs/assets/greeting.gif)

![](./out/controllers/greetingController/greetingController.png)

* [Building a home for your app 🏡, Revisited in Go](./docs/building_a_home.md) - [![Medium](https://img.shields.io/badge/Medium-12100E?style=for-the-badge&logo=medium&logoColor=white)](https://betterprogramming.pub/build-a-slack-app-home-in-golang-using-socket-mode-aff7b855bb31)

### App Home

![](./docs/assets/apphome_completed.gif)

![](./out/controllers/appHomeController/appHomeController.png)

* [Building a home for your app 🏡, Revisited in Go](./docs/building_a_home.md) - [![Medium](https://img.shields.io/badge/Medium-12100E?style=for-the-badge&logo=medium&logoColor=white)](https://betterprogramming.pub/build-a-slack-app-home-in-golang-using-socket-mode-aff7b855bb31)