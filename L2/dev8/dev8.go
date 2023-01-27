package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func cd(dir string) {
	if err := os.Chdir(dir); err != nil {
		fmt.Println(err.Error())
	}
}

func pwd() string {
	if wd, err := os.Getwd(); err != nil {
		fmt.Println(err)
		return ""
	} else {
		return wd
	}
}

func echo(args []string) {
	for _, arg := range args {
		fmt.Print(arg + " ")
	}
	fmt.Println()
}

func ps() {
	cmd := exec.Command("ps", "-A")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("ps: ", err)
	}
	fmt.Println(string(output))
}

func kill(pid int) {
	if p, err := os.FindProcess(pid); err != nil {
		fmt.Println(err.Error())
		return
	} else {
		if err := p.Kill(); err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}

func execute(query string) {
	commands := strings.Split(query, " | ")
	for _, command := range commands {
		commandSlice := strings.Split(command, " ")
		switch commandSlice[0] {
		case "pwd":
			fmt.Println(pwd())
		case "cd":
			cd(commandSlice[1])
		case "echo":
			echo(commandSlice[1:])
		case "ps":
			ps()
		case "kill":
			if pid, err := strconv.Atoi(commandSlice[1]); err != nil {
				fmt.Println(err.Error())
				return
			} else {
				kill(pid)
			}
		}
	}
}

func main() {
	wd, _ := os.Getwd()
	fmt.Print(wd + "> ")
	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); fmt.Print("$ ") {
		if query := scanner.Text(); query != "\\quit" {
			execute(query)
		} else {
			break
		}
		wd, _ = os.Getwd()
	}
}
