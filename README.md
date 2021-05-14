# Building Slack Bots in Golang

<p>
<a href="https://github.com/xNok/slack-go-demo-socketmode/actions/workflows/go.yml"><img src="https://github.com/xNok/slack-go-demo-socketmode/actions/workflows/go.yml/badge.svg" height=25></a>
<a href="https://couedeloalexandre.medium.com/"><img src="https://img.shields.io/badge/medium-%2312100E.svg?&style=for-the-badge&logo=medium&logoColor=white" height=25></a> 
<a href="https://dev.to/xnok"><img src="https://img.shields.io/badge/DEV.TO-%230A0A0A.svg?&style=for-the-badge&logo=dev-dot-to&logoColor=white" height=25></a>
</p>


This project demonstrates how to build a Slackbot in Golang; it uses the [slack-go](https://github.com/slack-go/slack) library and communicates with slack using the [socket mode](https://api.slack.com/apis/connections/socket).

Working on this project inspired me to write a couple of articles that you can read here or on [Medium](https://medium.com/@couedeloalexandre) (* I included friend links so you won't hit Medium Paywall)

### Slack Bot Tutorials

* Tutorial 0 : [Every Slack Application starts with a few simple steps](0_1_your_first_steps.md) - [![Medium](https://img.shields.io/badge/Medium-12100E?style=for-the-badge&logo=medium&logoColor=white)](https://medium.com/geekculture/build-slack-application-in-golang-using-socket-mode-setting-up-57f454e63c0a?sk=03196a8c492a76cd7ef9a73628fe4524)

* Tutorial 1 : [Build Your Slack App Home in Golang Using Socket Mode](./docs/building_a_home.md) - [![Medium](https://img.shields.io/badge/Medium-12100E?style=for-the-badge&logo=medium&logoColor=white)](https://betterprogramming.pub/build-a-slack-app-home-in-golang-using-socket-mode-aff7b855bb31?sk=b8b1f7f3c03972793b26bec02dc3d2cc)

* Tutorial 2 : [Properly Welcome Users in Slack with Golang using Socket Mode](./docs/building_a_home.md) - [![Medium](https://img.shields.io/badge/Medium-12100E?style=for-the-badge&logo=medium&logoColor=white)](https://levelup.gitconnected.com/properly-welcome-users-in-slack-with-golang-using-socket-mode-9a206d30a34a?sk=24fb8c44c1128cc3ef366d51d4a48812)

* WIP: Tutorial 3: [Slash commands in Slack]()

### Other topics

* Article 1 : [Manage Static Assets in Golang](./docs/0_1_go_1_16_embeded.md) - [![Medium](https://img.shields.io/badge/Medium-12100E?style=for-the-badge&logo=medium&logoColor=white)](https://couedeloalexandre.medium.com/manage-static-assets-with-embed-golang-1-16-75c89c3eea39?sk=d903d7b0532aff64243ef419346f804b)
* Article 2 : [Handler and Middleware design pattern in Golang](./docs/2_middleware_design_pattern.md) - [![Medium](https://img.shields.io/badge/Medium-12100E?style=for-the-badge&logo=medium&logoColor=white)](https://medium.com/codex/handler-and-middleware-design-pattern-in-golang-de23ec452fce?sk=0eed25a60858ad985ad22274505fb992)
* Article 3 : [Diagrams as code 3 must have tools](./docs/3_diagrame_as_code.md) - [![Medium](https://img.shields.io/badge/Medium-12100E?style=for-the-badge&logo=medium&logoColor=white)](https://medium.com/geekculture/3-diagram-as-code-tools-that-combined-cover-all-your-needs-8f40f57d5cd8?sk=52fe49e20d7b3a37123d07b29b102696)
* Article 4 : [Golang’s untyped constants might make your work easier. But there’s a catch](https://betterprogramming.pub/stop-mixing-constants-with-the-type-string-in-golang-d3589d8ae84d?sk=455a5bb28fc70eae0c3b40013c0526dd) - [![Medium](https://img.shields.io/badge/Medium-12100E?style=for-the-badge&logo=medium&logoColor=white)](https://betterprogramming.pub/stop-mixing-constants-with-the-type-string-in-golang-d3589d8ae84d?sk=455a5bb28fc70eae0c3b40013c0526dd)

* WIP: Article 5 : [You need a Dojo ⛪]()


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

* [Master one-on-one interactions](./docs/4_welcome_you_team_members) - [![Medium](https://img.shields.io/badge/Medium-12100E?style=for-the-badge&logo=medium&logoColor=white)](https://levelup.gitconnected.com/properly-welcome-users-in-slack-with-golang-using-socket-mode-9a206d30a34a?sk=24fb8c44c1128cc3ef366d51d4a48812)

### App Home

![](./docs/assets/apphome_completed.gif)

![](./out/controllers/appHomeController/appHomeController.png)

* [Building a home for your app 🏡, Revisited in Go](./docs/building_a_home.md) - [![Medium](https://img.shields.io/badge/Medium-12100E?style=for-the-badge&logo=medium&logoColor=white)](https://betterprogramming.pub/build-a-slack-app-home-in-golang-using-socket-mode-aff7b855bb31?sk=b8b1f7f3c03972793b26bec02dc3d2cc)
## References
* [Building a home for your app 🏡](https://api.slack.com/tutorials/app-home-with-modal)