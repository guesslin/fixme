package main

import (
	"flag"
	"fmt"

	"github.com/fatih/color"
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
	color.Cyan("Prints text in cyan.")
}

// match like // FIXME ....
//
func match(pattern, line string) bool {
	fmt.Println(pattern, line)
	return true
}
