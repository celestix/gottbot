package gottbot

import (
	"bytes"
	"io"
	"os"
	"time"
)

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

type InputFile interface {
	string | *FileInfo | *os.File
}

type MediaOpts struct {
	// Text (Caption) for the media
	Text string `json:"text,omitempty"`

	// DisableLinkPreview If `false`, server will not generate media preview for links in text
	DisableLinkPreview bool `form:"disable_link_preview,omitempty" json:"disable_link_preview,omitempty"`

	// Extrat Payload for the request (if any)
	Payload Payload `json:"attachments"`

	// Link to Message
	Link *MessageLink `json:"link,omitempty"`

	// If set, message text will be formated according to given markup
	Format TextFormat `json:"format"`

	// If false, chat participants would not be notified
	Notify bool `json:"notify,omitempty"`
}

func SendFile[input InputFile](bot *Bot, chatId int64, file input, opts *MediaOpts) (*SendMessageResult, error) {
	if opts == nil {
		opts = new(MediaOpts)
	}
	var fileInfo *FileInfo
	switch v := any(file).(type) {
	case string:
		var atts []AttachmentRequest
		if opts.Payload != nil {
			atts = []AttachmentRequest{
				{&FilePayload{
					Url: v,
				}},
				{opts.Payload},
			}
		} else {
			atts = []AttachmentRequest{
				{&FilePayload{
					Url: v,
				}},
			}
		}
		return bot.SendMessage(chatId, opts.Text, &SendMessageOpts{
			opts.DisableLinkPreview,
			atts,
			opts.Link,
			opts.Format,
			opts.Notify,
		})
	case *FileInfo:
		fileInfo = v
	case *os.File:
		fileInfo = &FileInfo{
			Name: v.Name(),
			File: v,
		}
	}

	wait := initUploadWait(fileInfo.File)

	payload, err := bot.Upload(UploadTypeFile, fileInfo)
	if err != nil {
		return nil, err
	}

	// wait for the payload to be processed
	wait.Sleep()

	var atts []AttachmentRequest
	if opts.Payload != nil {
		atts = []AttachmentRequest{
			{payload},
			{opts.Payload},
		}
	} else {
		atts = []AttachmentRequest{
			{opts.Payload},
		}
	}
	return bot.SendMessage(chatId, opts.Text, &SendMessageOpts{
		opts.DisableLinkPreview,
		atts,
		opts.Link,
		opts.Format,
		opts.Notify,
	})
}

func SendVideo[input InputFile](bot *Bot, chatId int64, video input, opts *MediaOpts) (*SendMessageResult, error) {
	if opts == nil {
		opts = new(MediaOpts)
	}
	var fileInfo *FileInfo
	switch v := any(video).(type) {
	case string:
		var atts []AttachmentRequest
		if opts.Payload != nil {
			atts = []AttachmentRequest{
				{&VideoPayload{
					Url: v,
				}},
				{opts.Payload},
			}
		} else {
			atts = []AttachmentRequest{
				{&VideoPayload{
					Url: v,
				}},
			}
		}
		return bot.SendMessage(chatId, opts.Text, &SendMessageOpts{
			opts.DisableLinkPreview,
			atts,
			opts.Link,
			opts.Format,
			opts.Notify,
		})
	case *FileInfo:
		fileInfo = v
	case *os.File:
		fileInfo = &FileInfo{
			Name: v.Name(),
			File: v,
		}
	}

	wait := initUploadWait(fileInfo.File)

	payload, err := bot.Upload(UploadTypeVideo, fileInfo)
	if err != nil {
		return nil, err
	}

	// wait for the payload to be processed
	wait.Sleep()

	var atts []AttachmentRequest
	if opts.Payload != nil {
		atts = []AttachmentRequest{
			{payload},
			{opts.Payload},
		}
	} else {
		atts = []AttachmentRequest{
			{opts.Payload},
		}
	}
	return bot.SendMessage(chatId, opts.Text, &SendMessageOpts{
		opts.DisableLinkPreview,
		atts,
		opts.Link,
		opts.Format,
		opts.Notify,
	})
}

