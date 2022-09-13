package gottbot

import (
	"encoding/json"
)

// Defines values for ChatAdminPermission.
const (
	AddAdmins        ChatAdminPermission = "add_admins"
	AddRemoveMembers ChatAdminPermission = "add_remove_members"
	ChangeChatInfo   ChatAdminPermission = "change_chat_info"
	PinMessage       ChatAdminPermission = "pin_message"
	ReadAllMessages  ChatAdminPermission = "read_all_messages"
	Write            ChatAdminPermission = "write"
)

// Defines values for MessageLinkType.
const (
	Forward MessageLinkType = "forward"
	Reply   MessageLinkType = "reply"
)

// Defines values for TextFormat.
const (
	Html     TextFormat = "html"
	Markdown TextFormat = "markdown"
)

const (
	UpdateTypeMessageCreated             UpdateType = "message_created"
	UpdateTypeMessageRemoved             UpdateType = "message_removed"
	UpdateTypeMessageCallback            UpdateType = "message_callback"
	UpdateTypeMessageEdited              UpdateType = "message_edited"
	UpdateTypeBotAdded                   UpdateType = "bot_added"
	UpdateTypeBotRemoved                 UpdateType = "bot_removed"
	UpdateTypeBotStarted                 UpdateType = "bot_started"
	UpdateTypeUserAdded                  UpdateType = "user_added"
	UpdateTypeUserRemoved                UpdateType = "user_removed"
	UpdateTypeChatTitleChanged           UpdateType = "chat_title_changed"
	UpdateTypeMessageConstructionRequest UpdateType = "message_construction_request"
	UpdateTypeMessageConstructed         UpdateType = "message_constructed"
	UpdateTypeMessageChatCreated         UpdateType = "message_chat_created"
)

var AllUpdates = []UpdateType{
	UpdateTypeBotAdded, UpdateTypeBotRemoved, UpdateTypeBotStarted,
	UpdateTypeChatTitleChanged,

	UpdateTypeMessageCallback, UpdateTypeMessageChatCreated, UpdateTypeMessageConstructed, UpdateTypeMessageConstructionRequest,
	UpdateTypeMessageCreated, UpdateTypeMessageEdited, UpdateTypeMessageRemoved,

	UpdateTypeUserAdded, UpdateTypeUserRemoved,
}

type UpdateType string

// ActionRequestBody defines model for ActionRequestBody.
type ActionRequestBody struct {
	// Action Different actions to send to chat members
	Action SenderAction `json:"action"`
}

// AttachmentRequest Request to attach some data to message
type AttachmentRequest struct {
	Payload Payload `json:"payload"`
}

func (a AttachmentRequest) MarshalJSON() ([]byte, error) {
	type temp AttachmentRequest
	v := struct {
		Type string `json:"type"`
		temp
	}{
		Type: a.Payload.GetPayloadType(),
		temp: temp(a),
	}
	return json.Marshal(v)
}

