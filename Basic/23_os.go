package main

import (
	"os"
	"syscall"
	"time"
)

func main() {
	hostname, _ := os.Hostname()
	println(hostname)

	computerName, _ := syscall.ComputerName()
	println(computerName)

	strings, _ := syscall.GetEnvironmentStrings()
	println(strings)

	now1 := time.Now()
	println(now1.Date())
	println("一年中的第", now1.YearDay(), "天")
	println(now1.Year(), now1.Month(), now1.Day(), now1.Minute(), now1.Second())

	os.Create("abc123.txt")

	file, _ := os.Open("123abc.txt")
	buf := make([]byte, 1024)
	file.Read(buf)
	println(string(buf))
	file.Close()

}
