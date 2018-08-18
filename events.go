package event

import "sync"

type EventTable struct {
	events map[string][]func(...interface{})
}

var instance *EventTable
var once sync.Once

func Events() *EventTable {
	once.Do(func() {
		instance = &EventTable{
			events: make(map[string][]func(...interface{})),
		}
	})
	return instance
}

func (e *EventTable) On(event string, do func(...interface{})) {
	e.events[event] = append(e.events[event], do)
}

func (e *EventTable) Emit(event string, args ...interface{}) {
	for _, f := range e.events[event] {
		f(args...)
	}
}
