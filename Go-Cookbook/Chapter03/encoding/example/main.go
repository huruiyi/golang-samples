package main

import (
	"../../encoding"
)

func main() {
	if err := encoding.Base64Example(); err != nil {
		panic(err)
	}

	if err := encoding.Base64ExampleEncoder(); err != nil {
		panic(err)
	}

	if err := encoding.GobExample(); err != nil {
		panic(err)
	}
}
