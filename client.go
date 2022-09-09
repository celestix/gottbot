package gottbot

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	API_URL     = "https://botapi.tamtam.chat"
	GET_TIMEOUT = time.Second * 3
)

type Bot struct {
	token  string
	client *http.Client
	*BotInfo
}

type BotOpts struct {
	Client                   *http.Client
	DisableTokenVerification bool
}

func NewBot(token string, opts *BotOpts) (*Bot, error) {
	if opts == nil {
		opts = new(BotOpts)
	}
	if opts.Client == nil {
		opts.Client = new(http.Client)
	}
	b := Bot{
		token:  token,
		client: opts.Client,
	}
	if !opts.DisableTokenVerification {
		info, err := b.GetInfo()
		if err != nil {
			return nil, err
		}
		b.BotInfo = info
	}
	return &b, nil
}

func (b *Bot) MakeRequest(httpMethod string, method string, params url.Values, body []byte) (io.ReadCloser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), GET_TIMEOUT)
	defer cancel()
	return b.makeRequestWithContext(httpMethod, ctx, method, params, body)
}

func (b *Bot) makeRequestWithContext(httpMethod string, ctx context.Context, method string, params url.Values, body []byte) (io.ReadCloser, error) {
	r, err := http.NewRequest(httpMethod, fmt.Sprintf("%s/%s", API_URL, method), safeBody(body))
	if err != nil {
		return nil, fmt.Errorf("failed to build GET request to %s: %w", method, err)
	}
	params.Add("access_token", b.token)
	r.URL.RawQuery = params.Encode()
	resp, err := b.client.Do(r)
	if err != nil {
		return nil, fmt.Errorf("failed to execute GET request to %s: %w", method, err)
	}
	if resp.StatusCode != 200 {
		defer resp.Body.Close()
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.New("unknown error occured while making the request")
		}
		var tamtamError Error
		if err := json.Unmarshal(bytes, &tamtamError); err != nil {
			return nil, errors.New(string(bytes))
		}
		return nil, &tamtamError
	}
	return resp.Body, nil
}

func safeBody(body []byte) io.Reader {
	if body != nil {
		return bytes.NewBuffer(body)
	}
	return nil
}