func (a *AttachmentRequest) UnmarshalJSON(b []byte) error {
	v := struct {
		Type string `json:"type"`
	}{}
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	switch v.Type {
	case "inline_keyboard":
		t := struct {
			Payload *ButtonsPayload `json:"payload"`
		}{}
		err := json.Unmarshal(b, &t)
		if err != nil {
			return err
		}
		a.Payload = t.Payload
	case "image":
		t := struct {
			Payload *ImagePayload `json:"payload"`
		}{}
		err := json.Unmarshal(b, &t)
		if err != nil {
			return err
		}
		a.Payload = t.Payload
	case "video":
		t := struct {
			Payload   *VideoPayload `json:"payload"`
			Thumbnail *Image        `json:"thumbnail,omitempty"`
			Duration  int64         `json:"duration,omitempty"`
		}{}
		err := json.Unmarshal(b, &t)
		if err != nil {
			return err
		}
		t.Payload.Thumbnail = t.Thumbnail
		t.Payload.Duration = t.Duration
		a.Payload = t.Payload
	case "audio":
		t := struct {
			Payload *AudioPayload `json:"payload"`
		}{}
		err := json.Unmarshal(b, &t)
		if err != nil {
			return err
		}
		a.Payload = t.Payload
	case "file":
		t := struct {
			Payload  *FilePayload `json:"payload"`
			Filename string       `json:"filename,omitempty"`
			Size     int64        `json:"size,omitempty"`
		}{}
		err := json.Unmarshal(b, &t)
		if err != nil {
			return err
		}
		t.Payload.Filename = t.Filename
		t.Payload.Size = t.Size
		a.Payload = t.Payload
	case "contact":
		t := struct {
			Payload *ContactPayload `json:"payload"`
		}{}
		err := json.Unmarshal(b, &t)
		if err != nil {
			return err
		}
		a.Payload = t.Payload
	case "sticker":
		t := struct {
			Payload *StickerPayload `json:"payload"`
			Width   int             `json:"width,omitempty"`
			Height  int             `json:"height,omitempty"`
		}{}
		err := json.Unmarshal(b, &t)
		if err != nil {
			return err
		}
		t.Payload.Width = t.Width
		t.Payload.Height = t.Height
		a.Payload = t.Payload
	case "location":
		// Well, This is weird but idk why tt bot api is sending it like that.
		// Maybe, A bug?
		t := LocationPayload{}
		err := json.Unmarshal(b, &t)
		if err != nil {
			return err
		}
		a.Payload = &t
	case "share":
		t := struct {
			Payload *SharePayload `json:"payload"`
		}{}
		err := json.Unmarshal(b, &t)
		if err != nil {
			return err
		}
		a.Payload = t.Payload
	}
	return nil
}

// BotCommand defines model for BotCommand.
type BotCommand struct {
	// Description Optional command description
	Description *string `json:"description"`

	// Name Command name
	Name string `json:"name"`
}

// BotInfo defines model for BotInfo.
type BotInfo struct {
	// AvatarUrl URL of avatar
	AvatarUrl *string `json:"avatar_url,omitempty"`

	// Commands Commands supported by bot
	Commands *[]BotCommand `json:"commands"`

	// Description User description. Can be `null` if user did not fill it out
	Description *string `json:"description"`

	// FullAvatarUrl URL of avatar of a bigger size
	FullAvatarUrl *string `json:"full_avatar_url,omitempty"`

	// IsBot `true` if user is bot
	IsBot bool `json:"is_bot"`

	// LastActivityTime Time of last user activity in TamTam (Unix timestamp in milliseconds). Can be outdated if user disabled its "online" status in settings
	LastActivityTime int64 `json:"last_activity_time"`

	// Name Users visible name
	Name string `json:"name"`

	// UserId Users identifier
	UserId int64 `json:"user_id"`

	// Username Unique public user name. Can be `null` if user is not accessible or it is not set
	Username *string `json:"username"`
}

// BotPatch defines model for BotPatch.
type BotPatch struct {
	// Commands Commands supported by bot. Pass empty list if you want to remove commands
	Commands []BotCommand `json:"commands,omitempty"`

	// Description Bot description up to 16k characters long
	Description string `json:"description,omitempty"`

	// Name Visible name of bot
	Name string `json:"name,omitempty"`

	// Photo Request to set bot photo
	Photo *ImagePayload `json:"photo,omitempty"`

	// Username Bot unique identifier. It can be any string 4-64 characters long containing any digit, letter or special symbols: "-" or "_". It **must** starts with a letter
	Username string `json:"username,omitempty"`
}

// // Button defines model for Button.
// type Button struct {
// 	// Text Visible text of button
// 	Text string `json:"text"`
// 	Type string `json:"type"`
// }

type Callback struct {
	// Unix-time when user pressed the button
	Timestamp int64 `json:"timestamp"`

	// Current keyboard identifier
	CallbackId string `json:"callback_id"`

	// Button payload
	Payload string `json:"payload,omitempty"`

	// User pressed the button
	User *User `json:"user,omitempty"`
}

