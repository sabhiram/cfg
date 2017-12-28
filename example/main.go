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
