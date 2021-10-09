package color

import (
	"fmt"
	"io"
	"os"
)

func Print(color int, a string) {
	text := addColor(color, a)
	_, _ = fmt.Fprint(os.Stdout, text)
}

func Println(color int, a string) {
	text := addColor(color, a)
	_, _ = fmt.Fprintln(os.Stdout, text)
}

func Sprint(color int, a string) string {
	text := addColor(color, a)
	return fmt.Sprint(text)
}

func Fprint(color int, w io.Writer, a string) (n int, err error) {
	text := addColor(color, a)
	return fmt.Fprint(w, text)
}

func Fprintln(color int, w io.Writer, a string) (n int, err error) {
	text := addColor(color, a)
	return fmt.Fprintln(w, text)
}

func addColor(color int, a string) string {
	text := fmt.Sprintf("\x1b[0;%dm%v\x1b[0m", color, a)
	return text
}
