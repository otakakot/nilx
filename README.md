# nilx

`nilx` is a library that provides utility functions to easily handle `nil` pointers in Go.

## Installation

```sh
go get github.com/otakakot/nilx
```

## Usage

### NilOr

The NilOr function returns the first non-nil value from multiple pointers. 
If all pointers are nil, it returns the zero value of the type.

```go
package main

import (
    "fmt"

    "github.com/otakakot/nilx"
)

func main() {
    val := 1
    result := nilx.NilOr(nil, &val)
    fmt.Println(result) // Output: 1
}
```

### NilZero

The NilZero function returns the zero value of the type if the pointer is nil, otherwise it returns the value of the pointer.

```go
package main

import (
    "fmt"

    "github.com/otakakot/nilx"
)

func main() {
    result := nilx.NilZero[int](nil)
    fmt.Println(result) // Output: 0
}
```

### NilDef

The NilDef function returns the default value if the pointer is nil, otherwise it returns the value of the pointer.

```go
package main

import (
    "fmt"

    "github.com/otakakot/nilx"
)

func main() {
    result := nilx.NilDef(nil, 1)
    fmt.Println(result) // Output: 1
}
```

## License

This project is licensed under the Apache License 2.0.
