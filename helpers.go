package gottbot

func (m *Message) Edit(bot *Bot, body NewMessageBody) (*SimpleQueryResult, error) {
	return bot.EditMessage(m.Body.Mid, body)
}

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
