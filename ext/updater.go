package ext

import (
	"fmt"
	"time"

	"github.com/anonyindian/gottbot"
)

type Updater struct {
	Dispatcher Dispatcher
	update     chan *gottbot.Update
}

type UpdaterOpts struct {
	Dispatcher Dispatcher
}

func NewUpdater(opts *UpdaterOpts) *Updater {
	if opts == nil {
		opts = new(UpdaterOpts)
	}
	updateChan := make(chan *gottbot.Update)
	if opts.Dispatcher == nil {
		opts.Dispatcher = NewDispatcher()
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

func (u *Updater) Idle() {
	for {
		time.Sleep(time.Second * 1)
	}
}