// CallbackAnswer Send this object when your bot wants to react to when a button is pressed
type CallbackAnswer struct {
	// Message Fill this if you want to modify current message
	Message *NewMessageBody `json:"message,omitempty"`

	// Notification Fill this if you just want to send one-time notification to user
	Notification string `json:"notification,omitempty"`
}

// Chat defines model for Chat.
type Chat struct {
	// ChatId Chats identifier
	ChatId int64 `json:"chat_id"`

	// ChatMessageId Identifier of message that contains `chat` button initialized chat
	ChatMessageId *string `json:"chat_message_id"`

	// Description Chat description
	Description *string `json:"description"`

	// DialogWithUser Another user in conversation. For `dialog` type chats only
	DialogWithUser *UserWithPhoto `json:"dialog_with_user"`

	// Icon Icon of chat
	Icon *Image `json:"icon"`

	// IsPublic Is current chat publicly available. Always `false` for dialogs
	IsPublic bool `json:"is_public"`

	// LastEventTime Time of last event occurred in chat
	LastEventTime int64 `json:"last_event_time"`

	// Link Link on chat
	Link *string `json:"link"`

	// MessagesCount Messages count in chat. Only for group chats and channels. **Not available** for dialogs
	MessagesCount *int `json:"messages_count"`

	// OwnerId Identifier of chat owner. Visible only for chat admins
	OwnerId *int64 `json:"owner_id"`

	// Participants Participants in chat with time of last activity. Can be *null* when you request list of chats. Visible for chat admins only
	Participants *map[string]int64 `json:"participants"`

	// ParticipantsCount Number of people in chat. Always 2 for `dialog` chat type
	ParticipantsCount int32 `json:"participants_count"`

	// PinnedMessage Pinned message in chat or channel. Returned only when single chat is requested
	PinnedMessage *Message `json:"pinned_message"`

	// Status Chat status. One of:
	//  - active: bot is active member of chat
	//  - removed: bot was kicked
	//  - left: bot intentionally left chat
	//  - closed: chat was closed
	//  - suspended: bot was stopped by user. *Only for dialogs*
	Status ChatStatus `json:"status"`

	// Title Visible title of chat. Can be null for dialogs
	Title *string `json:"title"`

	// Type Type of chat. One of: dialog, chat, channel
	Type ChatType `json:"type"`
}

// ChatAdminPermission Chat admin permissions
type ChatAdminPermission string

// ChatList defines model for ChatList.
type ChatList struct {
	// Chats List of requested chats
	Chats []Chat `json:"chats"`

	// Marker Reference to the next page of requested chats
	Marker *int64 `json:"marker"`
}

// ChatMember defines model for ChatMember.
type ChatMember struct {
	// AvatarUrl URL of avatar
	AvatarUrl *string `json:"avatar_url,omitempty"`

	// Description User description. Can be `null` if user did not fill it out
	Description *string `json:"description"`

	// FullAvatarUrl URL of avatar of a bigger size
	FullAvatarUrl *string `json:"full_avatar_url,omitempty"`
	IsAdmin       bool    `json:"is_admin"`

	// IsBot `true` if user is bot
	IsBot    bool  `json:"is_bot"`
	IsOwner  bool  `json:"is_owner"`
	JoinTime int64 `json:"join_time"`

	// LastAccessTime User last activity time in chat. Can be outdated for super chats and channels (equals to `join_time`)
	LastAccessTime int64 `json:"last_access_time"`

	// LastActivityTime Time of last user activity in TamTam (Unix timestamp in milliseconds). Can be outdated if user disabled its "online" status in settings
	LastActivityTime int64 `json:"last_activity_time"`

	// Name Users visible name
	Name string `json:"name"`

	// Permissions Permissions in chat if member is admin. `null` otherwise
	Permissions *[]ChatAdminPermission `json:"permissions"`

	// UserId Users identifier
	UserId int64 `json:"user_id"`

	// Username Unique public user name. Can be `null` if user is not accessible or it is not set
	Username *string `json:"username"`
}

