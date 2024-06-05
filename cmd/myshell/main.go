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

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprint(os.Stdout, "$ ")
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)
		fmt.Printf("%s: command not found\n", cmd)
		os.Exit(1)
	}
	
	
}
