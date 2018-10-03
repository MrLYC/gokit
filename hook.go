package gokit

import (
	"container/list"
	"context"
	"errors"
	"sync"
)

//
var (
	ErrHookNotReady = errors.New("hook not ready")
	ErrHookBusy     = errors.New("hook is busy")
)

// HookNextFunc :
type HookNextFunc func(context.Context) error

// HookHandler :
type HookHandler func(ctx context.Context, next HookNextFunc) error

type hookState struct {
	cancelFunc context.CancelFunc
	current    *list.Element
	nextFunc   HookNextFunc
}

// Hook :
type Hook struct {
	*hookState
	lock     sync.Mutex
	handlers *list.List
}

// Hooker :
type Hooker = *Hook

// Activate :
func (h Hooker) Activate(ctx context.Context) error {
	if h.hookState != nil {
		return ErrHookBusy
	}

	h.lock.Lock()
	defer h.lock.Unlock()

	state := &hookState{}
	h.hookState = state
	ctx, state.cancelFunc = context.WithCancel(ctx)
	state.nextFunc = func(c context.Context) error {
		if state.current != nil {
			handler := state.current.Value.(HookHandler)
			state.current = state.current.Next()
			return handler(c, state.nextFunc)
		}
		return nil
	}
	state.current = h.handlers.Front()
	err := state.nextFunc(ctx)

	state.cancelFunc()
	h.hookState = nil

	return err
}

// Add :
func (h Hooker) Add(handler HookHandler) error {
	if h.hookState != nil {
		return ErrHookBusy
	}

	h.lock.Lock()
	defer h.lock.Unlock()

	h.handlers.PushBack(handler)

	return nil
}

// NewHook : make a hook
func NewHook() Hooker {
	return &Hook{
		handlers: list.New(),
	}
}
