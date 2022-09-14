package gottbot

import (
	"encoding/json"
	"fmt"
)

type Payload interface {
	GetPayloadType() string
}

// ImagePayload Request to attach image. All fields are mutually exclusive
type ImagePayload struct {
	// Photo ID
	PhotoId int64 `json:"photo_id,omitempty"`

	// Photos Tokens were obtained after uploading images
	Photos map[string]PhotoToken `json:"photos,omitempty"`

	// Token of any existing attachment
	Token string `json:"token,omitempty"`

	// Url Any external image URL you want to attach
	Url string `json:"url,omitempty"`
}

func (*ImagePayload) GetPayloadType() string {
	return "image"
}

// VideoPayload Request to attach video.
type VideoPayload struct {
	// Url of the video
	Url string `json:"url,omitempty"`

	// Token of any existing attachment
	Token string `json:"token,omitempty"`

	// Video Id
	VideoId int64 `json:"id,omitempty"`

	// Thumbnail of the video
	Thumbnail *Image `json:"thumbnail,omitempty"`

	// Duration of the video
	Duration int64 `json:"duration,omitempty"`

	// Sticker width
	Width int `json:"width,omitempty"`

	// Sticker height
	Height int `json:"height,omitempty"`
}

func (*VideoPayload) GetPayloadType() string {
	return "video"
}

// AudioPayload Request to attach audio.
type AudioPayload struct {
	// Url of the video
	Url string `json:"url,omitempty"`

	// Token of any existing attachment
	Token string `json:"token,omitempty"`

	// Audio Id
	AudioId int64 `json:"id,omitempty"`
}

func (*AudioPayload) GetPayloadType() string {
	return "audio"
}

// FilePayload Request to attach file.
type FilePayload struct {
	// Url of file
	Url string `json:"url,omitempty"`

	// Token of any existing attachment
	Token string `json:"token,omitempty"`

	// File id
	FileId int64 `json:"fileId,omitempty"`

	// File name
	Filename string `json:"filename,omitempty"`

	// File size in bytes
	Size int64 `json:"size,omitempty"`
}

func (*FilePayload) GetPayloadType() string {
	return "file"
}

// ContactPayload Request to attach contact.
type ContactPayload struct {
	// TamTam user info of the contact
	TamInfo *User `json:"tam_info,omitempty"`

	// Full information about contact in VCF format
	VCFInfo string `json:"vcf_info,omitempty"`

	// Contact Name
	Name string `json:"name,omitempty"`

	// Contact identifier if it is registered TamTam user
	ContactId int64 `json:"contact_id,omitempty"`

	// Contact phone in VCF format
	VCFPhone string `json:"vcf_phone,omitempty"`
}

func (*ContactPayload) GetPayloadType() string {
	return "contact"
}

// StickerPayload Request to attach sticker.
type StickerPayload struct {
	// Sticker code
	Code string `json:"code,omitempty"`

	// Url Any external sticker URL you want to attach
	Url string `json:"url,omitempty"`

	// Sticker width
	Width int `json:"width,omitempty"`

	// Sticker height
	Height int `json:"height,omitempty"`
}

func (*StickerPayload) GetPayloadType() string {
	return "sticker"
}

// LocationPayload Request to attach location.
type LocationPayload struct {
	// latitude
	Latitude float64 `json:"latitude,omitempty"`
	// longitude
	Longitude float64 `json:"longitude,omitempty"`
}

func (*LocationPayload) GetPayloadType() string {
	return "location"
}

// SharePayload Request to attach share.
type SharePayload struct {
	// Attachment token
	Token string `json:"token,omitempty"`

	// URL attached to message as media preview
	Url string `json:"url,omitempty"`

	// Link Preview Title
	Title string `json:"title,omitempty"`

	// Link Preview Description
	Description string `json:"description,omitempty"`

	// Link Preview Image Url
	ImageUrl string `json:"image_url,omitempty"`
}

func (*SharePayload) GetPayloadType() string {
	return "share"
}

// ButtonsPayload Request to attach buttons.
type ButtonsPayload struct {
	Buttons [][]Button `json:"buttons"`
}

func (*ButtonsPayload) GetPayloadType() string {
	return "inline_keyboard"
}

func (p *ButtonsPayload) UnmarshalJSON(b []byte) error {
	v := struct {
		Payload struct {
			Buttons [][]json.RawMessage `json:"buttons"`
		} `json:"payload"`
	}{}
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	btnss := make([][]Button, len(v.Payload.Buttons))
	for x, rs := range v.Payload.Buttons {
		btns := make([]Button, len(rs))
		for y, r := range rs {
			btn, err := unmarshalButton(r)
			if err != nil {
				return err
			}
			btns[y] = btn
		}
		btnss[x] = btns
	}
	p.Buttons = btnss
	return nil
}

func unmarshalButton(r json.RawMessage) (Button, error) {
	v := struct {
		Type string `json:"type"`
	}{}
	err := json.Unmarshal(r, &v)
	if err != nil {
		return nil, err
	}
	switch v.Type {
	case "callback":
		t := CallbackButton{}
		err := json.Unmarshal(r, &t)
		if err != nil {
			return nil, err
		}
		return &t, nil
	case "link":
		t := LinkButton{}
		err := json.Unmarshal(r, &t)
		if err != nil {
			return nil, err
		}
		return &t, nil
	case "request_contact":
		t := RequestContactButton{}
		err := json.Unmarshal(r, &t)
		if err != nil {
			return nil, err
		}
		return &t, nil
	case "request_geo_location":
		t := RequestGeoLocationButton{}
		err := json.Unmarshal(r, &t)
		if err != nil {
			return nil, err
		}
		return &t, nil
	case "chat":
		t := ChatButton{}
		err := json.Unmarshal(r, &t)
		if err != nil {
			return nil, err
		}
		return &t, nil
	}
	return nil, fmt.Errorf("unknown button type: %s", v.Type)
}
