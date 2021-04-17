## Welcome Members in Slack with Golang using Slack Socket Mode

When creating a Slack Application to increase engagement, it is essential to start small meaningful one-on-one interaction. For instance, you can send a short tutorial on how to your App in the Slack [App Home](https://api.slack.com/start/overview#app_home). You can also introduce the purpose and rules of a channel whenever a user joins, with a message only he can see as not polluting other members.

I want this article to help you understand some core features of Slack. Before moving on into coding, let me showcase the two use cases and Slack's terminology. With this base setup, you will be able to create much more exciting interactions with your users. I consider those as the basis of any slack application.

This tutorial guides you into implementing those two use cases I mentioned in Golang with the [slack-go](https://github.com/slack-go/slack) library and using [Slack Socket Mode](https://api.slack.com/apis/connections/socket).

> Why Socket Mode

You don't need a public server. In other words, your laptop, your raspberry pi, or a private server can host your bot. Socket mode is perfect for small Application that you do not intend to distribute via [App Directory](https://slack.com/apps)

### Sending Ephemeral greeting message when a user join a channel

[Ephemeral messages](https://api.slack.com/messaging/managing#ephemeral) are visible only by a specific user and are not persisted in the conversation. Those are ideal for reaction to user interaction such as joining a channel, answering with sensitive information, giving instruction, etc.

![](./assets/appmentionned2.gif)

### React to messages that mentioned your App or Bot with a message in your App Home

[App Home](https://api.slack.com/start/overview#app_home) is a dedicated space for your Slack Application. You can create a customized landing page, add an about page to document your App, and have a private message thread between your App and a user. Unlike Ephemeral messages, those messages will be persisted, meaning that all the tips and tricks you send will be easily accessible. Also, I would ing informative messages in the Slack App Home over sending them as a private message via a Bot. You want to convince your users that your App Home one place to find information about your App or Bot.

![](./assets/appmentionned1.gif)

## Configure your Application

Activate Socket Mode in the appropriate section.

![](./assets/socketmode.png)

## Create the project repository

First, create a new go project and import `slack-go`

```
+ controllers
`- greetingController.go
+ drivers
`- Slack.go
+ views
`+ greetingViewsAssets
  `- greeting.json
`- greetingViews.go
+ main.go
```

## Create the controller

## Create the view

## Final thoughts