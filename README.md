# Go Solutions

Go (or Golang) is an open-source programming language developed by Google. It’s designed for simplicity, speed, and efficiency — combining the performance of C/C++ with the ease of use of modern languages. It’s especially popular for building scalable backend systems, cloud services, and microservices.

[Check the list of projects created](PROJECTS.md)

## key features of Go (Golang):S

1. 🧠 Simplicity – Clean, minimal syntax that’s easy to read and maintain.
2. ⚡ High Performance – Compiles to native machine code, close to C in speed.
3. 🔄 Concurrency Support – Built-in goroutines and channels make concurrent programming simple and efficient.
4. 📦 Fast Compilation – Very quick compile times compared to traditional compiled languages.
5. 🧰 Standard Library – Rich standard library for networking, HTTP, testing, and more.
6. 🧱 Static Typing – Strongly typed with compile-time checks for reliability.
7. 🌍 Cross-Platform – Easy to compile and run across different operating systems.
8. 🧹 Garbage Collection – Automatic memory management for simpler development.
9. 🧩 Built-in Tooling – Includes formatting (gofmt), testing, and dependency management tools.
10. ☁️ Great for Cloud/Server Apps – Ideal for microservices, APIs, and distributed systems.

## Most common Go CLI commands you’ll use:

- `go build` – Compiles the source code into an executable binary.
- `go run` – Compiles and runs a Go program in one step.
- `go mod init` – Initializes a new module (creates a go.mod file).
- `go get` – Downloads and installs packages or dependencies.\
- `go test` – Runs tests defined in \_test.go files.
- `go fmt` – Automatically formats code according to Go’s style guide.

## Package Main

In Go, package main indicates that the file belongs to an executable program, not a library.

🔹Explanation:

- Every Go program is organized into packages (package).
- When the package name is main, Go knows it should build an executable binary.
- This package must include a main() function, which is the entry point of the program (where execution starts).

### Example:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

#### Here:

- package main → tells Go to create an executable.
- func main() → the function that runs first when you execute the program.
- If you use a different package name (e.g., package utils), the file only defines reusable code (a library), not a runnable program.

## Types of Packages

In Go, packages are used to organize and reuse code. There are two main types of packages:

### Executable packages

- Declared with: package main
- Contain a func main() function — the program’s entry point.
- Compiled into an executable binary.

### Reusable packages

- Declared with any name other than main (e.g., package math, package utils).
- Contain reusable functions, structs, or constants.
- Cannot be run directly — they’re meant to be imported by other packages.
