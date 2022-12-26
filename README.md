# istor

istor is a small golang library to check if an ip address is a tor relay or not

# Installation

```bash
go get -u github.com/cmd777/istor
```

then, import it

```go
import (
    "github.com/cmd777/istor"
)
```

# Usage / Basic Example

```go
package main

import (
    "fmt"

    "github.com/cmd777/istor"
)

func main() {
    _, ResponseCode, err := istor.IsRelay("1.2.3.4", "")
	switch ResponseCode {
	case istor.IP_NOT_TOR: // 0
		fmt.Println("Not a TOR Relay")
	case istor.IP_TOR_RELAY: // 10
		fmt.Println("TOR Relay")
	default: // 1 - 9 are error codes
		fmt.Println("An error occoured...", err.Error())
	}
}
```