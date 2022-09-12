package handlers

import (
	"fmt"
	"strings"

	"github.com/anonyindian/gottbot"
	"github.com/anonyindian/gottbot/ext"
	"github.com/anonyindian/gottbot/filters"
)

type Command struct {
	Prefix    []rune
	Command   string
	Response  Callback
	Filter    filters.MessageFilter
	handlerID string
}

func CommandHandler(command string, callback Callback) *Command {
	return &Command{
		Prefix:   []rune{'/'},
		Command:  command,
		Response: callback,
	}
}

func (c *Command) SetPrefix(prefix []rune) *Command {
	c.Prefix = prefix
	return c
}

func (c *Command) SetFilter(filter filters.MessageFilter) *Command {
	c.Filter = filter
	return c
}

func (c *Command) checkCommand(text string) bool {
	if text == "" {
		return false
	}
	arg := strings.ToLower(strings.Fields(text)[0])
	for _, prefix := range c.Prefix {
		if rune(arg[0]) != prefix {
			continue
		}
		if arg[1:] != c.Command {
			continue
		}
		return true
	}
	return false
}

func (c *Command) CheckUpdate(update *gottbot.Update) bool {
	switch update.GetUpdateType() {
	case gottbot.UpdateTypeMessageCreated:
		return c.checkCommand(update.MessageCreated.Message.Body.Text)
	}
	return false
}

func (c *Command) HandleUpdate(bot *gottbot.Bot, ctx *ext.Context) error {
	return c.Response(bot, ctx)
}

func (c *Command) GetHandlerID() ext.HandlerID {
	if c.handlerID == "" {
		c.handlerID = makeHandlerID("command", fmt.Sprintf("%v", c.Response))
	}
	return ext.HandlerID(c.handlerID)
}
