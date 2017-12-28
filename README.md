# cfg

load tagged structs from JSON.

## Install

```
go get github.com/sabhiram/cfg
```

## Usage

To run the example (from the root directory of the library):

```
go run ./example/main.go
```

The main (and only) library API is the `cfg.Load` method.  It accepts a path to
a json file and an `interface{}` backed by a tagged struct as shown below.
	
```
package main

import (
	"fmt"
	"os"

	"github.com/sabhiram/cfg"
)

// Config implements a structure that we want to use to inject variables in from
// a JSON file.
type Config struct {
	A string `cfg:"A"`
	B int    `cfg:"B,required"`
}

func main() {
	var config Config

	if err := cfg.Load("./example/config.json", &config); err != nil {
		fmt.Printf("config error :: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("A: %s\n", config.A)
	fmt.Printf("B: %d\n", config.B)
}

```

## All else

Leave it here: https://github.com/sabhiram/cfg/issues
