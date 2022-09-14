package gottbot

//	type update interface {
//		GetUpdateType() UpdateType
//	}

type MessageCallback struct {
	// Timestamp Unix-time when event has occurred
	Timestamp int64 `json:"timestamp"`

	UpdateType UpdateType `json:"update_type"`

	Callback *Callback `json:"callback"`

	// Original message containing inline keyboard. Can be null in case it had been deleted by the moment a bot got this update
	Message *Message `json:"message,omitempty"`

	// Current user locale in IETF BCP 47 format
	UserLocale string `json:"user_locale,omitempty"`
}

func (*MessageCallback) GetUpdateType() UpdateType {
	return UpdateTypeMessageCallback
}

type MessageCreated struct {
	// Timestamp Unix-time when event has occurred
	Timestamp int64 `json:"timestamp"`

	UpdateType UpdateType `json:"update_type"`

	// Original message containing inline keyboard. Can be null in case it had been deleted by the moment a bot got this update
	Message *Message `json:"message,omitempty"`

	// Current user locale in IETF BCP 47 format
	UserLocale string `json:"user_locale,omitempty"`
}

func (*MessageCreated) GetUpdateType() UpdateType {
	return UpdateTypeMessageCreated
}

type MessageRemoved struct {
	// Timestamp Unix-time when event has occurred
	Timestamp int64 `json:"timestamp"`

	UpdateType UpdateType `json:"update_type"`

	// Identifier of removed message
	MessageId string `json:"message_id"`

	// Chat identifier where message has been deleted
	ChatId int64 `json:"chat_id"`

	// User who deleted this message
	UserId int64 `json:"user_id"`
}

func (*MessageRemoved) GetUpdateType() UpdateType {
	return UpdateTypeMessageRemoved
}

type MessageEdited struct {
	// Timestamp Unix-time when event has occurred
	Timestamp int64 `json:"timestamp"`

	UpdateType UpdateType `json:"update_type"`

	// Message edited
	Message *Message `json:"message,omitempty"`
}

func (*MessageEdited) GetUpdateType() UpdateType {
	return UpdateTypeMessageEdited
}

type BotAdded struct {
	// Timestamp Unix-time when event has occurred
	Timestamp int64 `json:"timestamp"`

	UpdateType UpdateType `json:"update_type"`

	// Chat identifier where bot is added
	ChatId int64 `json:"chat_id"`

	// User who added bot to chat
	User *User `json:"user"`

	// Indicates whether bot has been added to channel or not
	IsChannel bool `json:"is_channel"`
}

func (*BotAdded) GetUpdateType() UpdateType {
	return UpdateTypeBotAdded
}

type BotRemoved struct {
	// Timestamp Unix-time when event has occurred
	Timestamp int64 `json:"timestamp"`

	UpdateType UpdateType `json:"update_type"`

	// Chat identifier where bot is removed
	ChatId int64 `json:"chat_id"`

	// User who removed bot to chat
	User *User `json:"user"`

	// Indicates whether bot has been removed to channel or not
	IsChannel bool `json:"is_channel"`
}

func (*BotRemoved) GetUpdateType() UpdateType {
	return UpdateTypeBotRemoved
}

type UserAdded struct {
	// Timestamp Unix-time when event has occurred
	Timestamp int64 `json:"timestamp"`

	UpdateType UpdateType `json:"update_type"`

	// Chat identifier where event has occurred
	ChatId int64 `json:"chat_id"`

	// User added to chat
	User *User `json:"user"`

	// User who added user to chat. Can be null in case when user joined chat by link
	InviderId int64 `json:"inviter_id,omitempty"`

	// Indicates whether user has been added to channel or not
	IsChannel bool `json:"is_channel"`
}

func (*UserAdded) GetUpdateType() UpdateType {
	return UpdateTypeUserAdded
}

type UserRemoved struct {
	// Timestamp Unix-time when event has occurred
	Timestamp int64 `json:"timestamp"`

	UpdateType UpdateType `json:"update_type"`

	// Chat identifier where event has occurred
	ChatId int64 `json:"chat_id"`

	// User removed from chat
	User *User `json:"user"`

	// Administrator who removed user from chat. Can be null in case when user left chat
	AdminId int64 `json:"admin_id,omitempty"`

	// Indicates whether user has been removed from a channel or not
	IsChannel bool `json:"is_channel"`
}

func (*UserRemoved) GetUpdateType() UpdateType {
	return UpdateTypeUserRemoved
}

type BotStarted struct {
	// Timestamp Unix-time when event has occurred
	Timestamp int64 `json:"timestamp"`

	UpdateType UpdateType `json:"update_type"`

	// Dialog identifier where event has occurred
	ChatId int64 `json:"chat_id"`

	// User pressed the 'Start' button
	User *User `json:"user"`

	// Additional data from deep-link passed on bot startup
	Payload string `json:"payload,omitempty"`

	// Current user locale in IETF BCP 47 format
	UserLocale string `json:"user_locale,omitempty"`
}

func (*BotStarted) GetUpdateType() UpdateType {
	return UpdateTypeBotStarted
}

type ChatTitleChanged struct {
	// Timestamp Unix-time when event has occurred
	Timestamp int64 `json:"timestamp"`

	UpdateType UpdateType `json:"update_type"`

	// Dialog identifier where event has occurred
	ChatId int64 `json:"chat_id"`

	// User pressed the 'Start' button
	User *User `json:"user"`

	// New Title
	Title string `json:"title"`
}

func (*ChatTitleChanged) GetUpdateType() UpdateType {
	return UpdateTypeChatTitleChanged
}

type MessageConstructionRequest struct {
	// Timestamp Unix-time when event has occurred
	Timestamp int64 `json:"timestamp"`

	UpdateType UpdateType `json:"update_type"`

	// User pressed the 'Start' button
	User *User `json:"user"`

	// Constructor session identifier
	SessionId string `json:"session_id"`

	// data received from previous ConstructorAnswer
	Data string `json:"data,omitempty"`

	// User's input. It can be message (text/attachments) or simple button's callback
	Input *Input `json:"input"`

	// Current user locale in IETF BCP 47 format
	UserLocale string `json:"user_locale,omitempty"`
}

func (*MessageConstructionRequest) GetUpdateType() UpdateType {
	return UpdateTypeMessageConstructionRequest
}

type MessageConstructed struct {
	// Timestamp Unix-time when event has occurred
	Timestamp int64 `json:"timestamp"`

	UpdateType UpdateType `json:"update_type"`

	// User message
	Message *Message `json:"message"`

	// Constructor session identifier
	SessionId string `json:"session_id"`
}

func (*MessageConstructed) GetUpdateType() UpdateType {
	return UpdateTypeMessageConstructed
}

type MessageChatCreated struct {
	// Timestamp Unix-time when event has occurred
	Timestamp int64 `json:"timestamp"`

	UpdateType UpdateType `json:"update_type"`

	// Created chat
	Chat *Chat `json:"chat"`

	// Message identifier where the button has been clicked
	MessageId string `json:"message_id"`

	// Payload from chat button
	StartPayload string `json:"start_payload,omitempty"`
}

func (*MessageChatCreated) GetUpdateType() UpdateType {
	return UpdateTypeMessageChatCreated
}
