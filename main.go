package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func execInput(input string) error {
	// Remove new line char
	input = strings.TrimSuffix(input, "\n")

	args := strings.Split(input, " ")

	// Prepare cmd
	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return errors.New("cd requires at least one argument")
		}

		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}
	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	return cmd.Run()
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		// Read the keyboad input.
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
