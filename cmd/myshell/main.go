package main

import (
	"bufio"
	"fmt"
	"os"
    "strings"
    "slices"
)

built_ins := []string{"exit 0", "echo", "type"} 

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
        message := strings.TrimPrefix(incoming, "echo ")
        fmt.Fprint(os.Stdout, message+"\n")

    case strings.HasPrefix(incoming, "type"):
        check := strings.TrimPrefix(incoming, "type ")
        check = strings.TrimSpace(check)
        if slices.Contains(built_ins, "check"){
            fmt.Fprint(os.Stdout, check+" is a shell builtin\n")
        } else {
            fmt.Fprint(os.Stdout, check+" not found\n")
        }

    default:
        fmt.Fprint(os.Stdout, incoming +": command not found\n")
    }
}
