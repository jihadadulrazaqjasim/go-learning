# Go Learning

Projects built while following [Go: The Complete Developer's Guide (Golang)](https://www.udemy.com/course/go-the-complete-developers-guide/) by Stephen Grider.

---

## 01-deck-and-tests

A card deck application covering Go fundamentals.

**Concepts learned:**
- Variables, functions, and return values
- Custom types (`type deck []string`, `type color string`)
- Receiver functions (methods on types)
- Slices: creating, appending, and range-based iteration
- Slice splitting for dealing cards (`d[:n]`, `d[n:]`)
- String manipulation with `strings.Join` and `strings.Split`
- File I/O with `os.WriteFile` and `os.ReadFile`
- Random shuffling with `math/rand/v2`
- **Testing** with Go's `testing` package (`go test`)
- Test cleanup patterns (`os.Remove`)

---

## 02-custom-types

Exploring custom types and receiver functions with a book type and deck.

**Concepts learned:**
- Defining custom types from primitives (`type book string`)
- Receiver functions: `printTitle()`, `countLetters()`, `titleToUpperCase()`
- Using `strings.ToUpper` for string conversion
- Rewriting the deck type and `newDeck()` for practice

---

## 03-even-odd

A simple loop exercise printing whether numbers are even or odd.

**Concepts learned:**
- `for` loops with `range`
- Modulo operator (`%`) for conditionals
- `continue` keyword to skip loop iterations

---

## 04-structs

Modeling a person with nested contact info using structs.

**Concepts learned:**
- Defining structs (`person`, `contactInfo`)
- Nested/embedded structs (accessing `jihad.email` directly)
- Struct field updates
- Pointers and dereferencing (`&`, `*`)
- Pointer receivers (`func (p *person) updateName()`)
- Go's automatic pointer/value conversion on method calls

---

## 05-interfaces

A payment system demonstrating interfaces with CreditCard, PayPal, and Cash.

**Concepts learned:**
- Defining interfaces (`PaymentMethod` with `Process` and `GetName`)
- Multiple types satisfying the same interface
- Error handling with `fmt.Errorf` and `error` return type
- String slicing for masking card numbers
- Polymorphism: passing different types to a single `Checkout` function

---

## 06-interfaces-ii

A document processing system using Logger and Document interfaces.

**Concepts learned:**
- Interface composition (separate `Logger` and `Document` interfaces)
- Struct with interface field (`Processor` has a `Logger`)
- Dependency injection (swapping `ConsoleLogger` for `FileLogger`)
- Method call chaining through interfaces

---

## 07-interfaces-iii

Advanced multi-interface document processor with type assertions.

**Concepts learned:**
- Multiple interfaces: `Document`, `Compressible`, `Encryptable`, `Logger`
- **Type assertions** (`doc.(Encryptable)`) with the comma-ok pattern
- Selective behavior based on which interfaces a type implements
- Command-line arguments with `os.Args`
- `strings.LastIndex` for file extension parsing
- `switch` statements for routing logic

---

## 08-reader-writer

Shape area calculator and custom `io.Writer` implementation.

**Concepts learned:**
- Interfaces for polymorphism (`shape` with `area()`)
- `fmt.Printf` with `%T` for printing types
- Go's `io.Reader` and `io.Writer` interfaces
- HTTP requests with `net/http` (commented out)
- Implementing `Write([]byte) (int, error)` to satisfy `io.Writer`
- `io.Copy` to pipe data between Reader and Writer

---

## 09-file-reader

Read and print a file's contents using `os.Open` and `io.Copy`.

**Concepts learned:**
- Opening files with `os.Open`
- File metadata via `file.Stat()` (name, size, isDir)
- Piping file content to stdout with `io.Copy(os.Stdout, file)`
- Command-line arguments (`os.Args[1]`)
- Error handling and `os.Exit(1)`

---

## 10-channels

Order notification system using goroutines and channels.

**Concepts learned:**
- Goroutines (`go` keyword for concurrent execution)
- Channels: creating, sending, and receiving (`chan`, `<-`)
- Directional channels (`<-chan` for receive-only, `chan<-` for send-only)
- Worker pattern (goroutine processing jobs from a channel)
- `close(channel)` to signal no more data
- Blocking with `<-done` to wait for goroutine completion
- `time.Sleep` to simulate work

---

## 11-user-crud

A user management service with full CRUD operations, organized into packages.

**Concepts learned:**
- **Package organization** (`models/`, `service/`)
- Exported vs unexported types (capitalized names)
- Custom type aliases (`type UserID int`)
- Methods on structs (`IsAdult()`)
- Pointer receivers for mutation (`*UserService`)
- Constructor pattern (`NewUserService()`)
- Slice manipulation for delete (`append(s[:i], s[i+1:]...)`)
- Error handling across service methods
- `go.mod` and module management
