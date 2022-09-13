package handlers

import (
	"fmt"

	"github.com/anonyindian/gottbot"
	"github.com/anonyindian/gottbot/ext"
	"github.com/anonyindian/gottbot/filters"
)

type CallbackQuery struct {
	Response  Callback
	Filter    filters.CallbackQueryFilter
	handlerID string
}

func CallbackQueryHandler(filter filters.CallbackQueryFilter, callback Callback) *CallbackQuery {
	return &CallbackQuery{
		Response: callback,
		Filter:   filter,
	}
}

func (m *CallbackQuery) CheckUpdate(update *gottbot.Update) bool {
	switch update.GetUpdateType() {
	case gottbot.UpdateTypeMessageCallback:
		return true
	}
	return false
}

func (m *CallbackQuery) HandleUpdate(bot *gottbot.Bot, ctx *ext.Context) error {
	if !m.Filter(ctx.EffectiveQuery) {
		return ext.ContinueGroup
	}
	return m.Response(bot, ctx)
}

func (m *CallbackQuery) GetHandlerID() ext.HandlerID {
	if m.handlerID == "" {
		m.handlerID = makeHandlerID("callback_query", fmt.Sprintf("%v", m.Response))
	}
	return ext.HandlerID(m.handlerID)
}
