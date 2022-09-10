package handlers

import (
	"fmt"
	"strings"

	"github.com/anonyindian/gottbot"
	"github.com/anonyindian/gottbot/ext"
)

type Command struct {
	Prefix    []rune
	Command   string
	Response  Callback
	handlerID string
}

func CommandHandler(command string, callback Callback) *Command {
	return &Command{
		Prefix:   []rune{'/'},
		Command:  command,
		Response: callback,
	}
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
	switch update.UpdateType {
	case gottbot.UpdateMessageCreated:
		return c.checkCommand(update.Message.Body.Text)
	}
	return false
}

func (c *Command) HandleUpdate(bot *gottbot.Bot, update *gottbot.Update) error {
	return c.Response(bot, update)
}

func (c *Command) GetHandlerID() ext.HandlerID {
	if c.handlerID == "" {
		c.handlerID = makeHandlerID("command", fmt.Sprintf("%v", c.Response))
	}
	return ext.HandlerID(c.handlerID)
}
