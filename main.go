package main

import (
	"flag"
	"fmt"
)

var (
	rows int
)

func init() {
	// get number of rows
	flag.IntVar(&rows, "rows", 5, "Number of rows")
}

func main() {
	flag.Parse()
	tree := ""
	// print box
	cols := (rows * 2) - 1 + 4
	for i := 1; i <= cols; i++ {
		if i == 1 || i == cols {
			tree += "+"
		} else {
			tree += "-"
		}
	}
	tree += "\n"
	for i := 1; i <= rows; i++ {
		// set spaces
		tree += "| "
		for j := 1; j <= (rows - i); j++ {
			tree += " "
		}
		// set x
		for j := 1; j <= (i*2)-1; j++ {
			if false {
				tree += " "
			} else {
				tree += "x"
			}
		}
		for j := 1; j <= (rows - i); j++ {
			tree += " "
		}
		tree += " |\n"
	}
	for i := 1; i <= cols; i++ {
		if i == 1 || i == cols {
			tree += "+"
		} else {
			tree += "-"
		}
	}
	fmt.Println(tree)
}
