package main

import "../../storage"

func main() {
	if err := storage.Exec(); err != nil {
		panic(err)
	}
}
