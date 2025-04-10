package events

import (
	"sync"
	"time"
)

type EventInterface interface {
	GetName() string
	GetDataTime() time.Time
	GetPayload() interface{}
}

type EventHandlerInterface interface {
	Handler(event EventInterface, wg *sync.WaitGroup)
}

type EnventDispatcherInterface interface {
	Register(eventName string, handler EventHandlerInterface) error
	Dispatch(event EventInterface) error
	Remove(eventName string, handler EventHandlerInterface) error
	Has(eventName string, handler EventHandlerInterface) bool
	Clear() error
}
