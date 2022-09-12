package handlers

import (
	"fmt"

	"github.com/anonyindian/gottbot"
	"github.com/anonyindian/gottbot/ext"
)

type BotAdded struct {
	Response  Callback
	handlerID string
}

func BotAddedHandler(callback Callback) *BotAdded {
	return &BotAdded{
		Response: callback,
	}
}

func (m *BotAdded) CheckUpdate(update *gottbot.Update) bool {
	return update.GetUpdateType() == gottbot.UpdateTypeBotAdded
}

func (m *BotAdded) HandleUpdate(bot *gottbot.Bot, ctx *ext.Context) error {
	return m.Response(bot, ctx)
}

func (m *BotAdded) GetHandlerID() ext.HandlerID {
	if m.handlerID == "" {
		m.handlerID = makeHandlerID("bot_added", fmt.Sprintf("%v", m.Response))
	}
	return ext.HandlerID(m.handlerID)
}
