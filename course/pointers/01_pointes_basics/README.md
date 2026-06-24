# Pointers — Basics

A pointer holds the memory address of a value. The `&` operator generates a pointer, and the `*` operator dereferences a pointer to access or modify the underlying value.

## Key Concepts

- `&x` → address of `x` (creates a pointer to `x`).
- `*p` → dereference `p` (reads or writes the value stored at that address).
- Function parameters declared as `*T` expect a pointer; the caller passes `&x`.
- Modifications via a pointer inside a function affect the original variable.

## Code Example

```go
score := 95
fmt.Println("Before:", score)

addScore(&score)

fmt.Println("After:", score)  // 100

func addScore(score *int) {
    *score = *score + 5
}
```

## Running the Program

```bash
go run main.go
```

The program passes a pointer to `score` into `addScore`, which modifies the original variable through the pointer.
