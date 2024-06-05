package main
// Test comment
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	
	// Read a line of input from standard input (keyboard)
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	
	// Print the input (excluding the newline character) followed by ": command not found"
	fmt.Fprint(os.Stdout, input[:len(input)-1]+": command not found\n")
}