// ChatMembersList defines model for ChatMembersList.
type ChatMembersList struct {
	// Marker Pointer to the next data page
	Marker *int64 `json:"marker"`

	// Members Participants in chat with time of last activity. Visible only for chat admins
	Members []ChatMember `json:"members"`
}

// ChatPatch defines model for ChatPatch.
type ChatPatch struct {
	Icon *ImagePayload `json:"icon"`

	// Notify By default, participants will be notified about change with system message in chat/channel
	Notify *bool `json:"notify"`

	// Pin Identifier of message to be pinned in chat. In case you want to remove pin, use [unpin](#operation/unpinMessage) method
	Pin   *string `json:"pin"`
	Title *string `json:"title"`
}

// ChatStatus Chat status for current bot
type ChatStatus = string

// ChatType Type of chat. Dialog (one-on-one), chat or channel
type ChatType = string

// ConstructedMessageBody defines model for ConstructedMessageBody.
type ConstructedMessageBody struct {
	// Attachments Message attachments. See `AttachmentRequest` and it's inheritors for full information
	Attachments []AttachmentRequest `json:"attachments,omitempty"`

	// Format Message text format. If set,
	Format *TextFormat `json:"format"`

	// Markup Text markup
	Markup []MarkupElement `json:"markup,omitempty"`

	// Text Message text
	Text string `json:"text,omitempty"`
}

// ConstructorAnswer Bot's answer on construction request
type ConstructorAnswer struct {
	// AllowUserInput If `true` user can send any input manually. Otherwise, only keyboard will be shown
	AllowUserInput bool `json:"allow_user_input,omitempty"`

	// Data In this property you can store any additional data up to 8KB. We send this data back to bot within the
	// next construction request. It is handy to store here any state of construction session
	Data string `json:"data,omitempty"`

	// Hint Hint to user. Will be shown on top of keyboard
	Hint string `json:"hint,omitempty"`

	// Keyboard Keyboard to show to user in constructor mode
	Keyboard *Keyboard `json:"keyboard"`

	// Messages Array of prepared messages. This messages will be sent as user taps on "Send" button
	Messages []ConstructedMessageBody `json:"messages,omitempty"`

	// Placeholder Text to show over the text field
	Placeholder string `json:"placeholder,omitempty"`
}

// GetPinnedMessageResult defines model for GetPinnedMessageResult.
type GetPinnedMessageResult struct {
	// Message Pinned message. Can be `null` if no message pinned in chat
	Message *Message `json:"message"`
}

// GetSubscriptionsResult List of all WebHook subscriptions
type GetSubscriptionsResult struct {
	// Subscriptions Current subscriptions
	Subscriptions []Subscription `json:"subscriptions"`
}

// Image Generic schema describing image object
type Image struct {
	// Url URL of image
	Url string `json:"url"`
}

// Keyboard Keyboard is two-dimension array of buttons
type Keyboard struct {
	Buttons [][]Button `json:"buttons"`
}

// LinkedMessage defines model for LinkedMessage.
type LinkedMessage struct {
	// ChatId Chat where message has been originally posted. For forwarded messages only
	ChatId  *int64      `json:"chat_id,omitempty"`
	Message MessageBody `json:"message"`

	// Sender User sent this message. Can be `null` if message has been posted on behalf of a channel
	Sender *User `json:"sender,omitempty"`

	// Type Type of linked message
	Type MessageLinkType `json:"type"`
}

// MarkupElement defines model for MarkupElement.
type MarkupElement struct {
	// From Element start index (zero-based) in text
	From int32 `json:"from"`

	// Length Length of the markup element
	Length int32 `json:"length"`

	// Type Type of the markup element.  Can be **strong**,  *emphasized*, ~strikethrough~, ++underline++, `monospaced`, link or user_mention
	Type string `json:"type"`
}

