package filters

import (
	"strings"

	"github.com/anonyindian/gottbot"
)

type MessageFilter func(m *gottbot.Message) bool

func (*message) All(_ *gottbot.Message) bool {
	return true
}

func (*message) Text(m *gottbot.Message) bool {
	return m.Body.Text != ""
}

func (*message) Prefix(prefix string) MessageFilter {
	return func(m *gottbot.Message) bool {
		return strings.HasPrefix(m.Body.Text, prefix)
	}
}

func (*message) Suffix(prefix string) MessageFilter {
	return func(m *gottbot.Message) bool {
		return strings.HasSuffix(m.Body.Text, prefix)
	}
}

func (*message) User(userId int64) MessageFilter {
	return func(m *gottbot.Message) bool {
		return m.Recipient.UserId == userId
	}
}

func (*message) Chat(chatId int64) MessageFilter {
	return func(m *gottbot.Message) bool {
		return m.Recipient.ChatId == chatId
	}
}

func (*message) ChatType(chatId gottbot.ChatType) MessageFilter {
	return func(m *gottbot.Message) bool {
		return m.Recipient.ChatType == chatId
	}
}

func (*message) IsReply(m *gottbot.Message) bool {
	return m.Link.Type == gottbot.Reply
}

func (*message) IsForwarded(m *gottbot.Message) bool {
	return m.Link.Type == gottbot.Forward
}
