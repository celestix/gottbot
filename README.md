# GoTTBot
GoTTBot is an asynchronous wrapper for the TamTam Bot API written in Golang. It provides all the methods and types that are available on the official TamTam Bot API and aims to keep everything type-safe and rid of generics.

You can use this package to create bots easily in golang, for any futher help you can check out the [documentations](https://pkg.go.dev/github.com/anonyindian/gottbot) or reach us through the following:
- Updates Channel: [![Channel](https://img.shields.io/badge/GoTTBot-Channel-dark)](https://tt.me/gottbot)
- Support Chat: [![Chat](https://img.shields.io/badge/Bot-Support%20Chat-red)](https://tt.me/botchat)

## **Key Features**
- **Easy to use**: Heavily inspired by the python-telegram-bot and gotgbot, GoTTBot is designed in such a way that even a beginner can make a bot with it easily.
- **Asynchronous**: GoTTBot processes each update in a separate goroutine to keep it asynchronous.  
- **Easy Migration**: Bots source codes can be easily migrated from Telegram Bot API to TamTam Bot API using the GoTTBot as it has intercept with the GoTGBot library.
- **Filters**: GoTTBot provides filters to make it easy for you to sort different type of updates in a managed way.

[![Go Reference](https://pkg.go.dev/badge/github.com/anonyindian/gottbot.svg)](https://pkg.go.dev/github.com/anonyindian/gottbot) [![GPLv3 license](https://img.shields.io/badge/License-GPLv3-blue.svg)](http://perso.crans.org/besson/LICENSE.html)

## Installation
You can download the library with the help of standard `go get` command.

```bash
go get github.com/anonyindian/gottbot
```

## Usage
You can find various examples in the [examples directory](./examples/), a simple echo example is as follows:
```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/anonyindian/gottbot"
	"github.com/anonyindian/gottbot/ext"
	"github.com/anonyindian/gottbot/filters"
	"github.com/anonyindian/gottbot/handlers"
)

func main() {
	bot, err := gottbot.NewBot(os.Getenv("TAMTAM_BOT_TOKEN"), nil)
	if err != nil {
		panic(err)
	}
	updater := ext.NewUpdater(nil)
	updater.StartPolling(bot, nil)

	dispatcher := updater.Dispatcher

	dispatcher.AddHandler(handlers.MessageHandler(filters.Message.All, echo))

	fmt.Println("Started example bot with long polling...")

	updater.Idle()
}

func echo(bot *gottbot.Bot, ctx *ext.Context) error {
	msg := ctx.EffectiveMessage
	_, err := msg.Reply(bot, msg.Body.Text, nil)
	if err != nil {
		log.Println("failed to send message:", err.Error())
	}
	return ext.EndGroups
}

```

## Bot Projects
This section includes some bots written with the GoTTBot library:
- [evalbot](https://github.com/TamTamBots/evalbot): A bot to eval codes of various programming languages.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update examples as appropriate.

## License
[![GPLv3](https://www.gnu.org/graphics/gplv3-127x51.png)](https://www.gnu.org/licenses/gpl-3.0.en.html)
<br>Licensed Under <a href="https://www.gnu.org/licenses/gpl-3.0.en.html">GNU General Public License v3</a>