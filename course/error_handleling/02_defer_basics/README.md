# Defer Basics

The `defer` keyword schedules a function call to run immediately after the surrounding function returns (or panics). It is the idiomatic way to ensure cleanup always happens, regardless of the execution path.

## Key Concepts

- `defer` pushes a function call onto a stack; deferred calls execute in LIFO (last-in, first-out) order.
- Commonly used for resource cleanup: closing files, unlocking mutexes, closing network connections.
- Arguments to deferred functions are evaluated immediately, not when the deferred call runs.
- Even if the surrounding function panics or returns early, deferred calls still execute.

## Code Example

```go
func doWork(succ bool) error {
    fmt.Println("Starting work: resource aquire")
    defer fmt.Println("Finished work: resource released")

    if !succ {
        return errors.New("something went wrong")
    }

    fmt.Println("Doing work")
    fmt.Println("Work done")
    return nil
}
```

## Running the Program

```bash
go run main.go
```

The example shows how `defer` ensures the cleanup message is printed even when the function returns early with an error.
