package main

import (
	"flag"
	"fmt"
)

var (
	folder   string
	filename string
)

func init() {
	flag.StringVar(&folder, "d", "", "")
	flag.StringVar(&filename, "f", "", "")
}

func main() {
	flag.Parse()
	fmt.Println()
}

// match like // FIXME ....
//
func match(pattern, line string) bool {
	return true
}
