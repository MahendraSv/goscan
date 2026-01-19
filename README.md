# goscan

Let’s build a custom input utility in Go, similar to Java’s Scanner, that is:

- ✅ Easy to use
- ✅ Type-safe
- ✅ Handles spaces & newlines
- ✅ Centralized error handling
- ✅ Suitable for real projects & competitive programming

A simple Java-like Scanner utility for Go.

## Installation

```bash
go get github.com/MahendraSv/goscan
```

## Usage

```
import "github.com/MahendraSv/goscan/scanner"

in := scanner.New()

name := in.String("Name: ")
age := in.Int("Age: ")
```

## Features
- Full line input
- Type-safe parsing
- Retry on invalid input
- CLI-friendly


### Why This Is Idiomatic Go (Even With Generics)
| Feature                  | Value |
| ------------------------ | ----- |
| Compile-time type safety | ✅     |
| No reflection            | ✅     |
| Simple API               | ✅     |
| Reusable parsing logic   | ✅     |
| Go-style error loops     | ✅     |

## Run Tests

From the module root:
```
go test ./...
```

Expected output:
```
ok   github.com/MahendarSv/goscan/scanner  0.00Xs
```

### Coverage
``` go
go test -cover ./...
```