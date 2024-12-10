package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Uncomment this block to pass the first stage
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		inputStr, err := bufio.NewReader(os.Stdin).ReadString('\n')
		inputStr = strings.TrimRight(inputStr, "\n")
		inputStr = strings.TrimSpace(inputStr)
		if err != nil {
			log.Fatal(err)
		}

		switch inputStr {
		case "exit 0":
			os.Exit(0)
		}

		input := strings.SplitN(inputStr, " ", 2)
		switch input[0] {
		case "echo":
			fmt.Println(input[1])
		case "type":
			switch input[1] {
			case "echo":
				fmt.Println("echo is a shell builtin")
			case "exit":
				fmt.Println("exit is a shell builtin")
			default:
				fmt.Printf("%s: not found\n", input[1])
			}
		default:
			fmt.Printf("%s: command not found\n", inputStr)
		}
	}

}
