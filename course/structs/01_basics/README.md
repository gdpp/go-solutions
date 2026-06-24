# Structs — Basics

A struct is a collection of named fields that lets you group related data into a single type.

## Key Concepts

- Define a struct with `type StructName struct { … }`.
- Each field has a name and a type.
- Create instances with `StructName{Field: value, …}`.
- Access fields with dot notation: `instance.Field`.
- Fields not explicitly set get their zero value (`0`, `""`, `false`, `nil`).

## Code Example

```go
type User struct {
    Id    int
    Name  string
    Email string
    Age   int
}

u1 := User{
    Id:    1,
    Name:  "Gustavo",
    Email: "gdpp@dev.com",
    Age:   33,
}

u2 := User{Id: 2}  // Name, Email, Age get zero values
```

## Running the Program

```bash
go run main.go
```

The program creates a fully initialized `User` and a partially initialized one, then prints both to show how missing fields default to zero values.
