package ext

import (
	"errors"
	"fmt"
	"sort"

	"github.com/anonyindian/gottbot"
)

var (
	EndGroups        = errors.New("end")
	ContinueGroup    = errors.New("continue")
	SkipCurrentGroup = errors.New("skip")
)

type Dispatcher interface {
	// setUpdateChan(update chan gottbot.Update)
	AddHandlerToGroup(group int, handler Handler) HandlerID
	AddHandler(handler Handler) HandlerID
	RemoveHandler(id HandlerID) bool
	Run(bot *gottbot.Bot, updateChan chan *gottbot.Update)
}

type GeneralDispatcher struct {
	handlerGroups []int
	handlerMap    map[int][]Handler
	ErrorHandler  func(*gottbot.Bot, *gottbot.Update, error)
}

func NewDispatcher(errorHandler func(*gottbot.Bot, *gottbot.Update, error)) *GeneralDispatcher {
	return &GeneralDispatcher{
		handlerGroups: make([]int, 0),
		handlerMap:    make(map[int][]Handler),
		ErrorHandler:  errorHandler,
	}
}

// func (g *GeneralDispatcher) setUpdateChan(update chan gottbot.Update) {
// 	g.update = update
// }

func (g *GeneralDispatcher) Run(bot *gottbot.Bot, updateChan chan *gottbot.Update) {
	for update := range updateChan {
		g.processUpdate(bot, update)
	}
}

func (g *GeneralDispatcher) processUpdate(bot *gottbot.Bot, update *gottbot.Update) {
	for _, handlers := range g.handlerMap {
		for _, handler := range handlers {
			if !handler.CheckUpdate(update) {
				continue
			}
			err := handler.HandleUpdate(bot, update)
			if err == nil || errors.Is(err, SkipCurrentGroup) {
				break
			}
			switch {
			case errors.Is(err, EndGroups):
				return
			case errors.Is(err, ContinueGroup):
				continue
			case g.ErrorHandler != nil:
				g.ErrorHandler(bot, update, err)
			default:
				fmt.Println("An error occured:", err.Error())
			}
		}
	}
}

func (g *GeneralDispatcher) AddHandlerToGroup(group int, handler Handler) HandlerID {
	handlers, ok := g.handlerMap[group]
	if !ok {
		handlers = make([]Handler, 0)
		g.handlerGroups = append(g.handlerGroups, group)
		sort.Ints(g.handlerGroups)
	}
	handlers = append(handlers, handler)
	g.handlerMap[group] = handlers
	return handler.GetHandlerID()
}

func (g *GeneralDispatcher) AddHandler(handler Handler) HandlerID {
	return g.AddHandlerToGroup(0, handler)
}

func (g *GeneralDispatcher) RemoveGroup(group int) {
	delete(g.handlerMap, group)
}

func (g *GeneralDispatcher) RemoveHandler(id HandlerID) bool {
	for group, handlers := range g.handlerMap {
		for i, handler := range handlers {
			if handler.GetHandlerID() != id {
				continue
			}
			handlers[i] = handlers[len(handlers)-1]
			handlers[len(handlers)-1] = nil
			handlers = handlers[:len(handlers)-1]
			g.handlerMap[group] = handlers
			return true
		}
	}
	return false
}
