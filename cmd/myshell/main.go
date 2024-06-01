package main

import (
	"bufio"
	"fmt"
	"os"
    "os/exec"
    "strings"
    "slices/Contains"
    "path/filepath"
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
        pathToChange := cmds[1]
        path := parsePath(pathToChange)
        
        err := os.Chdir(path)
        if err != nil {
            fmt.Println(path + ": No such file or directory")
        }

    default:
        command := exec.Command(cmds[0], cmds[1:]...)
		command.Stderr = os.Stderr
        command.Stdout = os.Stdout
		err := command.Run()
		if err != nil {
            fmt.Fprint(os.Stdout, cmds[0]+": command not found\n")
        }
    }
}

func parsePath(path string) string {
    // Absolute Path
	if strings.HasPrefix(path, "/") {
		return path
	} else if strings.HasPrefix(path, "~") {
        homedir := os.Getenv("HOME")
        return homedir
    }
	// Relative Path
	currentPath, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return filepath.Join(currentPath, path)
}
