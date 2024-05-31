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
    
        go responseHandler(input)
    }
}

func responseHandler(incoming string) {
    fmt.Fprint(os.Stdout, strings.TrimRight(incoming, "\n") +": command not found\n")
}
