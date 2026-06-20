# Go Foundations

This module contains the initial and fundamental concepts for learning to develop with the **Go** programming language. Through a series of short, practical examples, it covers basic syntax, primary data types, package structure, control flow, and the main collection types used in the language.

---

## 📂 Module Content

Here is the list of directories and topics covered in this section of the course:

1. **01_setup_first-program**
   - First program "Hello, World!".
   - Basic structure of an executable file: declaring `package main` and the `func main()` entry point.
2. **02_variables_and_types**
   - Explicit variable declaration using the `var` keyword and specifying common types (`string`, `int`, `float64`).
3. **03_packages_imports**
   - Using grouped `import` syntax to bring in external or standard library packages (`fmt`, `math`).
   - Accessing public package functions (`math.Sqrt()`).
4. **04_var_vs_short_declare**
   - Comparison between formal declaration with `var` and the short variable declaration operator (`:=`).
   - Multiple variable assignment in a single line (e.g., `likes, comments := 10, 20`).
5. **05_basic_types_string**
   - Basic string manipulation, concatenation using the `+` operator, and utilizing the `strings` standard library package (e.g., `strings.ToUpper()`).
6. **06_basic_types_int**
   - Different sizes of signed integers (`int8`, `int16`, `int32`, `int64`) and unsigned integers (`uint8`, `uint16`, `uint32`, `uint64`).
   - Basic arithmetic operations and increment (`++`) / decrement (`--`) operators.
7. **07_if_else**
   - Classical conditional structures.
   - Declaring initializer variables in the `if` header whose scope is restricted to the conditional blocks.
8. **08_for_classic_loop**
   - Loop structure and control using Go's only iteration keyword: `for`.
9. **09_switch**
   - Multi-case selection using `switch` and `default`.
10. **10_milestone_project**
    - First integrated project: a **Video Stats Analyzer** that combines variables, formatted printing (`fmt.Printf`), type conversions, `math`, `strings`, conditionals, `switch`, and `for` loops.
    - Introduction to constants via `const` (e.g., `const daysInWeek = 5`).
11. **11_arrays**
    - Fixed-size collection declaration with `[N]T` syntax.
    - Array literals (e.g., `scores := [5]int{...}`) and the `len()` function to obtain the length.
    - Note: arrays have a fixed size and cannot grow.
12. **12_slice_literals**
    - Declaration of dynamic slices with `[]T{...}`.
    - Adding elements dynamically with the built-in `append` function (single or multiple values).
    - Indexing and access by position, including the last element with `slice[len(slice)-1]`.
13. **13_len_and_capacity_of_slices**
    - Creating slices with `make([]T, length, capacity)` and understanding the difference between `len` (current length) and `cap` (underlying capacity).
    - Automatic capacity growth when `append` exceeds the original capacity.
    - Merging slices with the spread operator `...` (e.g., `append(todos, more...)`).
    - The `for range` loop to iterate over slices, obtaining the `index` and the `value`.
14. **14_maps**
    - Declaration of maps with the `map[keyType]valueType` syntax and map literals.
    - Creating maps with `make` and inserting/deleting key-value pairs (`delete`).
    - The **comma-ok** idiom: `value, ok := map[key]` to check whether a key exists.
    - Using `for range` to iterate over a map, obtaining `key` and `value`.
15. **15_functions**
    - Declaration of simple functions with parameters and return value (e.g., `func add(a, b int) int`).
    - **Multiple return values** (e.g., `sumAndProduct` returns `int, int`).
    - **Named return values** (e.g., `func divide(a, b int) (x int)` with a naked `return`).
    - **Variadic functions** that accept a variable number of arguments using `...T` (e.g., `sumAll(nums ...int)`).
    - **Anonymous functions** assigned to variables (function values).
    - **IIFE** (_Immediately Invoked Function Expression_): an anonymous function defined and executed in the same statement.

---

## 🧠 What I Have Learned

Upon completing these topics, I have gained a solid understanding of the fundamental pillars of Go:

### 1. Entry Point and Compilation

- **`main` Package**: I learned that any executable Go program must belong to the `main` package. Go looks for `package main` and a `func main()` declaration as the application's entry point.

### 2. Strong Typing and Variable Declaration

