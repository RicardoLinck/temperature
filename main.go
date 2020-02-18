package main

import (
	"flag"
	"fmt"
	"temperature/tempconv"
)

var temperature = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temperature)
}
