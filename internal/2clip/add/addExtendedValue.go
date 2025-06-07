package add

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetExtendedInput() string {
	fmt.Println("Enter your text (type 'END' on a new line to finish):")
	reader := bufio.NewReader(os.Stdin)
	var lines []string

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			os.Exit(1)
		}

		line = strings.TrimSpace(line)
		if line == "END" {
			break
		}

		lines = append(lines, line)
	}

	input := strings.Join(lines, "\n")

	return input
}
