package ext

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/anonyindian/gottbot"
)

type Updater struct {
	Dispatcher Dispatcher
	update     chan *gottbot.Update
}

type UpdaterOpts struct {
	Dispatcher   Dispatcher
	ErrorHandler func(*gottbot.Bot, *gottbot.Update, error)
}

func NewUpdater(opts *UpdaterOpts) *Updater {
	if opts == nil {
		opts = new(UpdaterOpts)
	}
	updateChan := make(chan *gottbot.Update)
	if opts.Dispatcher == nil {
		opts.Dispatcher = NewDispatcher(opts.ErrorHandler)
	}
	return &Updater{
		Dispatcher: opts.Dispatcher,
		update:     updateChan,
	}
}

func (u *Updater) StartPolling(bot *gottbot.Bot, opts *gottbot.GetUpdatesOpts) {
	if opts == nil {
		opts = new(gottbot.GetUpdatesOpts)
	}
	go u.Dispatcher.Run(bot, u.update)
	go func() {
		for {
			updates, err := bot.GetUpdates(opts)
			if err != nil {
				fmt.Println("An error occured while fetching updates:", err.Error())
				continue
			}
			opts.Marker = updates.Marker
			for _, update := range updates.Updates {
				u.update <- &update
			}
		}
	}()
}

type WebhookOpts struct {
	Domain      string
	Port        int
	Path        string
	ReadTimeout time.Duration
}

func (u *Updater) StartWebhook(bot *gottbot.Bot, opts *WebhookOpts) {
	if opts == nil {
		opts = new(WebhookOpts)
	}
	if opts.Domain == "" {
		opts.Domain = "0.0.0.0"
	}
	if opts.Port == 0 {
		opts.Port = 8080
	}
	go u.Dispatcher.Run(bot, u.update)
	mux := http.NewServeMux()
	mux.HandleFunc("/"+opts.Path, func(w http.ResponseWriter, r *http.Request) {
		var update gottbot.Update
		_ = json.NewDecoder(r.Body).Decode(&update)
		u.update <- &update
	})
	server := http.Server{
		Addr:        fmt.Sprintf("%s:%d", opts.Domain, opts.Port),
		Handler:     mux,
		ReadTimeout: opts.ReadTimeout,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			panic("failed to create webhook server: " + err.Error())
		}
	}()
}

func (u *Updater) Idle() {
	for {
		time.Sleep(time.Second * 1)
	}
}
