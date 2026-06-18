# Go Foundations

This module contains the initial and fundamental concepts for learning to develop with the **Go** programming language. Through a series of short, practical examples, it covers basic syntax, primary data types, package structure, and control flow.

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

---

## 🧠 What I Have Learned

Upon completing these topics, I have gained a solid understanding of the fundamental pillars of Go:

### 1. Entry Point and Compilation
* **`main` Package**: I learned that any executable Go program must belong to the `main` package. Go looks for `package main` and a `func main()` declaration as the application's entry point.

### 2. Strong Typing and Variable Declaration
* **Type Inference vs. Explicit Typing**: Go is statically typed, but allows two ways of declaring variables:
  * **Explicit typing**: `var variable type = value` (e.g., `var year int = 2026`). Useful for package-level declarations or when the variable is not initialized immediately.
  * **Short declaration (`:=`)**: Automatically infers the type based on the assigned value (e.g., `population := 40000`). This can only be used inside function bodies.
* **Multiple Assignment**: Multiple variables can be declared and initialized simultaneously using commas (e.g., `likes, comments := 10, 20`).

### 3. Modularity and Imports
* **`import` Blocks**: It is good practice to group required packages within parentheses.
* **Exported Functions**: Go packages expose functions by starting their names with a capital letter (e.g., `math.Sqrt`, `strings.ToUpper`). If they start with a lowercase letter, they are package-private.

### 4. Precision of Numeric Types
* **Memory Optimization**: I understood the importance of selecting the appropriate integer type based on expected values. Types like `uint8` (0 to 255) are ideal for small unsigned values, while signed `int` types allow negative and positive values depending on the architecture (32 or 64 bits).

### 5. Clean Control Flow
* **Initialization in Conditionals**: A variable can be initialized directly within the `if` statement (e.g., `if total := items * pricePerItem; total >= 100`). The variable `total` only exists within the scope of the `if` and `else` blocks, helping to keep scopes narrow.
* **Simplicity in Loops**: Go does not have a `while` statement. All repetitive loops are expressed using the `for` keyword.
* **Switch statements without implicit fallthrough**: In a `switch` statement, Go executes only the matching case and does not automatically fall through to the next cases. No manual `break` statement is required at the end of each `case`.
