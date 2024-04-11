package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)



func main() {
	
	fmt.Println(" \033[34m____  _              ____  _          _ _\033[0m")
	fmt.Println("\033[34m| __ )| |_   _  ___  / ___|| |__   ___| | |\033[0m")
	fmt.Println("\033[34m|  _ \\| | | | |/ _ \\ \\___ \\| '_ \\ / _ \\ | |\033[0m")
	fmt.Println("\033[34m| |_) | | |_| |  __/  ___) | | | |  __/ | |\033[0m")
	fmt.Println("\033[34m|____/|_|\\__,_|\\___| |____/|_| |_|\\___|_|_|\033[0m")
		
    reader := bufio.NewReader(os.Stdin)
    for {
		fmt.Print("\033[34mblue_shell\033[0m> ")        		
		
        // Read the keyboad input.
        input, err := reader.ReadString('\n')
        if err != nil {
            fmt.Printf("Error reading string: %v", err)
        }

        // Handle the execution of the input.
        if err = execInput(input); err != nil {
            fmt.Printf("Error executing input command with \"%v\"\n", err)			
        }
    }
}

func execInput(input string) error {
    // Remove the newline character.
    input = strings.TrimSuffix(input, "\n")

    // Prepare the command to execute.
    cmd := exec.Command(input)

    // Set the correct output device.
    cmd.Stderr = os.Stderr
    cmd.Stdout = os.Stdout

    // Execute the command and return the error.
    return cmd.Run()
}