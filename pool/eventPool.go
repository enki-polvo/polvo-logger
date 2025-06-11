// pool/eventPool.go

package eventPool

import (
	"fmt"
	"sync"

	model "github.com/enki-polvo/polvo-logger/model"
	eventModel "github.com/enki-polvo/polvo-logger/model/event"
	stateConstants "github.com/enki-polvo/polvo-logger/model/state"
)

const (
	ErrEventNotFound              = "Event '%s' not found in pool"
	ErrGetEventFromPoolFailed     = "failed to get event from pool"
	ErrInvalidTypeAssertionInPool = "invalid type assertion for event pool"
)

var (
	modelMapper = map[model.EventCode]func() any{
		model.PROC_CREATE: func() any {
			obj := &model.CommonModel{}
			obj.CommonHeader.EventCode = model.PROC_CREATE
			obj.CommonHeader.EventName = model.PROC_CREATE.String()
			obj.Metadata = &eventModel.ProcessCreateMetadata{}
			return obj
		},
		model.PROC_TERMINATE: func() any {
			obj := &model.CommonModel{}
			obj.CommonHeader.EventCode = model.PROC_TERMINATE
			obj.CommonHeader.EventName = model.PROC_TERMINATE.String()
			obj.Metadata = &eventModel.ProcessTerminateMetadata{}
			return obj
		},
		model.PROC_BASH_READLINE: func() any {
			obj := &model.CommonModel{}
			obj.CommonHeader.EventCode = model.PROC_BASH_READLINE
			obj.CommonHeader.EventName = model.PROC_BASH_READLINE.String()
			obj.Metadata = &eventModel.BashReadlineMetadata{}
			return obj
		},
		model.PROC_SERVICE: func() any {
			obj := &model.CommonModel{}
			obj.CommonHeader.EventCode = model.PROC_SERVICE
			obj.CommonHeader.EventName = model.PROC_SERVICE.String()
			obj.Metadata = &eventModel.ServiceMetadata{}
			return obj
		},
		model.TCP_EVENT: func() any {
			obj := &model.CommonModel{}
			obj.CommonHeader.EventCode = model.TCP_EVENT
			obj.CommonHeader.EventName = model.TCP_EVENT.String()
			// Initialize the metadata Opcode for TCP events
			obj.Metadata = &eventModel.TcpMetadata{
				Op: stateConstants.TCP_OP_UNSET, // default value
			}
			return obj
		},
		model.FILE_OPEN_EVENT: func() any {
			obj := &model.CommonModel{}
			obj.CommonHeader.EventCode = model.FILE_OPEN_EVENT
			obj.CommonHeader.EventName = model.FILE_OPEN_EVENT.String()
			// Initializes the metadata Opcode for File Open events
			obj.Metadata = &eventModel.FileOpenMetadata{
				FileOpenPurposeOp: stateConstants.FILE_OPEN_TO_UNSET, // default value
			}
			return obj
		},
		model.FILE_RENAME_EVENT: func() any {
			obj := &model.CommonModel{}
			obj.CommonHeader.EventCode = model.FILE_RENAME_EVENT
			obj.CommonHeader.EventName = model.FILE_RENAME_EVENT.String()
			obj.Metadata = &eventModel.FileRenameMetadata{}
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
		return nil, fmt.Errorf(ErrInvalidTypeAssertionInPool)
	}

	value = eventPool.Get()
	if value == nil {
		return nil, fmt.Errorf(ErrGetEventFromPoolFailed)
	}

	event, ok := value.(*model.CommonModel)
	if !ok {
		return nil, fmt.Errorf(ErrInvalidTypeAssertionInPool)
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
		return fmt.Errorf(ErrInvalidTypeAssertionInPool)
	}
	eventPool.Put(event)
	return nil
}
