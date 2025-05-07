// logger/objPool.go

package logger

import (
	"bytes"
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

	SingletonPool = newObjectPool()
)

// Pool interface defines the methods for the object pool.
type Pool interface {
	GetEventModelFromPool(eventName eventModel.EventCode) (eventModel.Event, error)
	PutEventModelToPool(eventName eventModel.EventCode, event eventModel.Event) error
	GetBufferFromPool() *bytes.Buffer
	PutBufferToPool(buf *bytes.Buffer)
}

// objectPool implements the Pool interface.
type objectPool struct {
	bufPool      sync.Pool
	eventPoolMap sync.Map // key: eventModel.EventCode, value: *sync.Pool{eventModel.Event}
}

// newObjectPool initializes a new object pool.
func newObjectPool() Pool {
	newPool := new(objectPool)

	newPool.bufPool = sync.Pool{
		New: func() any {
			return bytes.NewBuffer(make([]byte, 0, 4096))
		},
	}

	newPool.eventPoolMap = sync.Map{}

	for eventCode, newFunc := range modelMapper {
		newPool.eventPoolMap.Store(eventCode, &sync.Pool{
			New: newFunc,
		})
	}

	// create
	return newPool
}

// GetEventModelFromPool retrieves an event model from the pool.
func (op *objectPool) GetEventModelFromPool(eventName eventModel.EventCode) (eventModel.Event, error) {
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

// PutEventModelToPool puts an event model back into the pool.
func (op *objectPool) PutEventModelToPool(eventName eventModel.EventCode, event eventModel.Event) error {
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

// GetBufferFromPool retrieves a bytes.Buffer from the pool.
func (op *objectPool) GetBufferFromPool() *bytes.Buffer {
	return op.bufPool.Get().(*bytes.Buffer)
}

// PutBufferToPool puts a bytes.Buffer back into the pool.
func (op *objectPool) PutBufferToPool(buf *bytes.Buffer) {
	buf.Reset()
	op.bufPool.Put(buf)
}
