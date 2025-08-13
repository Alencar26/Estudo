package events

import (
	"errors"
	"sync"
)

var ErrHandlerAlreadyRegister = errors.New("handler already registered")

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (ed *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {

	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return ErrHandlerAlreadyRegister
			}
		}
	}
	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}

func (ed *EventDispatcher) Clear() {
	ed.handlers = make(map[string][]EventHandlerInterface)
}

func (ed *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	if _, has := ed.handlers[eventName]; has {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return true
			}
		}
	}
	return false
}

func (ed *EventDispatcher) Dispatch(event EventInterface) error {
	if handlers, ok := ed.handlers[event.GetName()]; ok {
		wg := &sync.WaitGroup{}
		for _, handler := range handlers {
			wg.Add(1)
			go handler.Handle(event, wg)
		}
		wg.Wait()
	}
	return nil
}

func (ed *EventDispatcher) Remove(evetName string, handler EventHandlerInterface) error {
	if _, ok := ed.handlers[evetName]; ok {
		for i, h := range ed.handlers[evetName] {
			if h == handler {
				ed.handlers[evetName] = append(ed.handlers[evetName][:i], ed.handlers[evetName][i+1:]...)
				return nil
			}
		}
	}
	return nil
}
