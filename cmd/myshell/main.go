package main

import (

	"bufio"

	"fmt"

	"os"

	"os/exec"

	"strings"

)

func main() {

	// You can use print statements as follows for debugging, they'll be visible when running tests.

	// fmt.Println("Logs from your program will appear here!")

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Fprint(os.Stdout, "$ ")

		cmd, _ := reader.ReadString('\n')

		cmd = strings.TrimSpace(cmd) // clean the input

		_, err := exec.LookPath(cmd) // shut the compiler up now

		if err != nil {

			fmt.Printf("%s: command not found\n", cmd)
		}}}