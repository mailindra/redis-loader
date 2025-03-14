package console

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ConsoleInput struct {
	reader *bufio.Reader
}

func NewConsoleInput() *ConsoleInput {
	return &ConsoleInput{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (c *ConsoleInput) GetNumberInput(prompt string) int {
	for {
		fmt.Print(prompt + ": ")
		input, err := c.reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		count, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			fmt.Println("Please enter a valid number")
			continue
		}

		if count <= 0 {
			fmt.Println("Please enter a positive number")
			continue
		}

		return count
	}
}
