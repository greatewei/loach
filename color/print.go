package color

import (
	"fmt"
	"io"
	"os"
)

// Print is show color text with no newline
func Print(color int, a string) {
	text := addColor(color, a)
	_, _ = fmt.Fprint(os.Stdout, text)
}

// Println is show color text with newline
func Println(color int, a string) {
	text := addColor(color, a)
	_, _ = fmt.Fprintln(os.Stdout, text)
}

// Sprint is return color text string
func Sprint(color int, a string) string {
	text := addColor(color, a)
	return fmt.Sprint(text)
}

// Fprint is return color text string
func Fprint(color int, w io.Writer, a string) (n int, err error) {
	text := addColor(color, a)
	return fmt.Fprint(w, text)
}

// Fprintln is return color text string with newline
func Fprintln(color int, w io.Writer, a string) (n int, err error) {
	text := addColor(color, a)
	return fmt.Fprintln(w, text)
}

func addColor(color int, a string) string {
	text := fmt.Sprintf("\x1b[0;%dm%v\x1b[0m", color, a)
	return text
}
