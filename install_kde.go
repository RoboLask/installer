package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var i string
	fmt.Print("Enter your choice: ")
	fmt.Scan(&i)
	if i == "kde" {
		cmd := exec.Command("sudo", "pacman", "-S", "plasma", "dolphin", "konsole")
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(string(output))
	} else {
		fmt.Println("Invalid input")
	}
}
