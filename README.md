<p>
  <a href="https://github.com/filinvadim/wall-clock-ticker/actions"><img alt="golang-action status" src="https://github.com/filinvadim/wall-clock-ticker/workflows/Go/badge.svg"></a>
</p>

# wall-clock-ticker

Golang wall clock ticker is a [Go](http://golang.org/) library that provides a ticker which ticks according to wall clock.

Installation
------------
```sh
go get -u github.com/filinvadim/wall-clock-ticker
```

Usage
--------

Using Golang wall clock ticker is very much like time.NewTicker with the addition of an accuracy and start time.

The following creates a ticker which ticks on the minute according the hosts wall clock with an accuracy of plus or minus one second.
```go
package main

import (
	"fmt"
	"time"

	wct "github.com/filinvadim/wall-clock-ticker"
)

func main() {
	// ticks every 59th minute of every hour
	t := wct.NewWCTicker(time.Minute*59, time.Second)
	for tick := range t.C {
		// Process tick
		fmt.Println("tick:", tick)
	}
}
```

License
-------
Golang wall clock ticker is available under the [BSD 2-Clause License](https://opensource.org/licenses/BSD-2-Clause).
