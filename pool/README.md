# Pool

`Pool` is a feature provided to increase memory reusability in the log processing pipeline of polvo's applications.
It is recommended that the `Pool` be managed as a **Singleton** at the program's runtime.

## Example

Here is an example usage:

```Go
import (
    ...
    eventModel "github.com/enki-polvo/polvo-logger/model/event"
    eventPool "github.com/enki-polvo/polvo-logger/pool"
)

var pool = eventPool.NewEventPool()

...

func PrintProcCreate() error  {
    // Get object from pool
    logModel, err := pool.Allocate(eventModel.PROC_CREATE)
    if err != nil {
        return err
    }
    
    // Print logModel
    ...
    
    // Free method returns object to pool
    err = pool.Free(eventModel.PROC_CREATE, logModel)
    if err != nil {
        return err
    }
    
    return nil
}

```

Pool is basically threadSafe so can be used in multiple goroutines.
