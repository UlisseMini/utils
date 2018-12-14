// Package input implements functions for easily getting user input.
package input

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader  *bufio.Reader  = bufio.NewReader(os.Stdin)
	scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
)

// Input takes a prompt as a string then prints it
// and then returns user input when the user
// enters a newline.
func Input(prompt string) (string, error) {
	fmt.Print(prompt)
	text, err := reader.ReadString('\n')
	return text, err
}

// WaitForEnter holds up proceedings until the user presses return
func WaitForEnter() {
	fmt.Print("[Press enter to proceed]")
	scanner.Scan()
	fmt.Println()
}
