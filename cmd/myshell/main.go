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
    fmt.Errorf("%v: command not found\n", strings.TrimRight(incoming, "\n"))
}
