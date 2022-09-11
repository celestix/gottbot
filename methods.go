package gottbot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func (b *Bot) GetInfo() (*BotInfo, error) {
	data, err := b.MakeRequest(http.MethodGet, "me", url.Values{}, nil)
	if data != nil {
		defer data.Close()
	}
	if err != nil {
		return nil, err
	}
	var v BotInfo
	return &v, json.NewDecoder(data).Decode(&v)
}

func (b *Bot) PatchInfo(patch BotPatch) (*BotInfo, error) {
	bs, err := json.Marshal(patch)
	if err != nil {
		return nil, fmt.Errorf("failed to encode patch: %w", err)
	}
	data, err := b.MakeRequest(http.MethodPatch, "me", url.Values{}, bs)
	if data != nil {
		defer data.Close()
	}
	if err != nil {
		return nil, err
	}
	var v BotInfo
	return &v, json.NewDecoder(data).Decode(&v)
}

func (b *Bot) GetChats(opts *GetChatsOpts) (*ChatList, error) {
	if opts == nil {
		opts = &GetChatsOpts{}
	}
	u := url.Values{}
	if opts.Count != 0 {
		u.Add("count", strconv.FormatInt(int64(opts.Count), 10))
	}
	if opts.Marker != 0 {
		u.Add("marker", strconv.FormatInt(opts.Marker, 10))
	}
	data, err := b.MakeRequest(http.MethodGet, "chats", u, nil)
	if data != nil {
		defer data.Close()
	}
	if err != nil {
		return nil, err
	}
	var v ChatList
	return &v, json.NewDecoder(data).Decode(&v)
}

func (b *Bot) GetChatByLink(link string) (*Chat, error) {
	data, err := b.MakeRequest(http.MethodGet, fmt.Sprintf("chats/%s", link), url.Values{}, nil)
	if data != nil {
		defer data.Close()
	}
	if err != nil {
		return nil, err
	}
	var v Chat
	return &v, json.NewDecoder(data).Decode(&v)
}

func (b *Bot) GetChat(chatId int64) (*Chat, error) {
	data, err := b.MakeRequest(http.MethodGet, fmt.Sprintf("chats/%d", chatId), url.Values{}, nil)
	if data != nil {
		defer data.Close()
	}
	if err != nil {
		return nil, err
	}
	var v Chat
	return &v, json.NewDecoder(data).Decode(&v)
}

func (b *Bot) EditChat(chatId int64, patch ChatPatch) (*Chat, error) {
	bs, err := json.Marshal(patch)
	if err != nil {
		return nil, fmt.Errorf("failed to encode patch: %w", err)
	}
	data, err := b.MakeRequest(http.MethodPatch, fmt.Sprintf("chats/%d", chatId), url.Values{}, bs)
	if data != nil {
		defer data.Close()
	}
	if err != nil {
		return nil, err
	}
	var v Chat
	return &v, json.NewDecoder(data).Decode(&v)
}

func (b *Bot) SendAction(chatId int64, action SenderAction) (*SimpleQueryResult, error) {
	data, err := b.MakeRequest(http.MethodPost, fmt.Sprintf("chats/%d/actions", chatId), url.Values{}, []byte(action))
	if data != nil {
		defer data.Close()
	}
	if err != nil {
		return nil, err
	}
	var v SimpleQueryResult
	return &v, json.NewDecoder(data).Decode(&v)
}

func (b *Bot) GetPinnedMessage(chatId int64) (*GetPinnedMessageResult, error) {
	data, err := b.MakeRequest(http.MethodGet, fmt.Sprintf("chats/%d/pin", chatId), url.Values{}, nil)
	if data != nil {
		defer data.Close()
	}
	if err != nil {
		return nil, err
	}
	var v GetPinnedMessageResult
	return &v, json.NewDecoder(data).Decode(&v)
}

func (b *Bot) PinMessage(chatId int64, body PinMessageBody) (*SimpleQueryResult, error) {
	bs, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %w", err)
	}
	data, err := b.MakeRequest(http.MethodPut, fmt.Sprintf("chats/%d/pin", chatId), url.Values{}, bs)
	if data != nil {
		defer data.Close()
	}
	if err != nil {
		return nil, err
	}
	var v SimpleQueryResult
	return &v, json.NewDecoder(data).Decode(&v)
}

func (b *Bot) UnpinMessage(chatId int64) (*SimpleQueryResult, error) {
	data, err := b.MakeRequest(http.MethodDelete, fmt.Sprintf("chats/%d/pin", chatId), url.Values{}, nil)
	if data != nil {
		defer data.Close()
	}
	if err != nil {
		return nil, err
	}
	var v SimpleQueryResult
	return &v, json.NewDecoder(data).Decode(&v)
}

func (b *Bot) GetChatMembership(chatId int64) (*ChatMember, error) {
	data, err := b.MakeRequest(http.MethodGet, fmt.Sprintf("chats/%d/members/me", chatId), url.Values{}, nil)
	if data != nil {
		defer data.Close()
	}
	if err != nil {
		return nil, err
	}
	var v ChatMember
	return &v, json.NewDecoder(data).Decode(&v)
}

