# Structs

Structs are composite data types that group related fields under a single name. They are Go's primary mechanism for creating custom types with named properties.

## Module Content

1. **01_basics** — Declaring struct types, creating struct instances, accessing fields, and partial initialization.

## What I Have Learned

- A struct is defined with the `type` keyword: `type User struct { … }`.
- Fields are named and typed: `Name string`, `Age int`, etc.
- Struct instances are created with a composite literal: `User{Name: "Alice", Age: 30}`.
- Fields are accessed with dot notation: `user.Name`.
- Uninitialized fields get zero values (0 for int, "" for string, nil for pointers, etc.).
- Structs can be partially initialized — omitted fields default to their zero values.
