package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Format(time.Kitchen))

	// also, you can input directly as string :
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(time.Now().Format("15:04:05"))

	// Timestamp
	timestamp := time.Now().Unix()
	fmt.Println(timestamp) // Ex: 1257894000

	tm := time.Unix(1573020865, 0)
	fmt.Println(tm)

}
