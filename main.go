package main

import (
	"bufio"   //
	"fmt"     //
	"os"      //
	"os/exec" //
	// "path/filepath" //
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to mridulccsh! Type 'exit' to quit.")

	for {
		// print the shell prompt
		fmt.Print("mridulccsh>")

		// read the input from the user
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		// remove the newline character from the input
		if input == "exit" {
			break
		}

		// split the input into command and arguements
		args := strings.Fields(input)

		// handle built-in commands
		switch args[0] {
		case "ls":
			listFiles()
		case "cd":
			if len(args) > 1 {
				changeDir(args[1])
			} else {
				fmt.Println("cd: missing argument")
			}

		case "mkdir":
			if len(args) > 1 {
				makeDir(args[1])
			} else {
				fmt.Println("mkdir: missing argument")

			}
		case "cp":
			if len(args) > 2 {
				copyFile(args[1], args[2])

			} else {
				fmt.Println("cp: missing source or destination")
			}
		case "rm":
			if len(args) > 1 {
				removeFile(args[1])

			} else {
				fmt.Println("rm: missing argument")
			}
		case "cat":
			if len(args) > 1 {
				concatentateFile(args[1])
			} else {
				fmt.Println("cat: missing argument")
			}
		case "grep":
			if len(args) > 2 {
				grepFile(args[1], args[2])
			} else {
				fmt.Println("grep: missing argument")
			}
		case "version":
			fmt.Println("mridulccsh v1.0")
		case "man":
			displayManual()
		default:
			executeCommand(args)
		}
	}

	fmt.Println("Exiting mridulccsh. Goodbye!")
}

// function to list files in the current directory
func listFiles() {
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		fmt.Println(file.Name())

	}
}

// function to change the current working directory
func changeDir(dir string) {
	err := os.Chdir(dir)
	if err != nil {
		fmt.Println("cd:", err)
	}
}

// function to make a new directory
func makeDir(dir string) {
	err := os.Mkdir(dir, 0755)
	if err != nil {
		fmt.Println("mkdir:", err)
	}
}

// function to copy a file
func copyFile(src string, dest string) {
	input, err := os.ReadFile(src)
	if err != nil {
		fmt.Println("cp:", err)
		return
	}
	err = os.WriteFile(dest, input, 0644)
	if err != nil {
		fmt.Println("cp:", err)
	}
}

// function to remove a file
func removeFile(file string) {
	err := os.Remove(file)
	if err != nil {
		fmt.Println("rm:", err)
	}
}

// function to concatenate a file and display its contents
func concatentateFile(file string) {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("cat:", err)
		return
	}

	fmt.Println(string(data))
}

// function to grep a pattern from a file
func grepFile(pattern string, file string) {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("grep:", err)
		return
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.Contains(line, pattern) {
			fmt.Println(line)
		}
	}
}

// function to execute external commands
func executeCommand(args []string) {
	cmd := exec.Command(args[0], args[1:]...)
	cmdOutput, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}
	fmt.Println(string(cmdOutput))
}

// func executeCommand(args []string) {
//     // Check for piping
//     if strings.Contains(strings.Join(args, " "), "|") {
//         commands := strings.Split(strings.Join(args, " "), "|")
//         var lastCmd *exec.Cmd

//         for i, cmd := range commands {
//             cmdArgs := strings.Fields(strings.TrimSpace(cmd))
//             if i == 0 {
//                 lastCmd = exec.Command(cmdArgs[0], cmdArgs[1:]...)
//             } else {
//                 lastCmd.Stdout = lastCmd.Stdout
//             }

//             if i < len(commands)-1 {
//                 nextCmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
//                 nextCmd.Stdin, _ = lastCmd.StdoutPipe()
//                 lastCmd.Stdout = nextCmd.Stdout
//                 lastCmd = nextCmd
//             }
//         }

//         // Start the last command
//         if err := lastCmd.Start(); err != nil {
//             fmt.Println("Error starting command:", err)
//             return
//         }

//         // Wait for the last command to finish
//         if err := lastCmd.Wait(); err != nil {
//             fmt.Println("Error waiting for command:", err)
//             return
//         }
//     } else {
//         cmd := exec.Command(args[0], args[1:]...)
//         cmdOutput, err := cmd.CombinedOutput()
//         if err != nil {
//             fmt.Println("Error executing command:", err)
//             return
//         }
//         fmt.Println(string(cmdOutput))
//     }
// }

// function to display the manual
func displayManual() {
	fmt.Println("Manual for mridulccsh:")
	fmt.Println("ls          - List files in the current directory")
	fmt.Println("cd <dir>   - Change the current directory to <dir>")
	fmt.Println("mkdir <dir>- Create a new directory named <dir>")
	fmt.Println("cp <src> <dest> - Copy a file from <src> to <dest>")
	fmt.Println("rm <file>  - Remove a file named <file>")
	fmt.Println("cat <file> - Concatenate and display the contents of <file>")
	fmt.Println("grep <pattern> <file> - Search for <pattern> in <file>")
	fmt.Println("version     - Show the version of mridulccsh")
	fmt.Println("man         - Show this manual")
}
