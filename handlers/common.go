package handlers

import (
	"fmt"
	"time"

	"github.com/anonyindian/gottbot"
	"github.com/anonyindian/gottbot/ext"
)

type Callback func(bot *gottbot.Bot, ctx *ext.Context) error

func makeHandlerID(name, suffix string) string {
	return fmt.Sprintf("%d_%s=%s", time.Now().Unix(), name, suffix)
}