// Message Message in chat
type Message struct {
	// Body Body of created message. Text + attachments. Could be null if message contains only forwarded message
	Body MessageBody `json:"body"`

	// Constructor Bot-constructor created this message
	Constructor *User `json:"constructor"`

	// Link Forwarded or replied message
	Link *LinkedMessage `json:"link"`

	// Recipient Message recipient. Could be user or chat
	Recipient Recipient `json:"recipient"`

	// Sender User who sent this message. Can be `null` if message has been posted on behalf of a channel
	Sender *User `json:"sender,omitempty"`

	// Stat Message statistics. Available only for channels in [GET:/messages](#operation/getMessages) context
	Stat *MessageStat `json:"stat"`

	// Timestamp Unix-time when message was created
	Timestamp int64 `json:"timestamp"`

	// Url Message public URL. Can be `null` for dialogs or non-public chats/channels
	Url string `json:"url,omitempty"`
}

// MessageBody Schema representing body of message
type MessageBody struct {
	// Attachments Message attachments. Could be one of `Attachment` type. See description of this schema
	Attachments []AttachmentRequest `json:"attachments,omitempty"`

	// Markup Message text markup. See [Formatting](#section/About/Text-formatting) section for more info
	Markup []MarkupElement `json:"markup,omitempty"`

	// Mid Unique identifier of message
	Mid string `json:"mid"`

	// Seq Sequence identifier of message in chat
	Seq int64 `json:"seq"`

	// Text Message text
	Text string `json:"text,omitempty"`
}

// MessageLinkType Type of linked message
type MessageLinkType string

// MessageList Paginated list of messages
type MessageList struct {
	// Messages List of messages
	Messages []Message `json:"messages"`
}

// MessageStat Message statistics
type MessageStat struct {
	Views int `json:"views"`
}

// NewMessageBody defines model for NewMessageBody.
type NewMessageBody struct {
	// Attachments Message attachments. See `AttachmentRequest` and it's inheritors for full information
	Attachments []AttachmentRequest `json:"attachments,omitempty"`

	// Format If set, message text will be formated according to given markup
	Format *TextFormat `json:"format"`

	// Link Link to Message
	Link *NewMessageLink `json:"link"`

	// Notify If false, chat participants would not be notified
	Notify bool `json:"notify,omitempty"`

	// Text Message text
	Text string `json:"text,omitempty"`
}

// NewMessageLink defines model for NewMessageLink.
type NewMessageLink struct {
	// Mid Message identifier of original message
	Mid string `json:"mid"`

	// Type Type of message link
	Type MessageLinkType `json:"type"`
}

// PhotoToken defines model for PhotoToken.
type PhotoToken struct {
	// Token Encoded information of uploaded image
	Token string `json:"token"`
}

// PinMessageBody defines model for PinMessageBody.
type PinMessageBody struct {
	// MessageId Identifier of message to be pinned in chat
	MessageId string `json:"message_id"`

	// Notify If `true`, participants will be notified with system message in chat/channel
	Notify bool `json:"notify,omitempty"`
}

// Recipient New message recipient. Could be user or chat
type Recipient struct {
	// ChatId Chat identifier
	ChatId int64 `json:"chat_id,omitempty"`

	// ChatType Chat type
	ChatType ChatType `json:"chat_type"`

	// UserId User identifier, if message was sent to user
	UserId int64 `json:"user_id,omitempty"`
}

// SendMessageResult defines model for SendMessageResult.
type SendMessageResult struct {
	// Message Message in chat
	Message Message `json:"message"`
}

// SenderAction Different actions to send to chat members
type SenderAction string

func (s SenderAction) GetSenderAction() string {
	return string(s)
}

const (
	TypingOn     SenderAction = "typing_on"
	SendingPhoto SenderAction = "sending_photo"
	SendingVideo SenderAction = "sending_video"
	SendingAudio SenderAction = "sending_audio"
	SendingFile  SenderAction = "sending_file"
	MarkSeen     SenderAction = "mark_seen"
)

// SimpleQueryResult Simple response to request
type SimpleQueryResult struct {
	// Message Explanatory message if the result is not successful
	Message *string `json:"message,omitempty"`

	// Success `true` if request was successful. `false` otherwise
	Success bool `json:"success"`
}

