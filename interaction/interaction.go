package interaction

import (
	"bufio"
	"github.com/greatewei/loach/color"
	"os"
	"strings"
)

// ReadInput is interactive operation
func ReadInput(question string) (string, error) {
	if len(question) > 0 {
		color.Print(color.GreenText, question)
	}

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() { // reading
		return "", scanner.Err()
	}

	answer := scanner.Text()
	return strings.TrimSpace(answer), nil
}
