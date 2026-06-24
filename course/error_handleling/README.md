# Error Handling

This module covers Go's approach to error handling, which differs from exception-based languages. In Go, errors are returned as regular values from functions, and the `defer` keyword provides a clean mechanism for resource cleanup.

## Module Content

1. **01_error_return_pattern** — The idiomatic pattern of returning errors as values, using `if err != nil` checks, and creating descriptive errors with `fmt.Errorf`.
2. **02_defer_basics** — The `defer` keyword for scheduling cleanup actions, resource release, and execution order when a function returns or panics.

## What I Have Learned

- Go treats errors as values — there are no exception handling constructs like `try`/`catch`.
- The error return pattern (`val, err := someFunc(); if err != nil { … }`) is fundamental to Go code.
- `fmt.Errorf` creates structured error messages, often wrapping context about what went wrong.
- The zero value for the `error` interface is `nil`, meaning no error occurred.
- `defer` schedules a function call to run after the surrounding function completes.
- Deferred calls execute in LIFO (last-in, first-out) order.
- `defer` is commonly used for closing resources (files, network connections, etc.).
- A function can return an error while still running deferred cleanup logic.
