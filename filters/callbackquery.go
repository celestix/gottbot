package filters

import "github.com/anonyindian/gottbot"

type CallbackQueryFilter func(m *gottbot.Callback) bool

func (m *callbackQuery) All(_ *gottbot.Callback) bool {
	return true
}
