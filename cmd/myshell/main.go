package main

import (

	"bufio"

	"fmt"

	"os"

	"strconv"

	"strings"

)

func main() {

	for {

		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input

		cmd, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		cmd = strings.Trim(cmd, "\n") // trim new lines

		handleCmd(cmd)

	}

}

func handleCmd(cmd string) {

	cmd, args := getCmdAndArgs(cmd)

	switch cmd {

	case "exit":

		code, _ := strconv.Atoi(args[0])

		os.Exit(code)

	case "echo":

		fmt.Fprintln(os.Stdout, strings.Join(args, " "))

	default:

		fmt.Fprintf(os.Stdout, "%s: command not found\n", cmd)

	}

}

func getCmdAndArgs(cmd string) (string, []string) {

	l := strings.Split(cmd, " ")

	if len(l) < 2 {

		return l[0], []string{}

	}

	return l[0], l[1:]

}