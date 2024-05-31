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
    
        go responseHandler(strings.TrimSpace(input))
    }
}

func responseHandler(incoming string) {
    fmt.Fprint(os.Stdout, incoming +": command not found\n")
}
