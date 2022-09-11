package gottbot

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Button interface {
	GetButtonText() string
	GetButtonType() string
}

type CallbackButton struct {
	// Visible text of button
	Text string `json:"text"`

	// Button payload
	Payload string `json:"payload"`

	// Default: "default"
	// Enum: "positive" "negative" "default"
	// Intent of button. Affects clients representation
	Intent string `json:"intent,omitempty"`
}

func (b *CallbackButton) GetButtonText() string {
	return b.Text
}

func (b *CallbackButton) GetButtonType() string {
	return "callback"
}

func (b CallbackButton) MarshalJSON() ([]byte, error) {
	type temp CallbackButton
	v := struct {
		Type string `json:"type"`
		temp
	}{
		Type: b.GetButtonType(),
		temp: temp(b),
	}
	return json.Marshal(v)
}

type LinkButton struct {
	// Visible text of button
	Text string `json:"text"`

	// Button url
	Url string `json:"url"`
}

func (b *LinkButton) GetButtonText() string {
	return b.Text
}

func (b *LinkButton) GetButtonType() string {
	return "link"
}

func (b LinkButton) MarshalJSON() ([]byte, error) {
	type temp LinkButton
	v := struct {
		Type string `json:"type"`
		temp
	}{
		Type: b.GetButtonType(),
		temp: temp(b),
	}
	return json.Marshal(v)
}

type RequestContactButton struct {
	// Visible text of button
	Text string `json:"text"`
}

func (b *RequestContactButton) GetButtonText() string {
	return b.Text
}

func (b *RequestContactButton) GetButtonType() string {
	return "request_contact"
}

func (b RequestContactButton) MarshalJSON() ([]byte, error) {
	type temp RequestContactButton
	v := struct {
		Type string `json:"type"`
		temp
	}{
		Type: b.GetButtonType(),
		temp: temp(b),
	}
	return json.Marshal(v)
}

type RequestGeoLocationButton struct {
	// Visible text of button
	Text string `json:"text"`

	// If true, sends location without asking user's confirmation
	Quick bool `json:"quick,omitempty"`
}

func (b *RequestGeoLocationButton) GetButtonText() string {
	return b.Text
}

func (b *RequestGeoLocationButton) GetButtonType() string {
	return "request_geo_location"
}

func (b RequestGeoLocationButton) MarshalJSON() ([]byte, error) {
	type temp RequestGeoLocationButton
	v := struct {
		Type string `json:"type"`
		temp
	}{
		Type: b.GetButtonType(),
		temp: temp(b),
	}
	return json.Marshal(v)
}

type ChatButton struct {
	// Visible text of button
	Text string `json:"text"`

	// Title of chat to be created
	ChatTitle string `json:"chat_title"`

	// Chat description
	ChatDescription string `json:"chat_description,omitempty"`

	// Start payload will be sent to bot as soon as chat created
	StartPayload string `json:"start_payload,omitempty"`

	// Unique button identifier across all chat buttons in keyboard.
	// If uuid changed, new chat will be created on the next click.
	// Server will generate it at the time when button initially posted.
	// Reuse it when you edit the message.'
	UUID string `json:"uuid,omitempty"`
}

func (b *ChatButton) GetButtonText() string {
	return b.Text
}

func (b *ChatButton) GetButtonType() string {
	return "chat"
}

func (b ChatButton) MarshalJSON() ([]byte, error) {
	if b.UUID == "" {
		b.UUID = uuid.New().String()
	}
	type temp ChatButton
	v := struct {
		Type string `json:"type"`
		temp
	}{
		Type: b.GetButtonType(),
		temp: temp(b),
	}
	return json.Marshal(v)
}
