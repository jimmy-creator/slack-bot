#  Introduction

Section Overview: In this section, introduces the topic of building a Slack bot using Golang and explains the different levels of integration with the Slack API.

## Levels of Integration

- There are three levels of integration with the Slack API.
- The first level is direct API integration, which involves creating Golang functions that talk to the APIs directly.
- The second level is using a library or SDK for Golang that someone has already written on top of the Slack APIs.
- The third level is socket programming, which involves relying on a package that someone has written for Golang specifically for creating Slack bots.

#  Socket Programming

Section Overview: In this section, the speaker explains why socket programming is necessary for building a Slack bot and introduces the Slacker library.

## Using Socket Programming

- Building a Slack bot requires socket programming because it needs to be able to read and write messages in real-time.
- The Slacker library is used for socket programming in Golang specifically for creating Slack bots.

#  Process Overview

Section Overview: In this section, the speaker provides an overview of the process involved in building a Slack bot and highlights some potential challenges.

## Creating a Slack Account

- The first step in building a Slack bot is to create a new account on slack.com.

## Dashboard Settings

- A significant portion (50% - 60%) of the process involves setting up various settings on the dashboard.
- This includes creating channels, adding users, and generating API keys and tokens.

## Writing Code

- Once all necessary settings have been configured on the dashboard, writing code can begin.
- While writing actual Golang code may not be complex, understanding how to create a bot in the dashboard and obtaining API keys and tokens can be challenging.
# [t=261s] Creating a Slack Bot with Go and Socket Mode

Section Overview: In this section, the speaker introduces the necessary steps to create a Slack bot using Go and Socket Mode. The speaker explains how to create an app from scratch or from a template, enable socket mode, copy the app token, enable event subscriptions, add permissions for events, copy the bot token, install the bot to your workspace, and start building your bot.

## Creating a Slack Bot

- To create a Slack bot using Go and Socket Mode:
  - Create a Slack account.
  - Come to api.slack.com.
  - Create an app from scratch or from a template.
    - Templates can be useful when creating subsequent bots.
  - Enable socket mode in the left menu.
    - Copy the app token provided by socket mode.
  - Enable event subscriptions and add permissions for events needed by your bot.
    - Permissions include reading and writing messages in chat channels.
    - Save changes made to event subscriptions and permissions pages.
    - Copy the bot token provided on these pages.

## Installing Your Bot

- To install your bot:
  - Install it on your local server instead of remote server.
  - Mention it on a channel where you want it to work (e.g., general channel).
  - Pass it some command that you have written for it. 

## Talking to Your Bot

- To talk to your bot:
  - Mention it on a channel where you want it to work (e.g., general channel).
  - Pass it some command that you have written for it.
#  Building a Slack Bot: Part 1

Section Overview: In this section, the speaker introduces the upcoming videos in the series and explains that they will cover building a Slack bot and integrating with other APIs that Slack has.

## Creating a Slack App

-  To create a Slack app, go to api.slack.com and create an account and workspace.
-  Enable Socket Mode by finding it in the menu on the left and generating a socket token.
-  Switch on Event Subscriptions to subscribe to bot events from here or add them in OAuth permissions.
-  Add scopes such as app read, map mentions read, view messaging, channels history, channels read, chat right channel's message.

#  Building a Slack Bot: Part 1

Section Overview: In this section, the speaker emphasizes the importance of giving your bot the correct permissions for it to work properly.

## OAuth and Permissions

-  In OAuth and Permissions, ensure you have all necessary scopes including app mentions read, channels history, channels read, chat right i'm history i am read i'm writing mpm history.
#  Required Scopes and Permissions

Section Overview: In this section, the speaker discusses the scopes and permissions required for a chatbot to function properly. They provide examples of when certain files need to be read or written, and how a chatbot can upload files to a channel. The speaker emphasizes the importance of selecting all necessary events and permissions to avoid issues with code functionality.

## Required Scopes

-  Files read and write scopes are needed in certain situations.
-  MPIM read scope is necessary for group direct messages.
-  Messages.groups scope should also be added.

## Necessary Permissions

-  MPIM read permission is required.
-  MPIM history permission is necessary.
-  Chat right always permission is needed.
-  Messages.groups permission should be added.

#  Installing Bot Token

Section Overview: In this section, the speaker explains how to install a bot token in OAuth and permissions. They stress the importance of copying the bot token into OAuth and permissions after installation, as well as selecting all necessary events and permissions.

## Installing Bot Token

-  Copy bot token from OAuth and permissions after installation.
-  Reinstall bot if any changes are made to it.

#  Writing Golang Code for Slack Chatbot

Section Overview: In this section, the speaker demonstrates how to write Golang code for a Slack chatbot using VS Code. They explain how to create a directory called "slack test" using Go mod init, install socket package "slacker," and use it in the code.

## Writing Golang Code

-  Create a directory called "slack test" using Go mod init.
-  Install socket package "slacker."
-  Use "slacker" package in the code.
#  Creating a Slack Bot with Golang

Section Overview: In this section, the speaker introduces the process of creating a Slack bot using Golang. They explain that they will create a file called "go main.go" and follow three standard steps: package main, import packages, and have a function main.

## Setting up the Bot

-  The speaker creates a function called "print command events" to print all the command events given to Slack.
-  To initialize the bot, they use the Slacker library and set environment variables for slack bot token and slack app token.
-  The speaker explains that they keep their code modular so that in the future if they want to have a config file or set environment in any other way, it won't require many changes.
-  After creating their bot, they call "go print command events" to show whatever commands are passed to their bot.

## Printing Command Events

-  The speaker explains that they want all commands passed to their bot to show up in their terminal so that they know everything is working fine.
-  They range over analytics channel and print out the timestamp of the command, the command itself, parameters, and event.
#  Printing Commands and Handling Events

Section Overview: In this section, the speaker discusses how to print commands and handle events in a Go program.

## Printing Commands

- To print out a command, use `fmp.println` followed by the given parameters.
- Use `println` to print out the event as well.

## Handling Events

- Range over the events to get them one by one.
- Print out all values inside each event.
- Add an empty `println`.
- Import necessary packages such as context and log for error logging.

## The Command Function

- The command function is the most important function in this project.
- Pass the address of `slacker.commandDefinition` into it.
- Inside the definition, add a handler to handle the event with a function that sends a response using `response.reply`.
- Test your bot by mentioning its name in Slack.
#  Building Advanced Slack Bots

Section Overview: In this section, the speaker talks about building advanced slack bots that can perform various tasks such as providing weather updates, crypto market information, stock market information and help with getting tickets from Jira.

## Features of Advanced Slack Bots

-  Building advanced slack bots that can perform various tasks such as:
  - Providing weather updates
  - Getting crypto market information
  - Getting stock market information
  - Helping with getting tickets from Jira
  
