package ext

import (
	"github.com/anonyindian/gottbot"
)

type Context struct {
	EffectiveUser    *gottbot.User
	EffectiveMessage *gottbot.Message
	EffectiveQuery   *gottbot.Callback
	EffectiveChatId  int64
	Data             map[string]any
	*gottbot.Update
}

func NewContext(u *gottbot.Update) *Context {
	ctx := &Context{
		Update: u,
	}
	switch {
	case u.MessageCreated != nil:
		ctx.EffectiveMessage = u.MessageCreated.Message
		ctx.EffectiveUser = u.MessageCreated.Message.Sender
		ctx.EffectiveChatId = u.MessageCreated.Message.Recipient.ChatId

	case u.MessageCallback != nil:
		ctx.EffectiveQuery = u.MessageCallback.Callback
		ctx.EffectiveMessage = u.MessageCallback.Message
		ctx.EffectiveUser = u.MessageCallback.Message.Sender
		ctx.EffectiveChatId = u.MessageCallback.Message.Recipient.ChatId

	case u.MessageEdited != nil:
		ctx.EffectiveMessage = u.MessageEdited.Message
		ctx.EffectiveUser = u.MessageEdited.Message.Sender
		ctx.EffectiveChatId = u.MessageEdited.Message.Recipient.ChatId

	case u.BotAdded != nil:
		ctx.EffectiveUser = u.BotAdded.User
		ctx.EffectiveChatId = u.BotAdded.ChatId

	case u.BotRemoved != nil:
		ctx.EffectiveUser = u.BotRemoved.User
		ctx.EffectiveChatId = u.BotAdded.ChatId

	case u.BotStarted != nil:
		ctx.EffectiveUser = u.BotStarted.User
		ctx.EffectiveChatId = u.BotStarted.ChatId

	case u.UserAdded != nil:
		ctx.EffectiveUser = u.UserAdded.User
		ctx.EffectiveChatId = u.UserAdded.ChatId

	case u.UserRemoved != nil:
		ctx.EffectiveUser = u.UserRemoved.User
		ctx.EffectiveChatId = u.UserRemoved.ChatId

	case u.MessageConstructed != nil:
		ctx.EffectiveMessage = u.MessageConstructed.Message
		ctx.EffectiveUser = u.MessageConstructed.Message.Sender
		ctx.EffectiveChatId = u.MessageConstructed.Message.Recipient.ChatId

	case u.MessageConstructionRequest != nil:
		ctx.EffectiveUser = u.MessageConstructionRequest.User

	case u.ChatTitleChanged != nil:
		ctx.EffectiveUser = u.ChatTitleChanged.User
		ctx.EffectiveChatId = u.ChatTitleChanged.ChatId

	case u.MessageChatCreated != nil:
		ctx.EffectiveChatId = u.MessageChatCreated.Chat.ChatId

	}
	return ctx
}
