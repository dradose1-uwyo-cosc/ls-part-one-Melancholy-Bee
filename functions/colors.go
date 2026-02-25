package functions

import (
	"bufio"
	"io"
)

type color string

const (
	Reset color = "\033[0m"
	Green color = "\033[01;32m"
	Blue  color = "\033[01;34m"
)

func write_color(c color, w *bufio.Writer) {
	io.WriteString(w, string(c))
}

func reset_color(w *bufio.Writer) {
	io.WriteString(w, string(Reset))
}
