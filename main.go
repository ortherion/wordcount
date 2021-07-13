package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	src, err := readInput()
	if err != nil {
		fail(err)
	}
	fmt.Println(wordCount(src))
}

// match returns true if src matches pattern,
// false otherwise.
func wordCount(src string) int {
	str := strings.Fields(src)
	return len(str)
}

// readInput reads pattern and source string
// from command line arguments and returns them.
func readInput() (src string, err error) {
	flag.Parse()
	src = strings.Join(flag.Args(),"")
	if src == "" {
		return src, errors.New("missing string to count")
	}
	return src, nil
}

// fail prints the error and exits.
func fail(err error) {
	fmt.Println("word count:", err)
	os.Exit(1)
}