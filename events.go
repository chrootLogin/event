package event

import "sync"

type Events struct {
	events map[string][]func(...interface{})
}

var instance *Events
var once sync.Once

func Events() *Events {
	once.Do(func() {
		instance = &Events{
			events: make(map[string][]func(...interface{})),
		}
	})
	return instance
}

func (e *Events) On(event string, do func(...interface{})) {
	e.events[event] = append(e.events[event], do)
}

func (e *Events) Emit(event string, args ...interface{}) {
	for _, f := range e.events[event] {
		f(args...)
	}
}
