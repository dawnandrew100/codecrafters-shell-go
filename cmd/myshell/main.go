package main

import (
	"bufio"
	"fmt"
	"os"
    "strings"
    "slices"
)

var built_ins = []string{"exit", "echo", "type"} 

func main() {
    for {
        fmt.Fprint(os.Stdout, "$ ")

	    // Wait for user input
        input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
    
        responseHandler(strings.TrimSpace(input))
    }
}

func responseHandler(incoming string) {
    cmds := strings.Split(incoming, " ")

    switch cmds[0]{
    case "exit":
        if cmds[1] == "0" {
            os.Exit(0)
        } else { 
            fmt.Fprint(os.Stdout, "Not a valid exit code\n")
        }

    case "echo":
        message := strings.TrimPrefix(incoming, "echo ")
        fmt.Fprint(os.Stdout, message+"\n")

    case "type":
        check := strings.TrimPrefix(incoming, "type ")
        if slices.Contains(built_ins, check){
            fmt.Fprint(os.Stdout, check+" is a shell builtin\n")
        } else {
            env := os.Getenv("PATH")
            paths := strings.Split(env, ":")
            for _, path := range paths {
                exec := path + "/" + check
                if _, err := os.Stat(exec); err == nil {
                    fmt.Fprintf(os.Stdout, "%v is %v\n", check, exec)
                    return
                }
            }
            fmt.Fprint(os.Stdout, check+" not found\n")
        }

    default:
        fmt.Fprint(os.Stdout, incoming +": command not found\n")
    }
}
