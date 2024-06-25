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
		cmd := exec.Command("sudo", "pacman", "-S", "plasma", "dolphin", "konsole", "sddm")
		cmdq := exec.Command("systemctl", "enable", "sddm.service")

		cmd.Run()
		cmdq.Run()
	} else {
		fmt.Println("Invalid input")
	}
}
