package handlers

import (
	"fmt"
	"time"

	"github.com/anonyindian/gottbot"
)

type Callback func(bot *gottbot.Bot, update *gottbot.Update) error

func makeHandlerID(name, suffix string) string {
	return fmt.Sprintf("%d_%s=%s", time.Now(), name, suffix)
}
