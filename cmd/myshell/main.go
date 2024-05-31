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
    switch {
    case incoming == "exit 0":
        os.Exit(0)
    case strings.HasPrefix(incoming, "echo"):
        message := strings.Trim(incoming, "echo")
        fmt.Fprint(os.Stdout, message+"\n")

    default:
        fmt.Fprint(os.Stdout, incoming +": command not found\n")
    }
}