func (b *Bot) LeaveChat(chatId int64) (*SimpleQueryResult, error) {
	data, err := b.MakeRequest(http.MethodDelete, fmt.Sprintf("chats/%d/members/me", chatId), url.Values{}, nil)
	if data != nil {
		defer data.Close()
	}
	if err != nil {
		return nil, err
	}
	var v SimpleQueryResult
	return &v, json.NewDecoder(data).Decode(&v)
}

func (b *Bot) GetChatAdmins(chatId int64) (*ChatMembersList, error) {
	data, err := b.MakeRequest(http.MethodGet, fmt.Sprintf("chats/%d/members/admins", chatId), url.Values{}, nil)
	if data != nil {
		defer data.Close()
	}
	if err != nil {
		return nil, err
	}
	var v ChatMembersList
	return &v, json.NewDecoder(data).Decode(&v)
}

func (b *Bot) GetChatMembers(chatId int64, opts *GetChatMembersOpts) (*ChatMembersList, error) {
	if opts == nil {
		opts = &GetChatMembersOpts{}
	}
	u := url.Values{}
	if opts.UserIds != nil {
		bs, err := json.Marshal(opts.UserIds)
		if err != nil {
			return nil, fmt.Errorf("failed to encode userIds: %w", err)
		}
		u.Add("user_ids", string(bs))
	}
	if opts.Count != 0 {
		u.Add("count", strconv.FormatInt(int64(opts.Count), 10))
	}
	if opts.Marker != 0 {
		u.Add("marker", strconv.FormatInt(opts.Marker, 10))
	}
	data, err := b.MakeRequest(http.MethodGet, fmt.Sprintf("chats/%d/members", chatId), u, nil)
	if data != nil {
		defer data.Close()
	}
	if err != nil {
		return nil, err
	}
	var v ChatMembersList
	return &v, json.NewDecoder(data).Decode(&v)
}

func (b *Bot) AddMembers(chatId int64, userIds []int64) (*SimpleQueryResult, error) {
	bs, err := json.Marshal(userIds)
	if err != nil {
		return nil, fmt.Errorf("failed to encode userIds: %w", err)
	}
	data, err := b.MakeRequest(http.MethodPost, fmt.Sprintf("chats/%d/members", chatId), url.Values{}, bs)
	if data != nil {
		defer data.Close()
	}
	if err != nil {
		return nil, err
	}
	var v SimpleQueryResult
	return &v, json.NewDecoder(data).Decode(&v)
}

func (b *Bot) RemoveMember(chatId int64, userId int64, block bool) (*SimpleQueryResult, error) {
	u := url.Values{}
	u.Add("user_id", strconv.FormatInt(userId, 10))
	u.Add("block", strconv.FormatBool(block))
	data, err := b.MakeRequest(http.MethodDelete, fmt.Sprintf("chats/%d/members", chatId), u, nil)
	if data != nil {
		defer data.Close()
	}
	if err != nil {
		return nil, err
	}
	var v SimpleQueryResult
	return &v, json.NewDecoder(data).Decode(&v)
}

func (b *Bot) GetMessages(opts *GetMessagesOpts) (*MessageList, error) {
	if opts == nil {
		opts = &GetMessagesOpts{}
	}
	u := url.Values{}
	if opts.MessageIds != nil {
		bs, err := json.Marshal(opts.MessageIds)
		if err != nil {
			return nil, fmt.Errorf("failed to encode userIds: %w", err)
		}
		u.Add("message_ids", string(bs))
	}
	if opts.Count != 0 {
		u.Add("count", strconv.FormatInt(int64(opts.Count), 10))
	}
	if opts.ChatId != 0 {
		u.Add("chat_id", strconv.FormatInt(opts.ChatId, 10))
	}
	if opts.From != 0 {
		u.Add("from", strconv.FormatInt(opts.From, 10))
	}
	if opts.To != 0 {
		u.Add("to", strconv.FormatInt(opts.To, 10))
	}
	data, err := b.MakeRequest(http.MethodGet, "messages", u, nil)
	if data != nil {
		defer data.Close()
	}
	if err != nil {
		return nil, err
	}
	var v MessageList
	return &v, json.NewDecoder(data).Decode(&v)
}

