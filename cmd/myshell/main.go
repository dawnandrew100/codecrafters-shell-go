package main

import (
	"bufio"
	"fmt"
	"os"
    "strings"
)

func main() {
    fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
    input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
    fmt.Fprint(os.Stdout, strings.TrimRight(input, "\n")+": command not found"
}
