package main

import (
	"fmt"
	"github.com/yosssi/gocmd"
	"os/exec"
	"strings"
)

const (
	msgNothingFormat = "\nThere is nothing to format.\n"
)

func main() {
	output, err := gocmd.Pipe(exec.Command("find", "."), exec.Command("grep", "\\.go"))
	if err != nil {
		fmt.Println(msgNothingFormat)
		return
	}
	files := convOutputToStrings(output)
	targets := []string{}
	for _, file := range files {
		output, err := exec.Command("gofmt", "-l", file).Output()
		if err != nil {
			continue
		}
		target := convOutputToStrings(output)
		if len(target) == 0 {
			continue
		}
		targets = append(targets, target[0])
	}
	if len(targets) == 0 {
		fmt.Println(msgNothingFormat)
		return
	}
	fmt.Println("\nThe following go files will be formatted:\n")
	for _, target := range targets {
		fmt.Println(target)
	}
	var input string
	for {
		fmt.Print("\nAre you sure to format them? [y/n] ")
		fmt.Scanln(&input)
		if input == "y" {
			break
		} else if input == "n" {
			fmt.Println("\nThis process was canceled.\n")
			return
		} else {
			fmt.Println("\nInput must be \"y\" or \"n\".")
		}
	}
	for _, target := range targets {
		err := exec.Command("gofmt", "-w", target).Run()
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("\nDone.\n")
}

func convOutputToStrings(output []byte) []string {
	s := strings.Split(string(output), "\n")
	s = s[:len(s)-1]
	return s
}
