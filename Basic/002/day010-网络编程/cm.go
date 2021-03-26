package main

import "os/exec"

func main() {
	//exec.Command("calc").Run()

	command := exec.Command("tasklist", "/V")
	bytes, err := command.Output()

	println(string(bytes))
	println(err)

}
