package ext

import "github.com/anonyindian/gottbot"

type HandlerID string

type Handler interface {
	HandleUpdate(bot *gottbot.Bot, ctx *Context) error
	CheckUpdate(update *gottbot.Update) bool
	GetHandlerID() HandlerID
}
