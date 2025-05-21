package main

import (
	"fmt"
	"os/exec"
)

func runC() {
	app := "gcc"
	arg0 := "../../c/test.cpp"
	cmd := exec.Command(app, arg0)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println("error: ", err.Error())
	}
	fmt.Println(string(stdout))

	cmd2 := exec.Command("../../c/a.out")
	stdout2, err2 := cmd2.Output()
	if err2 != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(stdout2))
}
