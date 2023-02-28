package main

import (
	"flag"
	"fmt"
)

var namePointer = flag.String("name", "Yu-yuan", "Your name")
var agePointer = flag.Int("age", 25, "Your age.")

func main() {
	flag.Parse()

	fmt.Printf("Hello, I am %s.\n", *namePointer)
	fmt.Println("I am", *agePointer, "years ago.")
}
