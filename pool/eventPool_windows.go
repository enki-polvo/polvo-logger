//go:build windows
// +build windows

// pool/eventPool.go

package eventPool

import (
	"errors"
	"fmt"
	"sync"

	model "github.com/enki-polvo/polvo-logger/model"
	eventModel "github.com/enki-polvo/polvo-logger/model/event"
)

const (
	ErrEventNotFound              = "event '%s' not found in pool"
	ErrGetEventFromPoolFailed     = "failed to get event from pool"
	ErrInvalidTypeAssertionInPool = "invalid type assertion for event pool"
)

var (
	modelMapper = map[model.EventCode]func() any{
		model.PROCESS_CREATION: func() any {
			obj := &model.CommonModel{}
			obj.CommonHeader.EventCode = model.PROCESS_CREATION
			obj.CommonHeader.EventName = model.PROCESS_CREATION.String()
			obj.Metadata = &eventModel.ProcessCreationMetadata{}
			return obj
		},
		model.NETWORK_CONNECTION: func() any {
			obj := &model.CommonModel{}
			obj.CommonHeader.EventCode = model.NETWORK_CONNECTION
			obj.CommonHeader.EventName = model.NETWORK_CONNECTION.String()
			obj.Metadata = &eventModel.NetworkConnectionMetadata{}
			return obj
		},
		model.DRIVER_LOAD: func() any {
			obj := &model.CommonModel{}
			obj.CommonHeader.EventCode = model.DRIVER_LOAD
			obj.CommonHeader.EventName = model.DRIVER_LOAD.String()
			obj.Metadata = &eventModel.DriverLoadMetadata{}
			return obj
		},
		model.IMAGE_LOAD: func() any {
			obj := &model.CommonModel{}
			obj.CommonHeader.EventCode = model.IMAGE_LOAD
			obj.CommonHeader.EventName = model.IMAGE_LOAD.String()
			obj.Metadata = &eventModel.ImageLoadMetadata{}
			return obj
		},
		model.PROCESS_ACCESS: func() any {
			obj := &model.CommonModel{}
			obj.CommonHeader.EventCode = model.PROCESS_ACCESS
			obj.CommonHeader.EventName = model.PROCESS_ACCESS.String()
			obj.Metadata = &eventModel.ProcessAccessMetadata{}
			return obj
		},
		model.FILE_EVENT: func() any {
			obj := &model.CommonModel{}
			obj.CommonHeader.EventCode = model.FILE_EVENT
			obj.CommonHeader.EventName = model.FILE_EVENT.String()
			obj.Metadata = &eventModel.FileEventMetadata{}
			return obj
		},
		model.REGISTRY_ADD: func() any {
			obj := &model.CommonModel{}
			obj.CommonHeader.EventCode = model.REGISTRY_ADD
			obj.CommonHeader.EventName = model.REGISTRY_ADD.String()
			obj.Metadata = &eventModel.RegistryAddMetadata{}
			return obj
		},
		model.REGISTRY_DELETE: func() any {
			obj := &model.CommonModel{}
			obj.CommonHeader.EventCode = model.REGISTRY_DELETE
			obj.CommonHeader.EventName = model.REGISTRY_DELETE.String()
			obj.Metadata = &eventModel.RegistryDeleteMetadata{}
			return obj
		},
		model.REGISTRY_SET: func() any {
			obj := &model.CommonModel{}
			obj.CommonHeader.EventCode = model.REGISTRY_SET
			obj.CommonHeader.EventName = model.REGISTRY_SET.String()
			obj.Metadata = &eventModel.RegistrySetMetadata{}
			return obj
		},
		model.REGISTRY_RENAME: func() any {
			obj := &model.CommonModel{}
			obj.CommonHeader.EventCode = model.REGISTRY_RENAME
			obj.CommonHeader.EventName = model.REGISTRY_RENAME.String()
			obj.Metadata = &eventModel.RegistryEventMetadata{}
			return obj
		},
		model.CREATE_STREAM_HASH: func() any {
			obj := &model.CommonModel{}
			obj.CommonHeader.EventCode = model.CREATE_STREAM_HASH
			obj.CommonHeader.EventName = model.CREATE_STREAM_HASH.String()
			obj.Metadata = &eventModel.CreateStreamHashMetadata{}
			return obj
		},
		model.DNS_QUERY: func() any {
			obj := &model.CommonModel{}
			obj.CommonHeader.EventCode = model.DNS_QUERY
			obj.CommonHeader.EventName = model.DNS_QUERY.String()
			obj.Metadata = &eventModel.DnsQueryMetadata{}
			return obj
		},
		model.FILE_DELETE: func() any {
			obj := &model.CommonModel{}
			obj.CommonHeader.EventCode = model.FILE_DELETE
			obj.CommonHeader.EventName = model.FILE_DELETE.String()
			obj.Metadata = &eventModel.FileDeleteMetadata{}
			return obj
		},
	}
)

// Pool interface defines the methods for the object pool.
type Pool interface {
	Allocate(eventName model.EventCode) (*model.CommonModel, error)
	Free(event eventModel.Event) error
}

// eventPool implements the Pool interface.
type eventPool struct {
	eventPoolMap sync.Map // key: eventModel.EventCode, value: *sync.Pool{eventModel.Event}
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
func (op *eventPool) Allocate(eventName model.EventCode) (*model.CommonModel, error) {
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
		return nil, errors.New(ErrInvalidTypeAssertionInPool)
	}

	value = eventPool.Get()
	if value == nil {
		return nil, errors.New(ErrGetEventFromPoolFailed)
	}

	event, ok := value.(*model.CommonModel)
	if !ok {
		return nil, errors.New(ErrInvalidTypeAssertionInPool)
	}
	return event, nil
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
	eventName = event.(*model.CommonModel).EventCode
	// check if event pool exists
	value, isExists = op.eventPoolMap.Load(eventName)
	if !isExists {
		return fmt.Errorf(ErrEventNotFound, eventName.String())
	}

	// put event to pool
	eventPool, ok := value.(*sync.Pool)
	if !ok {
		return errors.New(ErrInvalidTypeAssertionInPool)
	}
	eventPool.Put(event)
	return nil
}
