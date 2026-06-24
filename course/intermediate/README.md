# Intermediate

This module introduces methods in Go — functions with a special **receiver** parameter that attaches them to a type. Understanding the difference between value receivers and pointer receivers is key to writing idiomatic Go.

## Module Content

1. **01_methods_val_receivers** — Methods with value receivers: how they work, when to use them, and the implications of operating on a copy.
2. **02_methods_pointers_receivers** — Methods with pointer receivers: mutating the receiver's state and avoiding unnecessary copying.

## What I Have Learned

- A method is a function with a receiver parameter (`func (r ReceiverType) MethodName()`).
- Value receivers operate on a **copy** of the struct — modifications are local to the method.
- Pointer receivers can **mutate** the original struct and avoid copying large structs.
- Choose value receivers for small, immutable types; choose pointer receivers for mutable state or large structs.
- Methods are called with dot notation: `instance.Method()`.
