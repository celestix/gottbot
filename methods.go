package gottbot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
)

// Returns info about current bot.
// Current bot can be identified by access token.
// Method returns bot identifier, name and avatar (if any)
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

// Edits current bot info. Fill only the fields you want to update.
// All remaining fields will stay untouched
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

// Returns information about chats that bot participated in:
// a result list and marker points to the next page
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

// Returns chat/channel information by its public link or dialog with user by username
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

// Returns info about chat.
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

// Edits chat info: title, icon, etc…
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

// Send bot action to chat.
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

// Get pinned message in chat or channel.
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

// Pins message in chat or channel.
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

// Unpins message in chat or channel.
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

// Returns chat membership info for current bot
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

// Removes bot from chat members.
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

// Returns all chat administrators. Bot must be administrator in requested chat.
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

// Returns users participated in chat.
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

// Adds members to chat. Additional permissions may require.
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

// Removes member from chat. Additional permissions may require.
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

// Returns messages in chat: result page and marker referencing to the next page.
//
// Messages traversed in reverse direction so the latest message in chat will be first in result array.
// Therefore if you use from and to parameters, to must be less than from
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

// Sends a message to a chat. As a result for this method new message identifier returns.
//
// Important notice:
// It may take time for the server to process your file (audio/video or any binary).
// While a file is not processed you can't attach it.
// It means the last step will fail with 400 error.
// Try to send a message again until you'll get a successful result.
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

// Updated message should be sent as NewMessageBody in a request body.
// In case attachments field is null, the current message attachments won’t be changed.
// In case of sending an empty list in this field, all attachments will be deleted.
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

// Deletes message in a dialog or in a chat if bot has permission to delete messages.
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

// Returns single message by its identifier.
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

// This method should be called to send an answer after a user has clicked the button.
// The answer may be an updated message or/and a one-time user notification.
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

// Sends answer on construction request.
// Answer can contain any prepared message and/or keyboard to help user interact with bot.
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

// You can use this method for getting updates in case your bot is not subscribed to WebHook.
// The method is based on long polling.
//
// Every update has its own sequence number. marker property in response points to the next upcoming update.
//
// All previous updates are considered as committed after passing marker parameter.
// If marker parameter is not passed, your bot will get all updates happened after the last commitment.
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

// In case your bot gets data via WebHook, the method returns list of all subscriptions
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

// Subscribes bot to receive updates via WebHook. After calling this method, the bot will receive notifications about new events in chat rooms at the specified URL.
//
// Your server must be listening on one of the following ports: 80, 8080, 443, 8443, 16384-32383
func (b *Bot) Subscribe(body SubscriptionRequestBody) (*SimpleQueryResult, error) {
	bs, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to encode SubscriptionRequestBody: %w", err)
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

// Unsubscribes bot from receiving updates via WebHook.
// After calling the method, the bot stops receiving notifications about new events.
// Notification via the long-poll API becomes available for the bot
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

func (b *Bot) getUploadUrl(uploadType UploadType) (*UploadEndpoint, error) {
	u := url.Values{}
	u.Add("type", string(uploadType))
	data, err := b.MakeRequest(http.MethodPost, "uploads", u, nil)
	if data != nil {
		defer data.Close()
	}
	if err != nil {
		return nil, err
	}
	var v UploadEndpoint
	return &v, json.NewDecoder(data).Decode(&v)
}

// FileInfo is the struct used to deliver the information of a file to bot.Upload
type FileInfo struct {
	// Name of the file
	Name string
	// File Reader
	File io.Reader
}

// Returns the URL for the subsequent file upload.
func (b *Bot) Upload(uploadType UploadType, fileInfo *FileInfo) (Payload, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile(string(uploadType), fileInfo.Name)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, fileInfo.File)
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	endpoint, err := b.getUploadUrl(uploadType)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, endpoint.Url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := b.client.Do(req)
	if err != nil {
		return nil, err
	}

	switch uploadType {
	case UploadTypeFile:
		v := FilePayload{}
		return &v, json.NewDecoder(resp.Body).Decode(&v)
	case UploadTypeImage:
		v := ImagePayload{}
		return &v, json.NewDecoder(resp.Body).Decode(&v)
	case UploadTypeVideo:
		v := VideoPayload{}
		return &v, json.NewDecoder(resp.Body).Decode(&v)
	case UploadTypeAudio:
		v := AudioPayload{}
		return &v, json.NewDecoder(resp.Body).Decode(&v)
	default:
		return nil, fmt.Errorf("failed to upload: Unknown UploadType: %s", uploadType)
	}
}
