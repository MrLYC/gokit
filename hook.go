package gokit

import (
	"container/list"
	"context"
	"errors"
	"sync"
)

//
var (
	ErrHookNotActivate = errors.New("hook not activate")
)

// HookHandler :
type HookHandler interface {
	Run(Hooker) error
}

// Hook :
type Hook struct {
	context.Context
	cancelFunc     context.CancelFunc
	currentHandler *list.Element
	lock           sync.Mutex
	handlers       *list.List
}

// Hooker :
type Hooker = *Hook

// Activate :
func (h Hooker) Activate(ctx context.Context) error {
	h.lock.Lock()
	h.Context, h.cancelFunc = context.WithCancel(ctx)
	defer func() {
		h.Context = nil
		h.cancelFunc = nil
		h.currentHandler = nil
		h.lock.Unlock()
	}()

	for h.currentHandler = h.handlers.Front(); h.currentHandler != h.handlers.Back(); {
		handler := h.currentHandler.Value.(HookHandler)
		err := handler.Run(h)
		if err != nil {
			return err
		}
	}

	return nil
}

// Next :
func (h Hooker) Next() error {
	if h.currentHandler == nil {
		return ErrHookNotActivate
	}
	h.currentHandler.Next()
	return nil
}

// Add :
func (h Hooker) Add(handler HookHandler) {
	h.lock.Lock()

	h.handlers.PushBack(handler)

	h.lock.Unlock()
}

// Remove :
func (h Hooker) Remove(handler HookHandler) {
	h.lock.Lock()

	var el *list.Element
	for el = h.handlers.Front(); el != h.handlers.Back(); el.Next() {
		if el.Value.(HookHandler) == handler {
			break
		}
	}

	if el != nil {
		h.handlers.Remove(el)
	}

	h.lock.Unlock()
}
