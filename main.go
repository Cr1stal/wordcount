package main

import (
  "flag"
  "fmt"
  "os"
  "strings"
)

func main() {
    src := readInput()
    count := wordcount(src)
    fmt.Println(count)
}

// match returns true if src matches pattern,
// false otherwise.
func wordcount(src string) int {
  return len(strings.Fields(src))
}

// readInput reads pattern and source string
// from command line arguments and returns them.
func readInput() (src string) {
	flag.Parse()
	src = strings.Join(flag.Args(), "")
	return src
}

// fail prints the error and exits.
func fail(err error) {
	fmt.Println("wordcount:", err)
	os.Exit(1)
}
