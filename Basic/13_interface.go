package main

import (
	"fmt"
)

type Connection interface {
	Connect()
}

type USB interface {
	Name() string
	Connection
}

type PhoneConnecter struct {
	name string
}

func (phone PhoneConnecter) Name() string {
	return phone.name
}

func (phone PhoneConnecter) Connect() {
	fmt.Println("Connected:", phone.Name())
}

func main() {
	phone := PhoneConnecter{"iPhone"}
	phone.Connect()
	DisConnect(phone)
	DisConnect2(phone)
}

func DisConnect(c Connection) {
	if phone, ok := c.(PhoneConnecter); ok {
		fmt.Println("DisConnect:", phone.Name())
		return
	}
	fmt.Println("Unknown Connection")
}

func DisConnect2(c interface{}) {
	switch v := c.(type) {
	case PhoneConnecter:
		fmt.Println("DisConnect2:", v.Name())
	default:
		fmt.Println("Unknown Connection")
	}
}
