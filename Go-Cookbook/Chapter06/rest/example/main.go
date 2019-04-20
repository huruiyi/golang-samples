package main

import "../../rest"

func main() {
	if err := rest.Exec(); err != nil {
		panic(err)
	}
}
