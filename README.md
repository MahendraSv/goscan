# goscan

A simple Java-like Scanner utility for Go.

## Installation

```bash
go get github.com/MahendraSv/goscan
```

## usage

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