func SendAudio[input InputFile](bot *Bot, chatId int64, audio input, opts *MediaOpts) (*SendMessageResult, error) {
	if opts == nil {
		opts = new(MediaOpts)
	}
	var fileInfo *FileInfo
	switch v := any(audio).(type) {
	case string:
		var atts []AttachmentRequest
		if opts.Payload != nil {
			atts = []AttachmentRequest{
				{&AudioPayload{
					Url: v,
				}},
				{opts.Payload},
			}
		} else {
			atts = []AttachmentRequest{
				{&AudioPayload{
					Url: v,
				}},
			}
		}
		return bot.SendMessage(chatId, opts.Text, &SendMessageOpts{
			opts.DisableLinkPreview,
			atts,
			opts.Link,
			opts.Format,
			opts.Notify,
		})
	case *FileInfo:
		fileInfo = v
	case *os.File:
		fileInfo = &FileInfo{
			Name: v.Name(),
			File: v,
		}
	}

	wait := initUploadWait(fileInfo.File)

	payload, err := bot.Upload(UploadTypeAudio, fileInfo)
	if err != nil {
		return nil, err
	}

	// wait for the payload to be processed
	wait.Sleep()

	var atts []AttachmentRequest
	if opts.Payload != nil {
		atts = []AttachmentRequest{
			{payload},
			{opts.Payload},
		}
	} else {
		atts = []AttachmentRequest{
			{opts.Payload},
		}
	}
	return bot.SendMessage(chatId, opts.Text, &SendMessageOpts{
		opts.DisableLinkPreview,
		atts,
		opts.Link,
		opts.Format,
		opts.Notify,
	})
}

func SendPhoto[input InputFile](bot *Bot, chatId int64, photo input, opts *MediaOpts) (*SendMessageResult, error) {
	if opts == nil {
		opts = new(MediaOpts)
	}
	var fileInfo *FileInfo
	switch v := any(photo).(type) {
	case string:
		var atts []AttachmentRequest
		if opts.Payload != nil {
			atts = []AttachmentRequest{
				{&ImagePayload{
					Url: v,
				}},
				{opts.Payload},
			}
		} else {
			atts = []AttachmentRequest{
				{&ImagePayload{
					Url: v,
				}},
			}
		}
		return bot.SendMessage(chatId, opts.Text, &SendMessageOpts{
			opts.DisableLinkPreview,
			atts,
			opts.Link,
			opts.Format,
			opts.Notify,
		})
	case *FileInfo:
		fileInfo = v
	case *os.File:
		fileInfo = &FileInfo{
			Name: v.Name(),
			File: v,
		}
	}

	wait := initUploadWait(fileInfo.File)

	payload, err := bot.Upload(UploadTypeImage, fileInfo)
	if err != nil {
		return nil, err
	}

	// wait for the payload to be processed
	wait.Sleep()

	var atts []AttachmentRequest
	if opts.Payload != nil {
		atts = []AttachmentRequest{
			{payload},
			{opts.Payload},
		}
	} else {
		atts = []AttachmentRequest{
			{opts.Payload},
		}
	}
	return bot.SendMessage(chatId, opts.Text, &SendMessageOpts{
		opts.DisableLinkPreview,
		atts,
		opts.Link,
		opts.Format,
		opts.Notify,
	})
}

type uploadWait struct {
	sleepDur time.Duration
}

func initUploadWait(reader io.Reader) *uploadWait {
	buf := new(bytes.Buffer)
	n, _ := buf.ReadFrom(reader)
	return &uploadWait{time.Duration(n / 100000)}
}

func (u *uploadWait) Sleep() {
	time.Sleep(u.sleepDur)
}
