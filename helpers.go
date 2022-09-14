package gottbot

// Edit is a message helper to bot.EditMessage
func (m *Message) Edit(bot *Bot, body NewMessageBody) (*SimpleQueryResult, error) {
	return bot.EditMessage(m.Body.Mid, body)
}

// Reply is a message helper to bot.SendMessage with reply message added
func (m *Message) Reply(bot *Bot, text string, opts *SendMessageOpts) (*SendMessageResult, error) {
	if opts == nil {
		opts = &SendMessageOpts{}
	}
	opts.Link = &MessageLink{
		Mid:  m.Body.Mid,
		Type: Reply,
	}
	return bot.SendMessage(m.Recipient.ChatId, text, opts)
}

// Forward is a message helper to bot.SendMessage with forward message added
func (m *Message) Forward(bot *Bot, chatId int64, text string, opts *SendMessageOpts) (*SendMessageResult, error) {
	if opts == nil {
		opts = &SendMessageOpts{}
	}
	opts.Link = &MessageLink{
		Mid:  m.Body.Mid,
		Type: Forward,
	}
	return bot.SendMessage(chatId, text, opts)
}