// Subscription Schema to describe WebHook subscription
type Subscription struct {
	// Time Unix-time when subscription was created
	Time int64 `json:"time"`

	// UpdateTypes Update types bot subscribed for
	UpdateTypes []string `json:"update_types,omitempty"`

	// Url Webhook URL
	Url     string `json:"url"`
	Version string `json:"version,omitempty"`
}

// SubscriptionRequestBody Request to set up WebHook subscription
type SubscriptionRequestBody struct {
	// UpdateTypes List of update types your bot want to receive. See `Update` object for a complete list of types
	UpdateTypes []UpdateType `json:"update_types,omitempty"`

	// Url URL of HTTP(S)-endpoint of your bot. Must starts with http(s)://
	Url string `json:"url"`

	// Version Version of API. Affects model representation
	Version string `json:"version,omitempty"`
}

// TextFormat Message text format
type TextFormat string

// UpdateList List of all updates in chats your bot participated in
type UpdateList struct {
	// Marker Pointer to the next data page
	Marker int64 `json:"marker,omitempty"`

	// Updates Page of updates
	Updates []Update `json:"updates"`
}

func (p *UpdateList) UnmarshalJSON(b []byte) error {
	v := struct {
		Marker  int64             `json:"marker,omitempty"`
		Updates []json.RawMessage `json:"updates"`
	}{}
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	updates := make([]Update, len(v.Updates))
	for i, r := range v.Updates {
		update, err := unmarshalUpdate(r)
		if err != nil {
			return err
		}
		updates[i] = *update
	}
	p.Updates = updates
	return nil
}

// Update `Update` object represents different types of events that happened in chat.
type Update struct {
	Type                       UpdateType
	MessageCreated             *MessageCreated
	MessageEdited              *MessageEdited
	MessageRemoved             *MessageRemoved
	MessageCallback            *MessageCallback
	BotAdded                   *BotAdded
	BotRemoved                 *BotRemoved
	BotStarted                 *BotStarted
	UserAdded                  *UserAdded
	UserRemoved                *UserRemoved
	ChatTitleChanged           *ChatTitleChanged
	MessageChatCreated         *MessageChatCreated
	MessageConstructed         *MessageConstructed
	MessageConstructionRequest *MessageConstructionRequest
}

func (u *Update) GetUpdateType() UpdateType {
	return u.Type
}

func unmarshalUpdate(r json.RawMessage) (*Update, error) {
	v := struct {
		UpdateType UpdateType `json:"update_type"`
	}{}
	err := json.Unmarshal(r, &v)
	if err != nil {
		return nil, err
	}
	update := new(Update)
	update.Type = v.UpdateType
	switch v.UpdateType {
	case UpdateTypeMessageCreated:
		t := MessageCreated{}
		err := json.Unmarshal(r, &t)
		if err != nil {
			return nil, err
		}
		update.MessageCreated = &t
	case UpdateTypeMessageCallback:
		t := MessageCallback{}
		err := json.Unmarshal(r, &t)
		if err != nil {
			return nil, err
		}
		update.MessageCallback = &t
	case UpdateTypeMessageEdited:
		t := MessageEdited{}
		err := json.Unmarshal(r, &t)
		if err != nil {
			return nil, err
		}
		update.MessageEdited = &t
	case UpdateTypeMessageRemoved:
		t := MessageRemoved{}
		err := json.Unmarshal(r, &t)
		if err != nil {
			return nil, err
		}
		update.MessageRemoved = &t
	case UpdateTypeBotStarted:
		t := BotStarted{}
		err := json.Unmarshal(r, &t)
		if err != nil {
			return nil, err
		}
		update.BotStarted = &t
	case UpdateTypeBotAdded:
		t := BotAdded{}
		err := json.Unmarshal(r, &t)
		if err != nil {
			return nil, err
		}
		update.BotAdded = &t
	case UpdateTypeBotRemoved:
		t := BotRemoved{}
		err := json.Unmarshal(r, &t)
		if err != nil {
			return nil, err
		}
		update.BotRemoved = &t
	case UpdateTypeUserAdded:
		t := UserAdded{}
		err := json.Unmarshal(r, &t)
		if err != nil {
			return nil, err
		}
		update.UserAdded = &t
	case UpdateTypeUserRemoved:
		t := UserRemoved{}
		err := json.Unmarshal(r, &t)
		if err != nil {
			return nil, err
		}
		update.UserRemoved = &t
	case UpdateTypeChatTitleChanged:
		t := ChatTitleChanged{}
		err := json.Unmarshal(r, &t)
		if err != nil {
			return nil, err
		}
		update.ChatTitleChanged = &t
	case UpdateTypeMessageChatCreated:
		t := MessageChatCreated{}
		err := json.Unmarshal(r, &t)
		if err != nil {
			return nil, err
		}
		update.MessageChatCreated = &t
	case UpdateTypeMessageConstructed:
		t := MessageConstructed{}
		err := json.Unmarshal(r, &t)
		if err != nil {
			return nil, err
		}
		update.MessageConstructed = &t
	case UpdateTypeMessageConstructionRequest:
		t := MessageConstructionRequest{}
		err := json.Unmarshal(r, &t)
		if err != nil {
			return nil, err
		}
		update.MessageConstructionRequest = &t
	}
	return update, nil
}

