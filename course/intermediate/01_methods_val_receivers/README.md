# Methods — Value Receivers

A method with a **value receiver** operates on a copy of the original value. This means any changes made to the receiver's fields inside the method do not affect the original instance.

## Key Concepts

- Syntax: `func (u User) MethodName() ReturnType`
- The receiver `u` is a copy of the original `User` value.
- Value receivers are safe to call from any goroutine — no shared state is mutated.
- Use value receivers for small structs or when the method does not need to modify the receiver.

## Code Example

```go
type User struct {
    Name string
    Age  int
}

func (u User) Intro() string {
    return fmt.Sprintf("Hi, I'm %s and I'm %d", u.Name, u.Age)
}
```

## Running the Program

```bash
go run main.go
```

The program creates a `User` and calls the `Intro()` method which returns a formatted introduction string.
