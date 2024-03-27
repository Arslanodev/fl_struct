package main

import (
	"fl_struct/cmd"
	"flag"
	"fmt"
)
	

func main() {
	flag.Parse()
	
	arg := flag.Args()
	if len(arg) > 0 {
		cmd.Structurize(arg[0])
	} else {
		fmt.Print("Enter dirpath positional argument")
	}
}