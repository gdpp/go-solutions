# Pointers

Pointers store the memory address of a value, enabling functions to modify variables outside their local scope and to avoid copying large data structures.

## Module Content

1. **01_pointes_basics** — Pointer fundamentals: creating pointers with `&`, dereferencing with `*`, and passing pointers to functions.
2. **02_pointers_lab** — Hands-on lab implementing a `swap` function using pointers to exchange two values.

## What I Have Learned

- `&x` produces a pointer to `x` (the memory address of `x`).
- `*p` dereferences pointer `p` — it reads or writes the value at that address.
- A pointer type is written as `*T`, e.g., `*int` is a pointer to an `int`.
- Passing a pointer to a function allows that function to modify the original variable.
- Pointers are useful for avoiding copies of large structs or slices.
- Go does not allow pointer arithmetic (unlike C/C++), which improves safety.
