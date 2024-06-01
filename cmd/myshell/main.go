package main

import (
	"bufio"
	"fmt"
	"os"
    "os/exec"
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
    case "pwd":
        pwd, _ := os.Getwd()
        fmt.Fprint(os.Stdout, pwd+"\n")

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
                    fmt.Fprint(os.Stdout, check+" is "+exec+"\n")
                    return
                }
            }
            fmt.Fprint(os.Stdout, check+" not found\n")
        }

    case "exit":
        if cmds[1] == "0" {
            os.Exit(0)
        } else { 
            fmt.Fprint(os.Stdout, "Not a valid exit code\n")
        }

    case "cd":
        err := os.Chdir(cmds[1]); if err != nil {
            fmt.Fprint(os.Stdout, cmds[1]+": No such file or directory\n")
	}

    default:
        command := exec.Command(cmds[0], cmds[1:]...)
		command.Stderr = os.Stderr
		err := command.Run()
		if err != nil {
            fmt.Fprint(os.Stdout, cmd[0]+": command not found\n")
        }
    }
}
