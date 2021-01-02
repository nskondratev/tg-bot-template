# tg-bot-template
Template repo with boilerplate code to write [Telegram bots in](https://core.telegram.org/bots/api) Go.

## What is this repo for?
I implemented several bots for Telegram and each time I started by writing/copying boilerplate code.
This repository is a template for a quick start of a new bot. It solves the following problems:
- App structure: follows the [Standard Go Project Layout](https://github.com/golang-standards/project-layout).
- Handy middlewares for HTTP-like processing updates from Telegram.
- Deploying bot as a [Google Cloud Function](https://cloud.google.com/functions) and convenient local debugging with polling updates.
- Structured logging with [zerolog](https://github.com/rs/zerolog) lib.

In short, this template will save you a couple of hours and allow you to immediately start implementing the bot's logic.

## Quickstart
1. Press "Use this template" button at the top or just follow [the link](https://github.com/nskondratev/tg-bot-template/generate).
2. Clone the generated repository to your machine.
3. Rename module and change import paths by calling the command (replace `github.com/author/newbot` with yours repo name):
```bash
./scripts/rename.sh github.com/author/newbot
```
4. Fill configuration in .env file:
```bash
mv .env.example .env && nano .env
```
5. Run your bot locally:
```bash
make run
```

To set up a webhook for receiving updates, fill the config in `.env` file and run the following command:
```bash
./scripts/set_webhook.sh
```

To clear a webhook run the same script with `-c` flag provided:
```bash
./scripts/set_webhook.sh -c
```

## Next steps
* Add domain-specific logic in [internal/app](./internal/app) package.
* Add update handlers in [internal/app/bot/handlers](./internal/app/bot/handlers) package.
* The library [telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api) is used to work with Telegram Bot API.

## Project structure
* `bin` - dir for compiled binary deps (look at the `tools` directory).
* `cmd/bot` - entry-point for running bot locally.
* `internal`:
  * `internal/app` - contains business-logic layer and adapters to external world in sub-packages.
  * `internal/bot` - wrappers to work with Telegram Bot API and middlewares implementation.
  * `internal/bot/handlers` - handlers for different update types.
  * `internal/bot/middleware` - middlewares for all updates.
  * `internal/boot` - bootstrapping code for bot creation (used in local entry-point and Google Cloud Function).
  * `internal/env` - utilities for getting env-vars values.
  * `internal/logger` - logger creation code and custom log fields constants.
* `scripts` - handy scripts for renaming module, changing import paths and setting up webhook URL.
* `tools` - binary deps of a project.