func (b *Bot) SendMessage(chatId int64, text string, opts *SendMessageOpts) (*SendMessageResult, error) {
	if opts == nil {
		opts = &SendMessageOpts{}
	}
	u := url.Values{}
	if chatId != 0 {
		u.Add("chat_id", strconv.FormatInt(chatId, 10))
	}
	// if opts.UserId != 0 {
	// 	u.Add("from", strconv.FormatInt(opts.UserId, 10))
	// }
	if opts.DisableLinkPreview {
		u.Add("disable_link_preview", strconv.FormatBool(opts.DisableLinkPreview))
	}

	bs, err := json.Marshal(SendMessageBody{
		Text:        text,
		Attachments: opts.Attachments,
		Link:        opts.Link,
		Format:      opts.Format,
		Notify:      opts.Notify,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to encode SendMessageBody: %w", err)
	}

	data, err := b.MakeRequest(http.MethodPost, "messages", u, bs)
	if data != nil {
		defer data.Close()
	}
	if err != nil {
		return nil, err
	}
	var v SendMessageResult
	return &v, json.NewDecoder(data).Decode(&v)
}

func (b *Bot) EditMessage(messageId string, body NewMessageBody) (*SimpleQueryResult, error) {
	u := url.Values{}
	u.Add("message_id", messageId)
	bs, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to encode NewMessageBody: %w", err)
	}
	data, err := b.MakeRequest(http.MethodPut, "messages", u, bs)
	if data != nil {
		defer data.Close()
	}
	if err != nil {
		return nil, err
	}
	var v SimpleQueryResult
	return &v, json.NewDecoder(data).Decode(&v)
}

func (b *Bot) DeleteMessage(messageId string) (*SimpleQueryResult, error) {
	u := url.Values{}
	u.Add("message_id", messageId)
	data, err := b.MakeRequest(http.MethodDelete, "messages", u, nil)
	if data != nil {
		defer data.Close()
	}
	if err != nil {
		return nil, err
	}
	var v SimpleQueryResult
	return &v, json.NewDecoder(data).Decode(&v)
}

func (b *Bot) GetMessage(messageId string) (*Message, error) {
	data, err := b.MakeRequest(http.MethodGet, fmt.Sprintf("messages/%s", messageId), url.Values{}, nil)
	if data != nil {
		defer data.Close()
	}
	if err != nil {
		return nil, err
	}
	var v Message
	return &v, json.NewDecoder(data).Decode(&v)
}

func (b *Bot) AnswerOnCallback(callbackId string, body CallbackAnswer) (*SimpleQueryResult, error) {
	u := url.Values{}
	u.Add("callback_id", callbackId)
	bs, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to encode CallbackAnswer: %w", err)
	}
	data, err := b.MakeRequest(http.MethodPost, "answers", u, bs)
	if data != nil {
		defer data.Close()
	}
	if err != nil {
		return nil, err
	}
	var v SimpleQueryResult
	return &v, json.NewDecoder(data).Decode(&v)
}

func (b *Bot) ConstructMessage(sessionId string, body ConstructorAnswer) (*SimpleQueryResult, error) {
	u := url.Values{}
	u.Add("session_id", sessionId)
	bs, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to encode ConstructorAnswer: %w", err)
	}
	data, err := b.MakeRequest(http.MethodPost, "answers/constructor", u, bs)
	if data != nil {
		defer data.Close()
	}
	if err != nil {
		return nil, err
	}
	var v SimpleQueryResult
	return &v, json.NewDecoder(data).Decode(&v)
}

func (b *Bot) GetUpdates(opts *GetUpdatesOpts) (*UpdateList, error) {
	u := url.Values{}
	if opts != nil {
		if opts.Limit != 0 {
			u.Add("limit", strconv.FormatInt(int64(opts.Limit), 10))
		}
		if opts.Marker != 0 {
			u.Add("marker", strconv.FormatInt(opts.Marker, 10))
		}
		if opts.Timeout != 0 {
			u.Add("timeout", strconv.FormatInt(int64(opts.Timeout), 10))
		}
		if opts.Types != nil {
			bs, err := json.Marshal(opts.Types)
			if err != nil {
				return nil, fmt.Errorf("failed to encode GetUpdateOpts.Types: %w", err)
			}
			u.Add("types", string(bs))
		}
	}
	data, err := b.MakeRequest(http.MethodGet, "updates", u, nil)
	if data != nil {
		defer data.Close()
	}
	if err != nil {
		return nil, err
	}
	var v UpdateList
	return &v, json.NewDecoder(data).Decode(&v)
}

func (b *Bot) GetSubscriptions() (*GetSubscriptionsResult, error) {
	data, err := b.MakeRequest(http.MethodGet, "subscriptions", url.Values{}, nil)
	if data != nil {
		defer data.Close()
	}
	if err != nil {
		return nil, err
	}
	var v GetSubscriptionsResult
	return &v, json.NewDecoder(data).Decode(&v)
}

func (b *Bot) Subscribe(body SubscriptionRequestBody) (*SimpleQueryResult, error) {
	bs, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to encode ConstructorAnswer: %w", err)
	}
	data, err := b.MakeRequest(http.MethodPost, "subscriptions", url.Values{}, bs)
	if data != nil {
		defer data.Close()
	}
	if err != nil {
		return nil, err
	}
	var v SimpleQueryResult
	return &v, json.NewDecoder(data).Decode(&v)
}

func (b *Bot) Unsubscribe(webhookUrl string) (*SimpleQueryResult, error) {
	u := url.Values{}
	u.Add("url", webhookUrl)
	data, err := b.MakeRequest(http.MethodDelete, "subscriptions", u, nil)
	if data != nil {
		defer data.Close()
	}
	if err != nil {
		return nil, err
	}
	var v SimpleQueryResult
	return &v, json.NewDecoder(data).Decode(&v)
}
