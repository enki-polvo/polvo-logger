// pool/eventPool.go

package eventPool

import (
	"fmt"
	"sync"

	eventModel "github.com/enki-polvo/polvo-logger/model/event"
)

const (
	ErrEventNotFound              = "Event '%s' not found in pool"
	ErrInvalidTypeAssertionInPool = "invalid type assertion for event pool"
)

var (
	modelMapper = map[eventModel.EventCode]func() any{
		eventModel.PROC_CREATE:        func() any { return &eventModel.ProcessCreateEvent{} },
		eventModel.PROC_TERMINATE:     func() any { return &eventModel.ProcessTerminateEvent{} },
		eventModel.PROC_BASH_READLINE: func() any { return &eventModel.BashReadlineEvent{} },
		eventModel.PROC_SERVICE:       func() any { return &eventModel.ServiceEvent{} },
		eventModel.TCP_CONNECT:        func() any { return &eventModel.TcpConnectEvent{} },
		eventModel.TCP_DISCONNECT:     func() any { return &eventModel.TcpDisconnectEvent{} },
		eventModel.FILE_EVENT:         func() any { return &eventModel.FileEvent{} },
	}
)

// Pool interface defines the methods for the object pool.
type Pool interface {
	Allocate(eventName eventModel.EventCode) (eventModel.Event, error)
	Free(eventName eventModel.EventCode, event eventModel.Event) error
}

// eventPool implements the Pool interface.
type eventPool struct {
	eventPoolMap sync.Map // key: eventModel.EventCode, value: *sync.Pool{eventModel.Event}
	size         uint32
}

// newEventPool initializes a new event pool.
func NewEventPool() Pool {
	newPool := new(eventPool)

	newPool.eventPoolMap = sync.Map{}
	// create a pool for each event type
	for eventCode, newFunc := range modelMapper {
		newPool.eventPoolMap.Store(eventCode, &sync.Pool{
			New: newFunc,
		})
	}
	// create
	return newPool
}

// Allocate retrieves an event model from the pool.
func (op *eventPool) Allocate(eventName eventModel.EventCode) (eventModel.Event, error) {
	var (
		value     any
		eventPool *sync.Pool
		isExists  bool
	)

	// check if event pool exists
	value, isExists = op.eventPoolMap.Load(eventName)
	if !isExists {
		return nil, fmt.Errorf(ErrEventNotFound, eventName.String())
	}

	// get event from pool
	eventPool, ok := value.(*sync.Pool)
	if !ok {
		return nil, fmt.Errorf(ErrInvalidTypeAssertionInPool)
	}
	return eventPool.Get(), nil
}

// Free puts an event model back into the pool.
func (op *eventPool) Free(eventName eventModel.EventCode, event eventModel.Event) error {
	var (
		value     any
		eventPool *sync.Pool
		isExists  bool
	)

	// check if event pool exists
	value, isExists = op.eventPoolMap.Load(eventName)
	if !isExists {
		return fmt.Errorf(ErrEventNotFound, eventName.String())
	}

	// put event to pool
	eventPool, ok := value.(*sync.Pool)
	if !ok {
		return fmt.Errorf(ErrInvalidTypeAssertionInPool)
	}
	eventPool.Put(event)
	return nil
}
