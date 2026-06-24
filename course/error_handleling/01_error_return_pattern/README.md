# Error Return Pattern

In Go, functions communicate failures by returning an `error` value as part of their normal return signature. There are no exceptions for routine errors — this explicit approach makes error handling visible and deliberate.

## Key Concepts

- Functions return errors as the last (or only) return value.
- A `nil` error means the call succeeded; a non-`nil` error means something went wrong.
- The `if err != nil` check is the standard way to handle errors.
- `fmt.Errorf` formats and creates a new error message, optionally wrapping context.

## Code Example

```go
func parseLevel(s string) (int, error) {
    n, err := strconv.Atoi(s)
    if err != nil {
        return 0, fmt.Errorf("Level must be a number")
    }
    if n < 1 || n > 5 {
        return 0, fmt.Errorf("Level must be 1 to 5")
    }
    return n, nil
}
```

## Running the Program

```bash
go run main.go
```

The program attempts to parse a level value from a string, validating that it's a valid number and falls within the range 1–5.