// UploadEndpoint Endpoint you should upload to your binaries
type UploadEndpoint struct {
	// Url URL to upload
	Url string `json:"url"`
}

// UploadType Type of file uploading
type UploadType string

// User defines model for User.
type User struct {
	// IsBot `true` if user is bot
	IsBot bool `json:"is_bot"`

	// LastActivityTime Time of last user activity in TamTam (Unix timestamp in milliseconds). Can be outdated if user disabled its "online" status in settings
	LastActivityTime int64 `json:"last_activity_time"`

	// Name Users visible name
	Name string `json:"name"`

	// UserId Users identifier
	UserId int64 `json:"user_id"`

	// Username Unique public user name. Can be `null` if user is not accessible or it is not set
	Username string `json:"username,omitempty"`
}

// UserIdsList defines model for UserIdsList.
type UserIdsList struct {
	UserIds interface{} `json:"user_ids"`
}

// UserWithPhoto defines model for UserWithPhoto.
type UserWithPhoto struct {
	// AvatarUrl URL of avatar
	AvatarUrl *string `json:"avatar_url,omitempty"`

	// Description User description. Can be `null` if user did not fill it out
	Description *string `json:"description"`

	// FullAvatarUrl URL of avatar of a bigger size
	FullAvatarUrl *string `json:"full_avatar_url,omitempty"`

	// IsBot `true` if user is bot
	IsBot bool `json:"is_bot"`

	// LastActivityTime Time of last user activity in TamTam (Unix timestamp in milliseconds). Can be outdated if user disabled its "online" status in settings
	LastActivityTime int64 `json:"last_activity_time"`

	// Name Users visible name
	Name string `json:"name"`

	// UserId Users identifier
	UserId int64 `json:"user_id"`

	// Username Unique public user name. Can be `null` if user is not accessible or it is not set
	Username *string `json:"username"`
}

// Bigint defines model for bigint.
type Bigint = int64

// GetChatsOpts defines optional parameters for GetChats.
type GetChatsOpts struct {
	// Count Number of chats requested
	Count int32 `form:"count,omitempty" json:"count,omitempty"`

	// Marker Points to next data page. `null` for the first page
	Marker Bigint `form:"marker,omitempty" json:"marker,omitempty"`
}

// GetChatMembersOpts defines optional parameters for GetMembers.
type GetChatMembersOpts struct {
	// UserIds *Since* version [0.1.4](#section/About/Changelog).
	//
	// Comma-separated list of users identifiers to get their membership. When this parameter is passed, both `count` and `marker` are ignored
	UserIds []int64 `json:"user_ids,omitempty"`

	// Marker Marker
	Marker int64 `form:"marker,omitempty" json:"marker,omitempty"`

	// Count Count
	Count int `form:"count,omitempty" json:"count,omitempty"`
}

