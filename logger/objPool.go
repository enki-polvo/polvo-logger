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
	SingletonPool = newObjectPool()
)

type Pool interface {
	GetEventModelFromPool(eventName eventModel.EventCode) (eventModel.Event, error)
	PutEventModelToPool(eventName eventModel.EventCode, event eventModel.Event) error
	GetBufferFromPool() *bytes.Buffer
	PutBufferToPool(buf *bytes.Buffer)
}

type objectPool struct {
	bufPool      sync.Pool
	eventPoolMap sync.Map // key: eventModel.EventCode, value: *sync.Pool{eventModel.Event}
}

func newObjectPool() Pool {
	newPool := new(objectPool)

	newPool.bufPool = sync.Pool{
		New: func() any {
			return bytes.NewBuffer(make([]byte, 0, 4096))
		},
	}

	newPool.eventPoolMap = sync.Map{}
	// TODO: add all event pools to the object pool
	// create procCreate event pool
	newPool.eventPoolMap.Store(eventModel.PROC_CREATE, &sync.Pool{
		New: func() any {
			return &eventModel.ProcessCreateEvent{}
		},
	})
	// create procTerminate event pool
	newPool.eventPoolMap.Store(eventModel.PROC_TERMINATE, &sync.Pool{
		New: func() any {
			return &eventModel.ProcessTerminateEvent{}
		},
	})
	// create bashReadline event pool
	newPool.eventPoolMap.Store(eventModel.PROC_BASH_READLINE, &sync.Pool{
		New: func() any {
			return &eventModel.BashReadlineEvent{}
		},
	})
	return newPool
}

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

func (op *objectPool) GetBufferFromPool() *bytes.Buffer {
	return op.bufPool.Get().(*bytes.Buffer)
}

func (op *objectPool) PutBufferToPool(buf *bytes.Buffer) {
	buf.Reset()
	op.bufPool.Put(buf)
}
