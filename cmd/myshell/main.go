package main

import (
	"bufio"
	"fmt"
	"os"
    "strings"
)

func main() {
    for {
        fmt.Fprint(os.Stdout, "$ ")

	    // Wait for user input
        input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
    
        responseHandler(strings.TrimSpace(input))
    }
}

func responseHandler(incoming string) {
    cmd, args := getCmdAndArgs(incoming)
    
    switch cmd{
    case "exit 0":
        os.Exit(0)

    case "echo":
        fmt.Fprint(os.Stdout, strings.Join(args," "))

    default:
        fmt.Fprint(os.Stdout, incoming +": command not found\n")
    }
}
