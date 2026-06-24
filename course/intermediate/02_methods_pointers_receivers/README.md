# Methods — Pointer Receivers

A method with a **pointer receiver** (`*T`) has access to the original value and can mutate its fields. It also avoids copying the entire struct when the method is called.

## Key Concepts

- Syntax: `func (u *User) MethodName()`
- Changes to `u`'s fields inside the method affect the original instance.
- Pointer receivers are the standard choice when a method needs to modify the receiver's state.
- Also preferred for large structs to avoid the performance cost of copying.
- Even if a variable is not a pointer, Go automatically takes its address when calling a pointer-receiver method (if the variable is addressable).

## Code Example

```go
type User struct {
    Name string
    Age  int
}

func (u *User) Birthday() {
    u.Age++
}
```

## Running the Program

```bash
go run main.go
```

The program creates a `User`, calls `Birthday()` which increments the age via a pointer receiver, and prints the updated value.