- **Type Inference vs. Explicit Typing**: Go is statically typed, but allows two ways of declaring variables:
  - **Explicit typing**: `var variable type = value` (e.g., `var year int = 2026`). Useful for package-level declarations or when the variable is not initialized immediately.
  - **Short declaration (`:=`)**: Automatically infers the type based on the assigned value (e.g., `population := 40000`). This can only be used inside function bodies.
- **Multiple Assignment**: Multiple variables can be declared and initialized simultaneously using commas (e.g., `likes, comments := 10, 20`).
- **Constants (`const`)**: Values that cannot be modified after declaration (e.g., `const daysInWeek = 5`), ideal for fixed values used throughout the program.

### 3. Modularity and Imports

- **`import` Blocks**: It is good practice to group required packages within parentheses.
- **Exported Functions**: Go packages expose functions by starting their names with a capital letter (e.g., `math.Sqrt`, `strings.ToUpper`). If they start with a lowercase letter, they are package-private.

### 4. Precision of Numeric Types

- **Memory Optimization**: I understood the importance of selecting the appropriate integer type based on expected values. Types like `uint8` (0 to 255) are ideal for small unsigned values, while signed `int` types allow negative and positive values depending on the architecture (32 or 64 bits).
- **Type Conversion**: To perform operations between numeric types, an explicit conversion is required (e.g., `float64(likes+comments) / float64(views)`). Go does not perform automatic numeric promotion.

### 5. Clean Control Flow

- **Initialization in Conditionals**: A variable can be initialized directly within the `if` statement (e.g., `if total := items * pricePerItem; total >= 100`). The variable `total` only exists within the scope of the `if` and `else` blocks, helping to keep scopes narrow.
- **Simplicity in Loops**: Go does not have a `while` statement. All repetitive loops are expressed using the `for` keyword, including the classic form and the `for range` form.
- **`for range`**: Used to iterate over slices, maps, strings, and channels. Returns `index, value` for slices and `key, value` for maps. The blank identifier `_` is used when you only need the value or the key.
- **Switch statements without implicit fallthrough**: In a `switch` statement, Go executes only the matching case and does not automatically fall through to the next cases. No manual `break` statement is required at the end of each `case`. A `switch` with no expression works as a clean chain of `if/else`.

### 6. Formatted Output

- **`fmt.Printf`**: Allows you to print text with format verbs such as `%d` (integers), `%f` (floating-point), `%.2f` (2 decimals), and `%s` (strings). Escape sequences like `\n` are used for new lines.

### 7. Collection Types

- **Arrays (`[N]T`)**: Fixed-size structures. Their size is part of the type, so `[3]int` and `[5]int` are different. They cannot grow, so they are rarely used directly in business code.
- **Slices (`[]T`)**: The most common collection in Go. They are dynamic and can grow with `append`. They are declared as literals (`[]string{...}`) or with `make([]T, len, cap)`.
- **`len` vs `cap`**: `len` is the number of elements currently in the slice; `cap` is the maximum size that the underlying array can hold before a new allocation is required. When you `append` and exceed `cap`, Go automatically allocates a new (larger) backing array.
- **Spread operator (`...`)**: Used to expand a slice into individual arguments (e.g., `append(todos, more...)`).
- **Maps (`map[K]V`)**: Unordered collections of key-value pairs. They are created with literals, with `make`, and manipulated with `delete(map, key)`.
- **Comma-ok idiom**: `val, ok := map[key]` returns the value and a boolean indicating whether the key actually exists. This is the idiomatic way to safely check existence and avoid the zero value being confused with "not present".

### 8. Functions

- **Multiple return values**: Go functions can return more than one value (e.g., `func sumAndProduct(a, b int) (int, int)`), very useful for returning a result along with an error or an additional state.
- **Named return values**: It is possible to name the return values in the signature (e.g., `(x int)`) and use a naked `return` to return them, which improves readability in short functions.
- **Variadic functions**: By appending `...` to the last parameter (e.g., `nums ...int`), the function can receive any number of arguments of that type. Internally, they are treated as a slice.
- **Functions as values**: A function can be assigned to a variable and passed as an argument (e.g., `res := func(n int) int { return n * n }`).
- **IIFE**: An anonymous function can be defined and invoked in the same expression (e.g., `func(a, b int) int { return a + b }(5, 10)`), useful for encapsulating isolated logic.
