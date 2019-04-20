package main

import "../../decorator"

func main() {
	if err := decorator.Exec(); err != nil {
		panic(err)
	}
}
