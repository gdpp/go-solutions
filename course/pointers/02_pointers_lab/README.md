# Pointers — Lab: Swap

A hands-on lab that implements a classic `swap` function using pointers to exchange the values of two variables.

## Objective

Write a function `swap` that takes two integer pointers and exchanges the values they point to.

## Key Concepts

- Using a temporary variable to hold one value during the swap.
- Dereferencing pointers to read and write values.
- After the swap, the original variables have their values exchanged.

## Code Example

```go
a, b := 10, 20
swap(&a, &b)
fmt.Println(a, b) // 20, 10

func swap(x *int, y *int) {
    temp := *x
    *x = *y
    *y = temp
}
```

## Running the Program

```bash
go run main.go
```

The program creates two variables, swaps their values using pointers, and prints the result.
