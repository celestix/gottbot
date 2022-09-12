package handlers

import (
	"fmt"

	"github.com/anonyindian/gottbot"
	"github.com/anonyindian/gottbot/ext"
	"github.com/anonyindian/gottbot/filters"
)

type Message struct {
	Response    Callback
	Filter      filters.MessageFilter
	AllowEdited bool
	handlerID   string
}

func MessageHandler(filter filters.MessageFilter, callback Callback) *Message {
	return &Message{
		Response: callback,
		Filter:   filter,
	}
}

func (m *Message) CheckUpdate(update *gottbot.Update) bool {
	switch update.GetUpdateType() {
	case gottbot.UpdateTypeMessageCreated:
		return true
	case gottbot.UpdateTypeMessageEdited:
		return m.AllowEdited
	}
	return false
}

func (m *Message) HandleUpdate(bot *gottbot.Bot, ctx *ext.Context) error {
	return m.Response(bot, ctx)
}

func (m *Message) GetHandlerID() ext.HandlerID {
	if m.handlerID == "" {
		m.handlerID = makeHandlerID("message", fmt.Sprintf("%v", m.Response))
	}
	return ext.HandlerID(m.handlerID)
}
