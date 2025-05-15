// pool/eventPool.go

package eventPool

import (
	"fmt"
	"sync"

	model "github.com/enki-polvo/polvo-logger/model"
	eventModel "github.com/enki-polvo/polvo-logger/model/event"
)

const (
	ErrEventNotFound              = "Event '%s' not found in pool"
	ErrInvalidTypeAssertionInPool = "invalid type assertion for event pool"
)

var (
	modelMapper = map[model.EventCode]func() any{
		model.PROC_CREATE: func() any {
			obj := &eventModel.ProcessCreateEvent{}
			obj.CommonHeader.EventCode = model.PROC_CREATE
			obj.CommonHeader.EventName = model.PROC_CREATE.String()
			return obj
		},
		model.PROC_TERMINATE: func() any {
			obj := &eventModel.ProcessTerminateEvent{}
			obj.CommonHeader.EventCode = model.PROC_TERMINATE
			obj.CommonHeader.EventName = model.PROC_TERMINATE.String()
			return obj
		},
		model.PROC_BASH_READLINE: func() any {
			obj := &eventModel.BashReadlineEvent{}
			obj.CommonHeader.EventCode = model.PROC_BASH_READLINE
			obj.CommonHeader.EventName = model.PROC_BASH_READLINE.String()
			return obj
		},
		model.PROC_SERVICE: func() any {
			obj := &eventModel.ServiceEvent{}
			obj.CommonHeader.EventCode = model.PROC_SERVICE
			obj.CommonHeader.EventName = model.PROC_SERVICE.String()
			return obj
		},
		model.TCP_CONNECT: func() any {
			obj := &eventModel.TcpConnectEvent{}
			obj.CommonHeader.EventCode = model.TCP_CONNECT
			obj.CommonHeader.EventName = model.TCP_CONNECT.String()
			return obj
		},
		model.TCP_DISCONNECT: func() any {
			obj := &eventModel.TcpDisconnectEvent{}
			obj.CommonHeader.EventCode = model.TCP_DISCONNECT
			obj.CommonHeader.EventName = model.TCP_DISCONNECT.String()
			return obj
		},
		model.FILE_EVENT: func() any {
			obj := &eventModel.FileEvent{}
			obj.CommonHeader.EventCode = model.FILE_EVENT
			obj.CommonHeader.EventName = model.FILE_EVENT.String()
			return obj
		},
	}
)

// Pool interface defines the methods for the object pool.
type Pool interface {
	Allocate(eventName model.EventCode) (eventModel.Event, error)
	Free(event eventModel.Event) error
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
func (op *eventPool) Allocate(eventName model.EventCode) (eventModel.Event, error) {
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
func (op *eventPool) Free(event eventModel.Event) error {
	var (
		value     any
		eventPool *sync.Pool
		isExists  bool
		eventName model.EventCode
	)

	// get event name from event
	eventName = event.(model.CommonModel).EventCode
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
