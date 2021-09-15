// Center text in a terminal with Go.
// 15-Sep-2021
// https://devtidbits.com/2021/09/15/center-text-in-a-terminal-with-go/
package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode/utf8"

	"golang.org/x/term"
)

func main() {
	const s = "Hello world! 👋"

	// print to the terminal
	buf := Center(s)
	fmt.Println(buf)
	fmt.Fprintln(os.Stdout, buf)
	io.Copy(os.Stdout, buf)

	// save to a text file
	buf = NCenter(80, s)
	if err := os.WriteFile("hi.txt", buf.Bytes(), 0666); err != nil {
		log.Fatal(err)
	}
}

// NCenter centers the string to the column width.
func NCenter(width int, s string) *bytes.Buffer {
	const half, space = 2, "\u0020"
	var b bytes.Buffer
	n := (width - utf8.RuneCountInString(s)) / half
	if n <= 0 {
		fmt.Fprintf(&b, s)
		return &b
	}
	fmt.Fprintf(&b, "%s%s", strings.Repeat(space, int(n)), s)
	return &b
}

// Center the string to the width of the terminal.
// When the width is unknown, the string is left-aligned.
func Center(s string) *bytes.Buffer {
	fd := int(os.Stdin.Fd())
	w, _, err := term.GetSize(fd)
	if err != nil {
		return NCenter(0, s)
	}
	return NCenter(w, s)
}
