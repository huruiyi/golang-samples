package main

import "../../mongodb"

func main() {
	if err := mongodb.Exec(); err != nil {
		panic(err)
	}
}
