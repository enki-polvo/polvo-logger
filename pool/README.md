# Pool

`Pool` is a feature provided to increase memory reusability in the log processing pipeline of polvo's applications.
It is recommended that the `Pool` be managed as a **Singleton** at the program's runtime.

## Example

Here is an example usage:

```Go
import (
    ...
    eventModel "github.com/enki-polvo/polvo-logger/model/event"
    model "github.com/enki-polvo/polvo-logger/model"
    eventPool "github.com/enki-polvo/polvo-logger/pool"
)

var pool = eventPool.NewEventPool()

...

func PrintProcCreate() error  {
    // Get object from pool
    logModel, err := pool.Allocate(model.PROC_CREATE)
    if err != nil {
        return err
    }
    
    // Process logModel
    switch (logModel.EventCode) {
        case eventModel.PROC_CREATE:
            // In CommonModel, type casting is required because Metadata is any type.
            // The Allocate function returns an address. Similarly, the Metadata inside CommonModel is also an address.
            metadata := logModel.Metadata.(*eventModel.ProcessCreateMetadata)
            // One thing to note is that Allocate does not guarantee initialization of values ​​inside log.
            metadata.PID = 1234
            metadata.PPID = 5678
            metadata.UID = 1000
            metadata.Username = "root"
            metadata.Commandline = "bash rm -rf /tmp"
            metadata.ENV = "PATH=/usr/bin:/bin"
            metadata.Image = "/usr/bin/bash"
            metadata.TGID = 1234
            ...
    }

    fmt.Printf("%v, %v\n", logModel, logModel.Metadata)
    
    // Free method returns object to pool
    err = pool.Free(logModel)
    if err != nil {
        return err
    }
    
    return nil
}

```

Pool is basically threadSafe so can be used in multiple goroutines.
