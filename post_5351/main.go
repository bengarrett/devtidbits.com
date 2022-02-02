package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
	"unicode"
)

var ErrParse = errors.New("invalid keypress")

func main() {
	b := YesNo("Is the weather good today?")
	fmt.Println(b)
}

// CtrlC intercepts any Ctrl+C keyboard input and exits to the shell.
func CtrlC() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Fprintf(os.Stdout, " 👋\n")
		os.Exit(0)
	}()
}

// YesNo prints s requesting a boolean yes-or-no answer from the stardard input.
func YesNo(s string) bool {
	CtrlC()
	for {
		fmt.Printf("%s [y/n] ", s)
		input, err := Read(os.Stdin)
		if err != nil {
			log.Fatalln(err)
		}
		if r, err := Parse(input); err == nil {
			return r
		}
	}
}

// Read and return the first rune from the reader.
func Read(r io.Reader) (rune, error) {
	reader := bufio.NewReader(r)
	ru, _, err := reader.ReadRune()
	return ru, err
}

// Parse the rune as a bool.
// Characters y or Y return true, n or N return false.
// Everything else returns an error.
func Parse(r rune) (bool, error) {
	switch unicode.ToLower(r) {
	case rune('y'):
		return true, nil
	case rune('n'):
		return false, nil
	}
	return false, ErrParse
}