// DeleteMessageParams defines parameters for DeleteMessage.
type DeleteMessageParams struct {
	// MessageId Deleting message identifier
	MessageId string `form:"message_id" json:"message_id"`
}

// GetMessagesOpts defines parameters for GetMessages.
type GetMessagesOpts struct {
	// ChatId Chat identifier to get messages in chat
	ChatId Bigint `form:"chat_id,omitempty" json:"chat_id,omitempty"`

	// MessageIds Comma-separated list of message ids to get
	MessageIds []string `json:"message_ids,omitempty"`

	// From Start time for requested messages
	From Bigint `form:"from,omitempty" json:"from,omitempty"`

	// To End time for requested messages
	To Bigint `form:"to,omitempty" json:"to,omitempty"`

	// Count Maximum amount of messages in response
	Count int32 `form:"count,omitempty" json:"count,omitempty"`
}

// SendMessageOpts defines optional parameters for SendMessage.
type SendMessageOpts struct {
	// // UserId Fill this parameter if you want to send message to user
	// UserId int64 `form:"user_id,omitempty" json:"user_id,omitempty"`

	// // ChatId Fill this if you send message to chat
	// ChatId int64 `form:"chat_id,omitempty" json:"chat_id,omitempty"`

	// DisableLinkPreview If `false`, server will not generate media preview for links in text
	DisableLinkPreview bool `form:"disable_link_preview,omitempty" json:"disable_link_preview,omitempty"`

	// Attachments Message attachments. Could be one of `Attachment` type. See description of this schema
	Attachments []AttachmentRequest `json:"attachments"`

	// Link to Message
	Link *MessageLink `json:"link,omitempty"`

	// If set, message text will be formated according to given markup
	Format TextFormat `json:"format"`

	// If false, chat participants would not be notified
	Notify bool `json:"notify,omitempty"`
}

type SendMessageBody struct {
	// Message text
	Text string `json:"text,omitempty"`

	// Attachments Message attachments. Could be one of `Attachment` type. See description of this schema
	Attachments []AttachmentRequest `json:"attachments,omitempty"`

	// Link to Message
	Link *MessageLink `json:"link,omitempty"`

	// If set, message text will be formated according to given markup
	Format TextFormat `json:"format,omitempty"`

	// If false, chat participants would not be notified
	Notify bool `json:"notify,omitempty"`
}

// Link to Message
type MessageLink struct {
	// Message identifier of original message
	Mid string `json:"mid"`
	// Type of message link
	Type MessageLinkType `json:"type"`
}

// EditMessageParams defines parameters for EditMessage.
type EditMessageParams struct {
	// MessageId Editing message identifier
	MessageId string `form:"message_id" json:"message_id"`
}

// UnsubscribeParams defines parameters for Unsubscribe.
type UnsubscribeParams struct {
	// Url URL to remove from WebHook subscriptions
	Url string `form:"url" json:"url"`
}

// GetUpdatesOpts defines parameters for GetUpdates.
type GetUpdatesOpts struct {
	// Limit Maximum number of updates to be retrieved
	Limit int `form:"limit,omitempty" json:"limit,omitempty"`

	// Timeout Timeout in seconds for long polling
	Timeout int `form:"timeout,omitempty" json:"timeout,omitempty"`

	// Marker Pass `null` to get updates you didn't get yet
	Marker int64 `form:"marker,omitempty" json:"marker,omitempty"`

	// Types Comma separated list of update types your bot want to receive
	Types []string `json:"types,omitempty"`
}

// GetUploadUrlParams defines parameters for GetUploadUrl.
type GetUploadUrlParams struct {
	// Type Uploaded file type: photo, audio, video, file
	Type UploadType `form:"type" json:"type"`
}

type Input struct {
	InputType string `json:"input"`

	// Pressed button payload
	Payload string `json:"payload,omitempty"`

	// Messages sent by user during construction process.
	// Typically it is single element array but sometimes it can contains multiple messages.
	// Can be empty on initial request when user just opened constructor
	Messages []ConstructedMessageBody `json:"messages,omitempty"`
}
