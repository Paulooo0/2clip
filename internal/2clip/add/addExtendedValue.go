package add

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Paulooo0/2clip/internal/2clip/util"
)

func GetExtendedInput() string {
	fmt.Println("Enter your text (type \033[33mEND\033[0m on a new line to finish):")
	reader := bufio.NewReader(os.Stdin)
	var lines []string

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("%s Error reading input: %v", util.Err, err)
			os.Exit(0)
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
