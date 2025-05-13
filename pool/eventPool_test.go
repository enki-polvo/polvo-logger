package eventPool_test

import (
	"sync"
	"testing"
	"time"

	"fmt"

	eventModel "github.com/enki-polvo/polvo-logger/model/event"
	eventPool "github.com/enki-polvo/polvo-logger/pool"
)

// Test Initialization of the event pool
func TestCreateNewPool(t *testing.T) {
	pool := eventPool.NewEventPool()
	if pool == nil {
		t.Fatal("Failed to create new event pool")
	}
}

// Test Allocation of a valid event
// Test that a valid event can be allocated from the pool
func TestAllocateEvent(t *testing.T) {
	pool := eventPool.NewEventPool()
	eventCode := eventModel.PROC_CREATE

	event, err := pool.Allocate(eventCode)
	if err != nil {
		t.Fatalf("Failed to allocate event: %v", err)
	}

	if event == nil {
		t.Fatal("Allocated event is nil")
	}

	if _, ok := event.(*eventModel.ProcessCreateEvent); !ok {
		t.Fatalf("Allocated event is not of type ProcessCreateEvent")
	}
	// Free the event back to the pool
	err = pool.Free(eventCode, event)
	if err != nil {
		t.Fatalf("Failed to free event: %v", err)
	}
}

// Test Allocation of an invalid event
// Test that an invalid event code returns an error
func TestAllocateInvalidEvent(t *testing.T) {
	pool := eventPool.NewEventPool()
	invalidEventCode := eventModel.EventCode(999) // Assuming 999 is not a valid event code

	event, err := pool.Allocate(invalidEventCode)
	if err == nil {
		t.Fatal("Expected error when allocating invalid event, but got none")
	}

	if event != nil {
		t.Fatal("Allocated event should be nil for invalid event code")
	}
}

// Test Freeing of a valid event
// Test that a valid event can be freed back to the pool
func TestFreeEvent(t *testing.T) {
	pool := eventPool.NewEventPool()
	eventCode := eventModel.PROC_CREATE

	event, err := pool.Allocate(eventCode)
	if err != nil {
		t.Fatalf("Failed to allocate event: %v", err)
	}

	if event == nil {
		t.Fatal("Allocated event is nil")
	}

	err = pool.Free(eventCode, event)
	if err != nil {
		t.Fatalf("Failed to free event: %v", err)
	}
}

// Test Freeing of an invalid event
// Test that freeing an invalid event code returns an error
func TestFreeInvalidEvent(t *testing.T) {
	pool := eventPool.NewEventPool()
	invalidEventCode := eventModel.EventCode(999) // Assuming 999 is not a valid event code
	event := &eventModel.ProcessCreateEvent{}     // Create a dummy event

	err := pool.Free(invalidEventCode, event)
	if err == nil {
		t.Fatal("Expected error when freeing invalid event, but got none")
	}
}

// Test Stress Test for Allocating and Freeing Events
// This test checks the performance of allocating and freeing events in a loop
func TestStressTestAllocateFree(t *testing.T) {
	pool := eventPool.NewEventPool()
	eventCode := eventModel.PROC_CREATE

	for i := 0; i < 1000; i++ {
		event, err := pool.Allocate(eventCode)
		if err != nil {
			t.Fatalf("Failed to allocate event: %v", err)
		}

		if event == nil {
			t.Fatal("Allocated event is nil")
		}

		err = pool.Free(eventCode, event)
		if err != nil {
			t.Fatalf("Failed to free event: %v", err)
		}
	}
}

// Test Stress Test for Allocating Invalid Events
// This test checks the performance of allocating invalid events in a loop
// and ensures that it returns an error
func TestStressTestAllocateInvalidEvent(t *testing.T) {
	pool := eventPool.NewEventPool()
	invalidEventCode := eventModel.EventCode(999) // Assuming 999 is not a valid event code

	for i := 0; i < 1000; i++ {
		event, err := pool.Allocate(invalidEventCode)
		if err == nil {
			t.Fatal("Expected error when allocating invalid event, but got none")
		}

		if event != nil {
			t.Fatal("Allocated event should be nil for invalid event code")
		}
	}
}

// Test Stress Test for Freeing Invalid Events in Multiple Goroutines
// This test checks the performance of freeing invalid events in multiple goroutines
// and ensures that it returns an error
func TestStressTestMultipleGoroutines(t *testing.T) {
	pool := eventPool.NewEventPool()
	eventCode := eventModel.PROC_CREATE

	var wg sync.WaitGroup
	var errChan chan error

	errChan = make(chan error, 0)
	defer close(errChan)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				event, err := pool.Allocate(eventCode)
				if err != nil {
					errChan <- fmt.Errorf("failed to allocate event: %v", err)
					return
				}

				if event == nil {
					errChan <- fmt.Errorf("allocated event is nil")
					return
				}

				err = pool.Free(eventCode, event)
				if err != nil {
					errChan <- fmt.Errorf("failed to free event: %v", err)
					return
				}
			}
		}()
	}
	wg.Wait()
	if len(errChan) > 0 {
		for err := range errChan {
			t.Log(err)
		}
		t.Fatal("Errors occurred during stress test")
	}
}

// Test Stress Test for Invalid Events in Multiple Goroutines
// This test checks the performance of allocating and freeing invalid events
// in multiple goroutines and ensures that it returns an error
func TestStressTestInvalidInMultipleGoroutines(t *testing.T) {
	pool := eventPool.NewEventPool()
	invalidEventCode := eventModel.EventCode(999) // Assuming 999 is not a valid event code

	var wg sync.WaitGroup
	var errChan chan error
	errChan = make(chan error, 0)
	defer close(errChan)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				event, err := pool.Allocate(invalidEventCode)
				if err == nil {
					errChan <- fmt.Errorf("expected error when allocating invalid event, but got none")
					return
				}

				if event != nil {
					errChan <- fmt.Errorf("allocated event should be nil for invalid event code")
					return
				}
				time.Sleep(1 * time.Millisecond) // Simulate some delay
				pool.Free(invalidEventCode, event)
				if err == nil {
					errChan <- fmt.Errorf("expected error when freeing invalid event, but got none")
					return
				}
			}
		}()
	}
	wg.Wait()
	if len(errChan) > 0 {
		for err := range errChan {
			t.Log(err)
		}
		t.Fatal("Errors occurred during stress test")
	}
}
