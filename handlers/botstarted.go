package handlers

import (
	"fmt"

	"github.com/anonyindian/gottbot"
	"github.com/anonyindian/gottbot/ext"
)

type BotStarted struct {
	Response  Callback
	handlerID string
}

func BotStartedHandler(callback Callback) *BotStarted {
	return &BotStarted{
		Response: callback,
	}
}

func (m *BotStarted) CheckUpdate(update *gottbot.Update) bool {
	return update.GetUpdateType() == gottbot.UpdateTypeBotStarted
}

func (m *BotStarted) HandleUpdate(bot *gottbot.Bot, ctx *ext.Context) error {
	return m.Response(bot, ctx)
}

func (m *BotStarted) GetHandlerID() ext.HandlerID {
	if m.handlerID == "" {
		m.handlerID = makeHandlerID("bot_started", fmt.Sprintf("%v", m.Response))
	}
	return ext.HandlerID(m.handlerID)
}
